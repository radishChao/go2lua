package injection

import "fmt"

type (
	luaFieldString struct {
		luaBaseField
	}
	luaFieldUint8 struct {
		luaBaseField
	}
	luaFieldInt8 struct {
		luaBaseField
	}
	luaFieldBool struct {
		luaBaseField
	}
	luaFieldUint16 struct {
		luaBaseField
	}
	luaFieldInt16 struct {
		luaBaseField
	}
	luaFieldUint struct {
		luaBaseField
	}
	luaFieldInt struct {
		luaBaseField
	}
	luaFieldUint32 struct {
		luaBaseField
	}
	luaFieldInt32 struct {
		luaBaseField
	}
	luaFieldUint64 struct {
		luaBaseField
	}
	luaFieldInt64 struct {
		luaBaseField
	}
	luaFieldFloat32 struct {
		luaBaseField
	}
	luaFieldFloat64 struct {
		luaBaseField
	}
)

// 基本数据类型
const (
	typeString  = "string"
	typeInt8    = "int8"
	typeUInt8   = "uint8"
	typeBool    = "bool"
	typeByte    = "byte"
	typeInt16   = "int16"
	typeUInt16  = "uint16"
	typeInt     = "int"
	typeUInt    = "uint"
	typeInt32   = "int32"
	typeUInt32  = "uint32"
	typeInt64   = "int64"
	typeUInt64  = "uint64"
	typeFloat32 = "float32"
	typeFloat64 = "float64"
)

// 复合数据类型
const (
	typeStruct = "struct"
)

func newLuaFieldString() *luaFieldString {
	return &luaFieldString{luaBaseField{}}
}

func newLuaFieldInt8() *luaFieldInt8 {
	return &luaFieldInt8{luaBaseField{}}
}

func newLuaFieldUint8() *luaFieldUint8 {
	return &luaFieldUint8{luaBaseField{}}
}
func newLuaFieldBool() *luaFieldBool {
	return &luaFieldBool{luaBaseField{}}
}

func newLuaFieldInt16() *luaFieldInt16 {
	return &luaFieldInt16{luaBaseField{}}
}

func newLuaFieldUint16() *luaFieldUint16 {
	return &luaFieldUint16{luaBaseField{}}
}

func newLuaFieldInt() *luaFieldInt {
	return &luaFieldInt{luaBaseField{}}
}

func newLuaFieldUint() *luaFieldUint {
	return &luaFieldUint{luaBaseField{}}
}
func newLuaFieldInt32() *luaFieldInt32 {
	return &luaFieldInt32{luaBaseField{}}
}

func newLuaFieldUint32() *luaFieldUint32 {
	return &luaFieldUint32{luaBaseField{}}
}

func newLuaFieldInt64() *luaFieldInt64 {
	return &luaFieldInt64{luaBaseField{}}
}

func newLuaFieldUint64() *luaFieldUint64 {
	return &luaFieldUint64{luaBaseField{}}
}

func newLuaFieldFloat32() *luaFieldFloat32 {
	return &luaFieldFloat32{luaBaseField{}}
}

func newLuaFieldFloat64() *luaFieldFloat64 {
	return &luaFieldFloat64{luaBaseField{}}
}

func checkNumber(index int, varName, typ, targetType, typePackage string) string {
	c := fmt.Sprintf("%s.CheckNumber(%d)", luaStateShort, index)
	if typ != targetType {
		c = fmt.Sprintf("%s.%s(%s)", typePackage, typ, c)
	} else {
		c = fmt.Sprintf("%s(%s)", targetType, c)
	}
	c = fmt.Sprintf("%s = %s", varName, c)
	return c
}

func luaValueToInterface(varName, typ, targetType, typePackage, luaType string) string {
	if typ != targetType {
		return fmt.Sprintf("%s.%s(%s.(%s.%s))", typePackage, typ, varName, gopherLuaShort, luaType)
	}
	return fmt.Sprintf("%s(%s.(%s.%s))", targetType, varName, gopherLuaShort, luaType)
}

// string check
func (fieldString *luaFieldString) check(varName string, index int) string {
	c := fmt.Sprintf("%s.CheckString(%d)", luaStateShort, index)
	if fieldString.typ != typeString {
		c = fmt.Sprintf("%s.%s(%s)", fieldString.typePackage, fieldString.typ, c)
	}
	c = fmt.Sprintf("%s = %s", varName, c)
	return c
}

func (fieldString *luaFieldString) luaValueToInterface(varName string) string {
	return luaValueToInterface(varName, fieldString.typ, typeString, fieldString.typePackage, "LString")
}

func (fieldString *luaFieldString) interfaceToLuaValue(varName string) string {
	return fmt.Sprintf("%s.LString(%s)", gopherLuaShort, varName)
}

// string push
func (fieldString *luaFieldString) push(varName string) string {
	return fmt.Sprintf("%s.Push(%s.LString(%s))", luaStateShort, gopherLuaShort, varName)
}

// uint8 byte check
func (fieldUint8 *luaFieldUint8) check(varName string, index int) string {
	if fieldUint8.typ == typeByte {
		return checkNumber(index, varName, fieldUint8.typ, typeByte, fieldUint8.typePackage)
	}
	return checkNumber(index, varName, fieldUint8.typ, typeUInt8, fieldUint8.typePackage)
}

func (fieldUint8 *luaFieldUint8) luaValueToInterface(varName string) string {

	if fieldUint8.typ == typeByte {
		return luaValueToInterface(varName, fieldUint8.typ, typeByte, fieldUint8.typePackage, "LNumber")
	}
	return luaValueToInterface(varName, fieldUint8.typ, typeUInt8, fieldUint8.typePackage, "LNumber")
}

func (fieldUint8 *luaFieldUint8) interfaceToLuaValue(varName string) string {
	return fmt.Sprintf("%s.LNumber(%s)", gopherLuaShort, varName)
}

// uint8 byte push
func (fieldUint8 *luaFieldUint8) push(varName string) string {
	return fmt.Sprintf("%s.Push(%s.LNumber(%s))", luaStateShort, gopherLuaShort, varName)
}

// int8 check
func (fieldInt8 *luaFieldInt8) check(varName string, index int) string {

	return checkNumber(index, varName, fieldInt8.typ, typeInt8, fieldInt8.typePackage)
}

// int8 push
func (fieldInt8 *luaFieldInt8) push(varName string) string {
	return fmt.Sprintf("%s.Push(%s.LNumber(%s))", luaStateShort, gopherLuaShort, varName)
}

func (fieldInt8 *luaFieldInt8) luaValueToInterface(varName string) string {
	return luaValueToInterface(varName, fieldInt8.typ, typeInt8, fieldInt8.typePackage, "LNumber")
}
func (fieldInt8 *luaFieldInt8) interfaceToLuaValue(varName string) string {
	return fmt.Sprintf("%s.LNumber(%s)", gopherLuaShort, varName)
}

// bool check
func (fieldBool *luaFieldBool) check(varName string, index int) string {
	c := fmt.Sprintf("%s.CheckBool(%d)", luaStateShort, index)
	if fieldBool.typ != typeBool {
		c = fmt.Sprintf("%s.%s(%s)", fieldBool.typePackage, fieldBool.typ, c)
	} else {
		c = fmt.Sprintf("%s(%s)", typeBool, c)
	}
	c = fmt.Sprintf("%s = %s", varName, c)
	return c
}

// bool push
func (fieldBool *luaFieldBool) push(varName string) string {
	return fmt.Sprintf("%s.Push(%s.LBool(%s))", luaStateShort, gopherLuaShort, varName)
}

func (fieldBool *luaFieldBool) interfaceToLuaValue(varName string) string {
	return fmt.Sprintf("%s.LBool(%s)", gopherLuaShort, varName)
}
func (fieldBool *luaFieldBool) luaValueToInterface(varName string) string {
	return luaValueToInterface(varName, fieldBool.typ, typeBool, fieldBool.typePackage, "LBool")
}

// uint16 check
func (fieldUint16 *luaFieldUint16) check(varName string, index int) string {
	return checkNumber(index, varName, fieldUint16.typ, typeUInt16, fieldUint16.typePackage)
}

func (fieldUint16 *luaFieldUint16) luaValueToInterface(varName string) string {
	return luaValueToInterface(varName, fieldUint16.typ, typeUInt16, fieldUint16.typePackage, "LNumber")
}
func (fieldUint16 *luaFieldUint16) interfaceToLuaValue(varName string) string {
	return fmt.Sprintf("%s.LNumber(%s)", gopherLuaShort, varName)
}

// uint16 push
func (fieldUint16 *luaFieldUint16) push(varName string) string {
	return fmt.Sprintf("%s.Push(%s.LNumber(%s))", luaStateShort, gopherLuaShort, varName)
}

// int16 check
func (fieldInt16 *luaFieldInt16) check(varName string, index int) string {
	return checkNumber(index, varName, fieldInt16.typ, typeInt16, fieldInt16.typePackage)
}

// int16 push
func (fieldInt16 *luaFieldInt16) push(varName string) string {
	return fmt.Sprintf("%s.Push(%s.LNumber(%s))", luaStateShort, gopherLuaShort, varName)
}
func (fieldInt16 *luaFieldInt16) interfaceToLuaValue(varName string) string {
	return fmt.Sprintf("%s.LNumber(%s)", gopherLuaShort, varName)
}
func (fieldInt16 *luaFieldInt16) luaValueToInterface(varName string) string {
	return luaValueToInterface(varName, fieldInt16.typ, typeInt16, fieldInt16.typePackage, "LNumber")
}

// uint check
func (fieldUint *luaFieldUint) check(varName string, index int) string {
	return checkNumber(index, varName, fieldUint.typ, typeUInt, fieldUint.typePackage)
}

// uint push
func (fieldUint *luaFieldUint) push(varName string) string {
	return fmt.Sprintf("%s.Push(%s.LNumber(%s))", luaStateShort, gopherLuaShort, varName)
}
func (fieldUint *luaFieldUint) interfaceToLuaValue(varName string) string {
	return fmt.Sprintf("%s.LNumber(%s)", gopherLuaShort, varName)
}
func (fieldUint *luaFieldUint) luaValueToInterface(varName string) string {
	return luaValueToInterface(varName, fieldUint.typ, typeUInt, fieldUint.typePackage, "LNumber")
}

// int check
func (fieldInt *luaFieldInt) check(varName string, index int) string {
	c := fmt.Sprintf("%s.CheckInt(%d)", luaStateShort, index)
	if fieldInt.typ != typeInt {
		c = fmt.Sprintf("%s.%s(%s)", fieldInt.typePackage, fieldInt.typ, c)
	}
	c = fmt.Sprintf("%s = %s", varName, c)
	return c
}

// int push
func (fieldInt *luaFieldInt) push(varName string) string {
	return fmt.Sprintf("%s.Push(%s.LNumber(%s))", luaStateShort, gopherLuaShort, varName)
}
func (fieldInt *luaFieldInt) interfaceToLuaValue(varName string) string {
	return fmt.Sprintf("%s.LNumber(%s)", gopherLuaShort, varName)
}
func (fieldInt *luaFieldInt) luaValueToInterface(varName string) string {
	return luaValueToInterface(varName, fieldInt.typ, typeInt, fieldInt.typePackage, "LNumber")
}

// uint32 check
func (fieldUint32 *luaFieldUint32) check(varName string, index int) string {
	return checkNumber(index, varName, fieldUint32.typ, typeUInt32, fieldUint32.typePackage)
}

// uint32 push
func (fieldUint32 *luaFieldUint32) push(varName string) string {
	return fmt.Sprintf("%s.Push(%s.LNumber(%s))", luaStateShort, gopherLuaShort, varName)
}
func (fieldUint32 *luaFieldUint32) interfaceToLuaValue(varName string) string {
	return fmt.Sprintf("%s.LNumber(%s)", gopherLuaShort, varName)
}
func (fieldUint32 *luaFieldUint32) luaValueToInterface(varName string) string {
	return luaValueToInterface(varName, fieldUint32.typ, typeUInt32, fieldUint32.typePackage, "LNumber")
}

// int32 check
func (fieldInt32 *luaFieldInt32) check(varName string, index int) string {

	return checkNumber(index, varName, fieldInt32.typ, typeInt32, fieldInt32.typePackage)
}

// uint32 push
func (fieldInt32 *luaFieldInt32) push(varName string) string {
	return fmt.Sprintf("%s.Push(%s.LNumber(%s))", luaStateShort, gopherLuaShort, varName)
}
func (fieldInt32 *luaFieldInt32) interfaceToLuaValue(varName string) string {
	return fmt.Sprintf("%s.LNumber(%s)", gopherLuaShort, varName)
}
func (fieldInt32 *luaFieldInt32) luaValueToInterface(varName string) string {
	return luaValueToInterface(varName, fieldInt32.typ, typeInt32, fieldInt32.typePackage, "LNumber")
}

// uint64 check
func (fieldUint64 *luaFieldUint64) check(varName string, index int) string {
	return checkNumber(index, varName, fieldUint64.typ, typeUInt64, fieldUint64.typePackage)
}

// uint64 push
func (fieldUint64 *luaFieldUint64) push(varName string) string {
	return fmt.Sprintf("%s.Push(%s.LNumber(%s))", luaStateShort, gopherLuaShort, varName)
}
func (fieldUint64 *luaFieldUint64) interfaceToLuaValue(varName string) string {
	return fmt.Sprintf("%s.LNumber(%s)", gopherLuaShort, varName)
}
func (fieldUint64 *luaFieldUint64) luaValueToInterface(varName string) string {
	return luaValueToInterface(varName, fieldUint64.typ, typeUInt64, fieldUint64.typePackage, "LNumber")
}

// int64 check
func (fieldInt64 *luaFieldInt64) check(varName string, index int) string {
	c := fmt.Sprintf("%s.CheckInt64(%d)", luaStateShort, index)
	if fieldInt64.typ != typeInt64 {
		c = fmt.Sprintf("%s.%s(%s)", fieldInt64.typePackage, fieldInt64.typ, c)
	}
	c = fmt.Sprintf("%s = %s", varName, c)
	return c
}

// int64 push
func (fieldInt64 *luaFieldInt64) push(varName string) string {
	return fmt.Sprintf("%s.Push(%s.LNumber(%s))", luaStateShort, gopherLuaShort, varName)
}
func (fieldInt64 *luaFieldInt64) interfaceToLuaValue(varName string) string {
	return fmt.Sprintf("%s.LNumber(%s)", gopherLuaShort, varName)
}
func (fieldInt64 *luaFieldInt64) luaValueToInterface(varName string) string {
	return luaValueToInterface(varName, fieldInt64.typ, typeInt64, fieldInt64.typePackage, "LNumber")
}

// float32 check
func (fieldFloat32 *luaFieldFloat32) check(varName string, index int) string {
	return checkNumber(index, varName, fieldFloat32.typ, typeFloat32, fieldFloat32.typePackage)
}

// float32 push
func (fieldFloat32 *luaFieldFloat32) push(varName string) string {
	return fmt.Sprintf("%s.Push(%s.LNumber(%s))", luaStateShort, gopherLuaShort, varName)
}
func (fieldFloat32 *luaFieldFloat32) interfaceToLuaValue(varName string) string {
	return fmt.Sprintf("%s.LNumber(%s)", gopherLuaShort, varName)
}
func (fieldFloat32 *luaFieldFloat32) luaValueToInterface(varName string) string {
	return luaValueToInterface(varName, fieldFloat32.typ, typeFloat32, fieldFloat32.typePackage, "LNumber")
}

//float64 check
func (fieldFloat64 *luaFieldFloat64) check(varName string, index int) string {
	return checkNumber(index, varName, fieldFloat64.typ, typeFloat64, fieldFloat64.typePackage)
}

// float64 push
func (fieldFloat64 *luaFieldFloat64) push(varName string) string {
	return fmt.Sprintf("%s.Push(%s.LNumber(%s))", luaStateShort, gopherLuaShort, varName)
}

func (fieldFloat64 *luaFieldFloat64) luaValueToInterface(varName string) string {
	return luaValueToInterface(varName, fieldFloat64.typ, typeFloat64, fieldFloat64.typePackage, "LNumber")
}
func (fieldFloat64 *luaFieldFloat64) interfaceToLuaValue(varName string) string {
	return fmt.Sprintf("%s.LNumber(%s)", gopherLuaShort, varName)
}
