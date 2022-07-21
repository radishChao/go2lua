package injection

import (
	"fmt"
	"go/ast"
	"strings"
)

type (
	luaStruct struct {
		//structFunc    []*luaFunc
		constructFunc *luaFunc
		// 外部定义的字段
		fields   []luaField
		typeName string
		comment  *luaComment
		belong   *luaFile
	}
	luaStructOutField struct {
		typeName string
	}
)

func (ls *luaStruct) getName() string {
	return ls.getVarName()
}

func newLuaStruct(comment *luaComment, belong *luaFile) *luaStruct {
	return &luaStruct{
		belong:  belong,
		comment: comment,
	}
}

// 解析type中定义的struct
func (ls *luaStruct) parseStructFromType(t *ast.TypeSpec) {
	ls.typeName = t.Name.Name
	// 解析字段
	fields := t.Type.(*ast.StructType).Fields.List
	for i := 0; i < len(fields); i++ {
		f := fields[i]
		field := ls.parseField(f)
		if field != nil {
			ls.fields = append(ls.fields, field)
		}

	}
}

// 解析字段
func (ls *luaStruct) parseField(f *ast.Field) luaField {
	name := f.Names[0]
	if !name.IsExported() && f.Tag != nil {
		// 函数的参数 和 返回值 无视是否导出
		return nil
	}
	fTag := parseFieldLuaTag(f.Tag)
	if !fTag.get && !fTag.set {
		fmt.Printf("结构体[%s]字段[%s]未支持get set，忽略导出\n", ls.typeName, name)
		return nil
	}
	lField := ls.parseLuaFieldByExpr(f, f.Type)
	if lField == nil {
		fmt.Printf("结构体[%s]字段[%s]不支持导出\n", ls.typeName, name)
		return nil
	}
	lField.parse(f, fTag.get, fTag.set)
	err := ls.checkFieldByExpr(lField, f.Type)

	if err != nil {
		fmt.Printf("结构体[%s]字段[%s]%s\n", ls.typeName, name, err.Error())
		return nil
	}
	return lField

}
func (ls *luaStruct) parseLuaFieldByArrayType(field *ast.Field, t *ast.ArrayType) luaField {
	var f luaField
	elem := ls.parseLuaFieldByExpr(field, t.Elt)
	if elem != nil {
		f = newLuaArrayField(elem)
	}
	return f
}
func (ls *luaStruct) parseLuaFieldByMapType(field *ast.Field, t *ast.MapType) luaField {
	key := ls.parseLuaFieldByExpr(field, t.Key)
	value := ls.parseLuaFieldByExpr(field, t.Value)
	if key != nil && value != nil {
		return newLuaMapField(key, value)
	}
	return nil

}

func (ls *luaStruct) parseLuaFieldByExpr(field *ast.Field, expr ast.Expr) luaField {
	var (
		f luaField
	)
	switch expr.(type) {
	case *ast.Ident:
		f = ls.parseLuaFieldByIdent(field, expr.(*ast.Ident))
	case *ast.StarExpr:
		// 指针
		f = ls.parseLuaFieldByStart(field, expr.(*ast.StarExpr))
	case *ast.SelectorExpr:
		// 外部定义
		f = ls.parseLuaFieldBySelector(field, expr.(*ast.SelectorExpr))
	case *ast.InterfaceType:
		f = newLuaInterfaceField(ls.getVarName())
	case *ast.ArrayType:
		f = ls.parseLuaFieldByArrayType(field, expr.(*ast.ArrayType))
	case *ast.MapType:
		f = ls.parseLuaFieldByMapType(field, expr.(*ast.MapType))
	}
	return f
}

func (ls *luaStruct) checkArrayFiled(at *ast.ArrayType, lField luaField) error {
	switch lField.(type) {
	case *luaArrayField:
		af := lField.(*luaArrayField)
		if af.elem == nil {
			return fmt.Errorf("数组元素类型不支持导出")
		}
		switch at.Elt.(type) {
		case *ast.InterfaceType:
			return fmt.Errorf("不支持数组元素为interface的导出")
		case *ast.StarExpr:
			// 指针 只接受结构体指针
			switch af.elem.(type) {
			case *luaStructField:
			default:
				p := af.elem.getTypePackage()
				if p == ls.belong.sourcePackage || p == "" {
					return fmt.Errorf("不可导出,只接受%s", af.elem.getType())
				}
				return fmt.Errorf("不可导出,只接受%s.%s", p, af.elem.getType())
			}
		case *ast.Ident:
			switch af.elem.(type) {
			case *luaStructField:
				p := af.elem.getTypePackage()
				if p == ls.belong.sourcePackage || p == "" {
					return fmt.Errorf("不可导出,数组元素类型只接受*%s", af.elem.getType())
				}
				return fmt.Errorf("不可导出,数组元素类型只接受*%s.%s", p, af.elem.getType())
			case *luaArrayField:
				return fmt.Errorf("不支持多维数组的导出,建议使用函数调用的方式操作多维数组")
			}
		case *ast.SelectorExpr:
			switch af.elem.(type) {
			case *luaStructField:
				return fmt.Errorf("不可导出,数组元素类型只接受*%s.%s", af.elem.getTypePackage(), af.elem.getType())
			}

		case *ast.ArrayType:
			return fmt.Errorf("不支持多维数组的导出,建议使用函数调用的方式操作多维数组")

		}

	}
	return nil
}

func (ls *luaStruct) checkFieldByExpr(lField luaField, expr ast.Expr) error {
	switch expr.(type) {
	case *ast.Ident:
		switch lField.(type) {
		case *luaStructField:
			return fmt.Errorf("不可导出,只接受*%s", lField.getType())
		case *luaArrayField:
			return ls.checkArrayFiled(expr.(*ast.Ident).Obj.Decl.(*ast.TypeSpec).Type.(*ast.ArrayType), lField)
		default:
		}
	case *ast.StarExpr:
		// 指针 只接受结构体指针
		switch lField.(type) {
		case *luaStructField:
		default:
			p := lField.getTypePackage()
			if p == ls.belong.sourcePackage || p == "" {
				return fmt.Errorf("不可导出,只接受%s", lField.getType())
			}
			return fmt.Errorf("不可导出,只接受%s.%s", p, lField.getType())
		}

	case *ast.SelectorExpr:
		// 外部定义
		switch lField.(type) {
		case *luaStructField:
			return fmt.Errorf("不可导出,只接受*%s.%s", lField.getTypePackage(), lField.getType())
		case *luaArrayField:
			//expr.(*ast.SelectorExpr).X
			//return ls.checkArrayFiled(lField.(*luaArrayField))
		}
	case *ast.ArrayType:
		return ls.checkArrayFiled(expr.(*ast.ArrayType), lField)
	}
	return nil
}

func (ls *luaStruct) parseLuaFieldByIdent(field *ast.Field, ident *ast.Ident) luaField {
	name := ident.Name
	f := newLuaFieldByTyName(name)
	if f == nil && ident.Obj != nil {
		switch ident.Obj.Decl.(type) {
		case *ast.TypeSpec:
			// 在文件内的type定义的类型
			t := ident.Obj.Decl.(*ast.TypeSpec).Type
			switch t.(type) {
			case *ast.StructType:
				f = newLuaStructField("", ls.belong.sourcePackage)
			case *ast.Ident:
				// 定义的基础数据类型
				f = ls.parseLuaFieldByIdent(field, t.(*ast.Ident))
			case *ast.ArrayType:
				f = ls.parseLuaFieldByArrayType(field, t.(*ast.ArrayType))
			case *ast.InterfaceType:
				f = newLuaInterfaceField(ls.getVarName())
			}
		}
	}

	if f != nil {
		f.setType(name)
		f.setTypePackage(ls.belong.sourcePackage)
	}
	return f

}
func (ls *luaStruct) parseLuaFieldByStart(field *ast.Field, expr *ast.StarExpr) luaField {
	var (
		f luaField
	)
	switch expr.X.(type) {
	case *ast.Ident:
		f = ls.parseLuaFieldByIdent(field, expr.X.(*ast.Ident))
	case *ast.SelectorExpr:
		f = ls.parseLuaFieldBySelector(field, expr.X.(*ast.SelectorExpr))
	case *ast.InterfaceType:
		f = newLuaInterfaceField(ls.belong.sourcePackage)
	case *ast.ArrayType:
		f = ls.parseLuaFieldByArrayType(field, expr.X.(*ast.ArrayType))

	}
	return f
}

func (ls *luaStruct) parseLuaFieldBySelector(field *ast.Field, expr *ast.SelectorExpr) luaField {
	var (
		f   luaField
		typ string
	)

	typeTag := parseFieldLuaTypeTag(field.Tag)
	sourcePackage := expr.X.(*ast.Ident).Name
	if typeTag != nil {
		typ = typeTag.typ
	} else {
		fieldName := field.Names[0].Name
		outField, ok := ls.comment.fields[fieldName]
		if ok {
			typ = outField.typeName
		}

	}
	if typ == "" {
		return nil
	}
	switch typ {
	case typeStruct:
		// 获取到tolua的包名
		outFile, ok := ls.belong.comment.packages[sourcePackage]
		if !ok {
			return nil
		}
		s := outFile.short
		if outFile.short == ls.belong.targetPackage {
			s = ""
		}

		f = newLuaStructField(s, sourcePackage)
	default:
		f = newLuaFieldByTyName(typ)
	}
	if f != nil {
		f.setType(expr.Sel.Name)
		f.setTypePackage(sourcePackage)
	}
	return f

}

func (ls *luaStruct) getFullName() string {
	return fmt.Sprintf("%s.%s", ls.belong.sourcePackage, ls.typeName)
}
func (ls *luaStruct) getVarName() string {
	s := ls.typeName
	return strings.ToLower(string(s[0])) + s[1:]
}

func (ls *luaStruct) getAddressName() string {
	return "&" + ls.getFullName()
}
func (ls *luaStruct) getBelongSourcePackageUpper() string {
	return strings.ToUpper(string(ls.belong.sourcePackage[0])) + ls.belong.sourcePackage[1:]
}

// 生成文件var域下的代码
func (ls *luaStruct) genVar(coder *codeWriter) {
	fLen := len(ls.fields)
	if fLen == 0 {
		return
	}
	coder.writeString("map")
	coder.writeString(ls.typeName)
	coder.writeString("ToLuaMethods =  map[string]lua.LGFunction{")
	coder.writeLineEnd()
	for i := 0; i < fLen; i++ {
		f := ls.fields[i]
		coder.writeQuote()
		coder.writeString(f.getFieldName())
		coder.writeQuote()
		coder.writeString(":")
		coder.writeString("register")
		coder.writeString(ls.typeName)
		coder.writeString(f.getFieldName())
		coder.writeLine(",")
	}
	coder.writeString(rbrace)
	coder.writeLineEnd()
}

// 生成注入代码
func (ls *luaStruct) genRegisterCode(coder *codeWriter) {
	// 注入结构体
	coder.writeString("err = ")
	coder.writeString(toLuaHelperShort)
	coder.writeString(".RegisterStruct(")
	coder.writeString(luaStateShort)
	coder.writeString(",")
	coder.writeQuote()
	coder.writeString(ls.belong.luaPackage)
	coder.writeQuote()
	coder.writeString(",")
	coder.writeQuote()
	coder.writeString(ls.typeName)
	coder.writeQuote()
	coder.writeString(",(*")
	coder.writeString(ls.belong.luaPackage)
	coder.writeString(".")
	coder.writeString(ls.typeName)
	coder.writeString(")(nil),")

	coder.writeString("new")
	coder.writeString(ls.getBelongSourcePackageUpper())
	coder.writeString(ls.typeName)
	coder.writeString("ToLua,")
	if len(ls.fields) != 0 {
		coder.writeString("map")
		coder.writeString(ls.typeName)
		coder.writeLine("ToLuaMethods)")
	} else {
		coder.writeLine("nil)")
	}
	coder.writeLine("if err != nil {")
	coder.writeLine("return err")
	coder.writeLine("}")
}

func (ls *luaStruct) genGenUserDataBlockStmt(coder *codeWriter) {
	coder.writeString("func GenUserData")
	coder.writeString(ls.getBelongSourcePackageUpper())
	coder.writeString(ls.typeName)
	coder.writeString("(")
	coder.writeString(luaStateShort)
	coder.writeString(" *")
	coder.writeString(gopherLuaShort)
	coder.writeString(".LState, ")
	coder.writeString(ls.getVarName())
	coder.writeString(" *")
	coder.writeString(ls.getFullName())
	coder.writeString(")*")
	coder.writeString(gopherLuaShort)
	coder.writeString(".LUserData")
	coder.writeLine("{")

	coder.writeString("ud := ")
	coder.writeString(luaStateShort)
	coder.writeLine(".NewUserData()")
	coder.writeString("ud.Value = ")
	coder.writeLine(ls.getVarName())
	coder.writeString(luaStateShort)
	coder.writeString(".SetMetatable(ud, L.GetTypeMetatable(\"")
	//coder.writeString(ls.belong.sourcePackage)
	//coder.writeString(".")
	coder.writeString(ls.typeName)
	coder.writeQuote()
	coder.writeLine("))")
	coder.writeLine("return ud")
	coder.writeLine(rbrace)
}

// 生成push代码
func (ls *luaStruct) genPushBlockStmt(coder *codeWriter) {
	coder.writeString("func Push")
	coder.writeString(ls.getBelongSourcePackageUpper())
	coder.writeString(ls.typeName)
	coder.writeString("ToLua(")
	coder.writeString(luaStateShort)
	coder.writeString(" *")
	coder.writeString(gopherLuaShort)
	coder.writeString(".LState, ")
	coder.writeString(ls.getVarName())
	coder.writeString(" *")
	coder.writeString(ls.getFullName())
	coder.writeLine("){")
	coder.writeString(luaStateShort)
	coder.writeString(".Push(")
	coder.writeString("GenUserData")
	coder.writeString(ls.getBelongSourcePackageUpper())
	coder.writeString(ls.typeName)
	coder.writeString("(")
	coder.writeString(luaStateShort)
	coder.writeString(",")
	coder.writeString(ls.getVarName())
	coder.writeString(")")
	coder.writeLine(")")
	coder.writeLine(rbrace)
}

// 生成check代码
func (ls *luaStruct) genCheckBlockStmt(coder *codeWriter) {
	coder.writeString("func Check")
	coder.writeString(ls.getBelongSourcePackageUpper())
	coder.writeString(ls.typeName)
	coder.writeString("ToLua(")
	coder.writeString(luaStateShort)
	coder.writeString(" *")
	coder.writeString(gopherLuaShort)
	coder.writeString(".LState, index int)*")
	coder.writeString(ls.getFullName())
	coder.writeLine(lbrace)
	coder.writeString("ud := ")
	coder.writeString(luaStateShort)
	coder.writeLine(".CheckUserData(index)")
	coder.writeString("if v, ok := ud.Value.(*")
	coder.writeString(ls.getFullName())
	coder.writeLine("); ok {")
	coder.writeLine("return v")
	coder.writeLine(rbrace)
	coder.writeString(luaStateShort)
	coder.writeString(".ArgError(index,\"")
	coder.writeString(ls.getFullName())
	coder.writeLine(" expected\")")
	coder.writeLine("return nil")
	coder.writeLine(rbrace)
}

// 生成构造函数代码
func (ls *luaStruct) genConstructorBlockStmt(coder *codeWriter) {
	if ls.constructFunc == nil {
		coder.writeString("func ")
		coder.writeString("new")
		coder.writeString(ls.getBelongSourcePackageUpper())
		coder.writeString(ls.typeName)
		coder.writeString("ToLua(")
		coder.writeString(luaStateShort)
		coder.writeLine(" *lua.LState)int{")
		coder.writeString("Push")
		coder.writeString(ls.getBelongSourcePackageUpper())
		coder.writeString(ls.typeName)
		coder.writeString("ToLua(")
		coder.writeString(luaStateShort)
		coder.writeString(",")
		coder.writeString(ls.getAddressName())
		coder.writeLine("{})")
		coder.writeLine("return 1}")
	} else {
		ls.constructFunc.genBlockStmt(coder)
	}
}

// 生成函数代码
func (ls *luaStruct) genFieldGetSetBlockStmt(coder *codeWriter) {
	for i := 0; i < len(ls.fields); i++ {
		f := ls.fields[i]
		funcName := fmt.Sprintf("register%s%s", ls.typeName, f.getFieldName())
		coder.writeString("// ")
		coder.writeString(funcName)
		coder.writeSpace()
		coder.writeString(ls.typeName)
		coder.writeString(".")
		coder.writeString(f.getFieldName())
		if f.getter() {
			coder.writeSpace()
			coder.writeString("get")
		}
		if f.setter() {
			coder.writeSpace()
			coder.writeString("set")
		}
		coder.writeLineEnd()
		comment := f.getComment()
		if len(comment) != 0 {
			for j := 0; j < len(comment); j++ {
				coder.writeLine(comment[j])
			}
		}

		coder.writeString("func ")
		coder.writeString(funcName)
		coder.writeString("(")
		coder.writeString(luaStateShort)
		coder.writeString(" *")
		coder.writeString(gopherLuaShort)
		coder.writeLine(".LState)int{")
		coder.writeString(ls.getVarName())
		coder.writeString(" := Check")
		coder.writeString(ls.getBelongSourcePackageUpper())
		coder.writeString(ls.typeName)
		coder.writeString("ToLua(")
		coder.writeString(luaStateShort)
		coder.writeLine(",1)")
		if f.setter() {
			coder.writeString("if ")
			coder.writeString(luaStateShort)
			coder.writeLine(".GetTop() == 2{")
			varName := fmt.Sprintf("%s.%s", ls.getVarName(), f.getFieldName())
			coder.writeLine(f.check(varName, 2))
			if f.getter() {
				coder.writeLine("return 0")
			}
			coder.writeLine("}")
		}
		if f.getter() {
			fieldName := fmt.Sprintf("%s.%s", ls.getVarName(), f.getFieldName())
			coder.writeLine(f.push(fieldName))
			coder.writeLine("return 1")
		} else {
			coder.writeLine("return 0")
		}
		coder.writeLine("}")
	}

}

// 生成代码
func (ls *luaStruct) genBlockStmt(coder *codeWriter) {
	ls.genGenUserDataBlockStmt(coder)
	ls.genPushBlockStmt(coder)
	ls.genCheckBlockStmt(coder)
	ls.genConstructorBlockStmt(coder)
	ls.genFieldGetSetBlockStmt(coder)
}
