package injection

type (
	luaFunc struct {
		// 参数
		param []luaField
		// 返回值
		result []luaField
		// 函数名
		funcName string
		// 注入到lua的函数名
		funcLuaName string
		// 根据字段生成函数
		fromField luaField
	}
)

//func newLuaFunc(caller luaType, funcName, funcLuaName string, param []luaField, result []luaField) *luaFunc {
//	return &luaFunc{
//		luaTypeBase: luaTypeBase{
//			caller:  caller,
//		},
//		param:       param,
//		result:      result,
//		funcName:    funcName,
//		funcLuaName: funcLuaName,
//	}
//}

//func newLuaFuncFromField(funcName string, fromField luaField) *luaFunc {
//
//	funcLuaName := fromField.getFieldName()
//
//	return &luaFunc{
//		funcLuaName: funcLuaName,
//		funcName:    funcName,
//		fromField:   fromField,
//	}
//}

func (lf *luaFunc) genVar(coder *codeWriter) {
	coder.writeQuote()
	coder.writeString(lf.funcLuaName)
	coder.writeQuote()
	coder.writeString(": ")
	coder.writeString(lf.funcName)
	coder.writeLine(",")
}

// 生成字段的set代码
func (lf *luaFunc) genBlockStmtWithFieldSet(coder *codeWriter, isReturn bool) {
	//coder.writeString("if ")
	//coder.writeString(luaStateShort)
	//coder.writeString(".GetTop() == ")
	//if lf.callerIsStruct() {
	//	coder.writeString("2")
	//} else {
	//	coder.writeString("1")
	//}
	//coder.writeLine(" {")
	//coder.writeString(lf.getCaller().getVarName())
	//coder.writeString(".")
	//coder.writeString(lf.fromField.getFieldName())
	//coder.writeString(" = ")
	//coder.writeLine(lf.fromField.getLuaType().check(2))
	//if isReturn {
	//	coder.writeLine("return 0")
	//}
	//coder.writeLine(" }")
}

// 生成字段的get代码
func (lf *luaFunc) genBlockStmtWithFieldGet(coder *codeWriter) {
	//coder.writeLine(lf.fromField.getLuaType().push(fmt.Sprintf("%s.%s", lf.getCaller().getVarName(), lf.fromField.getFieldName())))
	//coder.writeLine("return 1")
}

// 根据字段的get set生成代码
func (lf *luaFunc) genBlockStmtWithFieldGetSet(coder *codeWriter) {
	if lf.fromField.getter() && lf.fromField.setter() {

		lf.genBlockStmtWithFieldSet(coder, true)
		lf.genBlockStmtWithFieldGet(coder)
		return
	}
	if lf.fromField.setter() {
		lf.genBlockStmtWithFieldSet(coder, false)
		coder.writeLine("return 0")
		return
	}
	if lf.fromField.getter() {
		lf.genBlockStmtWithFieldGet(coder)
	}

}

// 生成代码
func (lf *luaFunc) genBlockStmt(coder *codeWriter) {
	//coder.writeString("// ")
	//coder.writeString(lf.funcName)
	//coder.writeSpace()
	//coder.writeString(lf.caller.getTypeName())
	//coder.writeString(".")
	//coder.writeLine(lf.funcLuaName)
	//for i := 0; i < len(lf.comment); i++ {
	//	coder.writeString(lf.comment[i])
	//	coder.writeLineEnd()
	//}
	//coder.writeString("func ")
	//coder.writeString(lf.funcName)
	//coder.writeString("(")
	//coder.writeString(luaStateShort)
	//coder.writeSpace()
	//coder.writeString("*")
	//coder.writeString(gopherLuaShort)
	//coder.writeLine(".LState) int{")
	//coder.writeLine(lf.checkCaller(1))
	//if lf.fromField != nil {
	//	lf.genBlockStmtWithFieldGetSet(coder)
	//}
	//
	//coder.writeLine("}")
	//
	//fmt.Println("luaFunc genBlockStmt", lf.funcName, lf.funcLuaName, lf.caller.getTypeName(), len(lf.param), len(lf.result))
}
