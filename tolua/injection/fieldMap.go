package injection

import (
	"strconv"
)

type (
	luaMapField struct {
		luaBaseField
		key   luaField
		value luaField
	}
)

func (l *luaMapField) check(varName string, index int) string {
	writer := newCodeWriter()
	writer.writeString(varName)
	writer.writeString(" = make(map[")
	writer.writeString(l.key.getType())
	writer.writeString("]")
	writer.writeString(l.value.getType())
	writer.writeLine(")")

	writer.writeString("t := ")
	writer.writeString(luaStateShort)
	writer.writeString(".CheckTable(")
	writer.writeString(strconv.Itoa(index))
	writer.writeLine(")")
	writer.writeLine("t.ForEach(func(key lua.LValue, value lua.LValue) {")

	writer.writeString(varName)
	writer.writeString("[")
	writer.writeString(l.key.luaValueToInterface("key"))
	writer.writeString("]=")
	writer.writeLine(l.value.luaValueToInterface("value"))
	writer.writeString(rbrace)
	writer.writeLine(")")
	return writer.code()
}

func (l *luaMapField) push(varName string) string {
	writer := newCodeWriter()
	writer.writeString("t := ")
	writer.writeString(luaStateShort)
	writer.writeLine(".NewTable()")
	writer.writeString("for k, v := range ")
	writer.writeString(varName)
	writer.writeLine(" {")
	writer.writeString("t.RawSet(")
	writer.writeString(l.key.interfaceToLuaValue("k"))
	writer.writeString(",")
	writer.writeString(l.value.interfaceToLuaValue("v"))
	writer.writeLine(")")
	writer.writeLine("}")
	return writer.code()
}

func (l *luaMapField) luaValueToInterface(varName string) string {
	return "map luaValueToInterface"
}

func (l *luaMapField) interfaceToLuaValue(varName string) string {
	return "map interfaceToLuaValue"
}

func newLuaMapField(key, value luaField) *luaMapField {
	return &luaMapField{luaBaseField{}, key, value}
}
