package injection

import "fmt"

type (
	luaArrayField struct {
		luaBaseField
		elem luaField
	}
)

func (l *luaArrayField) luaValueToInterface(varName string) string {
	return ""
}

func (l *luaArrayField) interfaceToLuaValue(varName string) string {
	return ""
}

func newLuaArrayField(elem luaField) *luaArrayField {
	return &luaArrayField{luaBaseField: luaBaseField{}, elem: elem}
}

func (l *luaArrayField) check(varName string, index int) string {
	writer := newCodeWriter()
	writer.writeString(varName)
	writer.writeString(" = ")
	writer.writeString(varName)
	writer.writeString("[len(")
	writer.writeString(varName)
	writer.writeLine("):1]")
	writer.writeString("luaT := ")
	writer.writeString(luaStateShort)
	writer.writeString(".CheckTable(")
	writer.writeString(fmt.Sprintf("%d", index))
	writer.writeLine(")")
	writer.writeLine("luaTLen := luaT.Len()")
	writer.writeLine("for i := 0; i < luaTLen; i++ {")
	writer.writeString(varName)
	writer.writeString(" = ")
	writer.writeString("append(")
	writer.writeString(varName)
	writer.writeString(",")
	writer.writeString(l.elem.luaValueToInterface("luaT.RawGetInt(i+1)"))
	writer.writeLine(")")
	writer.writeLine(rbrace)
	return writer.code()
}

func (l *luaArrayField) push(varName string) string {
	writer := newCodeWriter()
	writer.writeLine("luaT := L.NewTable()")
	writer.writeString("for i := 0; i < len(")
	writer.writeString(varName)
	writer.writeLine("); i++ {")
	writer.writeString("luaT.Append(")

	loopValueName := fmt.Sprintf("%s[i]", varName)
	writer.writeString(l.elem.interfaceToLuaValue(loopValueName))
	writer.writeLine(")")
	writer.writeLine("}")
	writer.writeString(luaStateShort)
	writer.writeLine(".Push(luaT)")
	return writer.code()
}
