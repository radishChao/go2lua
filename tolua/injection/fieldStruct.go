package injection

import (
	"fmt"
	"strings"
)

type (
	luaStructField struct {
		luaBaseField
	}
)

func newLuaStructField(callPackage, typePackage string) *luaStructField {
	return &luaStructField{luaBaseField{typePackage: typePackage, callPackage: callPackage}}
}

func (lsf *luaStructField) getTypePackageUpper() string {
	return strings.ToUpper(string(lsf.typePackage[0])) + lsf.typePackage[1:]
}
func (lsf *luaStructField) check(varName string, index int) string {
	c := fmt.Sprintf(" Check%s%sToLua(%s,%d)", lsf.getTypePackageUpper(), lsf.typ, luaStateShort, index)
	if lsf.callPackage != "" {
		c = fmt.Sprintf("%s.%s", lsf.callPackage, c)
	}
	c = fmt.Sprintf("%s = %s", varName, c)
	return c
}

func (lsf *luaStructField) push(varName string) string {
	c := fmt.Sprintf("Push%s%sToLua(%s,%s)", lsf.getTypePackageUpper(), lsf.typ, luaStateShort, varName)
	if lsf.callPackage != "" {
		c = fmt.Sprintf("%s.%s", lsf.callPackage, c)
	}
	return c
}

func (lsf *luaStructField) luaValueToInterface(varName string) string {
	return fmt.Sprintf("%s.(*lua.LUserData).Value.(*%s.%s)", varName, lsf.typePackage, lsf.typ)
}

func (lsf *luaStructField) interfaceToLuaValue(varName string) string {
	if lsf.callPackage != "" {
		return fmt.Sprintf("%s.GenUserData%s%s(%s,%s)", lsf.callPackage, lsf.getTypePackageUpper(), lsf.typ, luaStateShort, varName)
	}
	return fmt.Sprintf("GenUserData%s%s(%s,%s)", lsf.getTypePackageUpper(), lsf.typ, luaStateShort, varName)
}
func (lsf *luaStructField) getType() string {
	if lsf.typePackage == lsf.callPackage {
		return "*" + lsf.typ
	}

	return "*" + lsf.typePackage + "." + lsf.typ
}
