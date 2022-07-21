package injection

import (
	"fmt"
	"go/ast"
)

type (
	luaField interface {
		getComment() []string
		getFieldName() string
		getter() bool
		setter() bool
		parse(f *ast.Field, get, set bool)
		check(varName string, index int) string
		push(varName string) string
		setType(s string)
		getType() string
		setTypePackage(s string)
		getTypePackage() string
		luaValueToInterface(varName string) string
		interfaceToLuaValue(varName string) string
	}
	luaBaseField struct {
		comment     *luaComment
		callPackage string
		typePackage string
		typ         string
		fieldName   string
		get         bool
		set         bool
	}

	luaInterfaceField struct {
		luaBaseField
	}
)

func newLuaInterfaceField(callPackage string) *luaInterfaceField {
	return &luaInterfaceField{luaBaseField{callPackage: callPackage}}
}

func (l *luaInterfaceField) check(varName string, index int) string {
	return fmt.Sprintf("%s = %s.Check(%s,%s.%s,%d)", varName, toLuaHelperShort, luaStateShort,
		l.callPackage, l.getFieldName(), index)
}

func (l *luaInterfaceField) push(varName string) string {
	return fmt.Sprintf("%s.Push(%s,%s.%s)", toLuaHelperShort, luaStateShort, l.callPackage, l.getFieldName())
}

func (l *luaInterfaceField) luaValueToInterface(varName string) string {
	return fmt.Sprintf("luaHelper.LuaValueToInterface(%s)", varName)
}
func (l *luaInterfaceField) interfaceToLuaValue(varName string) string {
	return fmt.Sprintf("luaHelper.InterfaceToLuaValue(%s, %s)", luaStateShort, varName)
}

func (lf *luaBaseField) getComment() []string {
	return lf.comment.comment
}

func (lf *luaBaseField) parse(f *ast.Field, get, set bool) {
	lf.fieldName = f.Names[0].Name
	lf.comment = parseLuaComment(f.Comment)
	lf.get = get
	lf.set = set
}

func (lf *luaBaseField) getFieldName() string {
	return lf.fieldName
}

func (lf *luaBaseField) getter() bool {
	return lf.get
}

func (lf *luaBaseField) setter() bool {
	return lf.set
}

func (lf *luaBaseField) setType(s string) {
	lf.typ = s
}
func (lf *luaBaseField) getType() string {
	return lf.typ
}

func (lf *luaBaseField) setTypePackage(s string) {
	lf.typePackage = s
}
func (lf *luaBaseField) getTypePackage() string {
	return lf.typePackage
}
