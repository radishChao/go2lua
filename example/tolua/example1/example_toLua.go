// Package example1 gen register toLua code with tag @lua from file github.com/radishChao/go2lua/example/go/example1/example.go
// Package example 说明

package example1

import (
	"github.com/radishChao/go2lua/example/go/example1"
	example2 "github.com/radishChao/go2lua/example/go/example2"
	example2Tolua "github.com/radishChao/go2lua/example/tolua/example2"
	luaHelper "github.com/radishChao/go2lua/tolua/helper"
	lua "github.com/yuin/gopher-lua"
)

var (
	mapExportToLua1ToLuaMethods = map[string]lua.LGFunction{
		"S1":  registerExportToLua1S1,
		"S2":  registerExportToLua1S2,
		"S3":  registerExportToLua1S3,
		"S4":  registerExportToLua1S4,
		"S5":  registerExportToLua1S5,
		"S6":  registerExportToLua1S6,
		"S7":  registerExportToLua1S7,
		"S8":  registerExportToLua1S8,
		"S9":  registerExportToLua1S9,
		"S10": registerExportToLua1S10,
		"S11": registerExportToLua1S11,
		"S12": registerExportToLua1S12,
		"S13": registerExportToLua1S13,
		"S14": registerExportToLua1S14,
		"S15": registerExportToLua1S15,
		"S17": registerExportToLua1S17,
		"S18": registerExportToLua1S18,
		"A1":  registerExportToLua1A1,
		"A2":  registerExportToLua1A2,
		"A3":  registerExportToLua1A3,
		"A4":  registerExportToLua1A4,
		"A5":  registerExportToLua1A5,
		"A6":  registerExportToLua1A6,
		"A7":  registerExportToLua1A7,
		"A8":  registerExportToLua1A8,
		"A9":  registerExportToLua1A9,
		"A10": registerExportToLua1A10,
		"A11": registerExportToLua1A11,
		"A12": registerExportToLua1A12,
		"A13": registerExportToLua1A13,
		"A14": registerExportToLua1A14,
		"A15": registerExportToLua1A15,
		"A16": registerExportToLua1A16,
		"A17": registerExportToLua1A17,
		"A18": registerExportToLua1A18,
		"A19": registerExportToLua1A19,
		"A20": registerExportToLua1A20,
		"A21": registerExportToLua1A21,
	}
	mapMyStructToLuaMethods = map[string]lua.LGFunction{
		"Float": registerMyStructFloat,
	}
)

// RegisterExampleToLua call this func register
func RegisterExampleToLua(L *lua.LState) error {
	var err error
	err = luaHelper.RegisterStruct(L, "example1", "ExportToLua1", (*example1.ExportToLua1)(nil), newExample1ExportToLua1ToLua, mapExportToLua1ToLuaMethods)
	if err != nil {
		return err
	}
	err = luaHelper.RegisterStruct(L, "example1", "MyStruct", (*example1.MyStruct)(nil), newExample1MyStructToLua, mapMyStructToLuaMethods)
	if err != nil {
		return err
	}
	return nil
}
func GenUserDataExample1ExportToLua1(L *lua.LState, exportToLua1 *example1.ExportToLua1) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = exportToLua1
	L.SetMetatable(ud, L.GetTypeMetatable("ExportToLua1"))
	return ud
}
func PushExample1ExportToLua1ToLua(L *lua.LState, exportToLua1 *example1.ExportToLua1) {
	L.Push(GenUserDataExample1ExportToLua1(L, exportToLua1))
}
func CheckExample1ExportToLua1ToLua(L *lua.LState, index int) *example1.ExportToLua1 {
	ud := L.CheckUserData(index)
	if v, ok := ud.Value.(*example1.ExportToLua1); ok {
		return v
	}
	L.ArgError(index, "example1.ExportToLua1 expected")
	return nil
}
func newExample1ExportToLua1ToLua(L *lua.LState) int {
	PushExample1ExportToLua1ToLua(L, &example1.ExportToLua1{})
	return 1
}

// registerExportToLua1S1 ExportToLua1.S1 get set
func registerExportToLua1S1(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.S1 = exportToLua1.S1[len(exportToLua1.S1):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua1.S1 = append(exportToLua1.S1, bool(luaT.RawGetInt(i+1).(lua.LBool)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua1.S1); i++ {
		luaT.Append(lua.LBool(exportToLua1.S1[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua1S2 ExportToLua1.S2 get set
func registerExportToLua1S2(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.S2 = exportToLua1.S2[len(exportToLua1.S2):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua1.S2 = append(exportToLua1.S2, byte(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua1.S2); i++ {
		luaT.Append(lua.LNumber(exportToLua1.S2[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua1S3 ExportToLua1.S3 get set
func registerExportToLua1S3(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.S3 = exportToLua1.S3[len(exportToLua1.S3):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua1.S3 = append(exportToLua1.S3, string(luaT.RawGetInt(i+1).(lua.LString)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua1.S3); i++ {
		luaT.Append(lua.LString(exportToLua1.S3[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua1S4 ExportToLua1.S4 get set
func registerExportToLua1S4(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.S4 = exportToLua1.S4[len(exportToLua1.S4):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua1.S4 = append(exportToLua1.S4, uint8(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua1.S4); i++ {
		luaT.Append(lua.LNumber(exportToLua1.S4[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua1S5 ExportToLua1.S5 get set
func registerExportToLua1S5(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.S5 = exportToLua1.S5[len(exportToLua1.S5):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua1.S5 = append(exportToLua1.S5, int8(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua1.S5); i++ {
		luaT.Append(lua.LNumber(exportToLua1.S5[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua1S6 ExportToLua1.S6 get set
func registerExportToLua1S6(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.S6 = exportToLua1.S6[len(exportToLua1.S6):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua1.S6 = append(exportToLua1.S6, uint16(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua1.S6); i++ {
		luaT.Append(lua.LNumber(exportToLua1.S6[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua1S7 ExportToLua1.S7 get set
func registerExportToLua1S7(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.S7 = exportToLua1.S7[len(exportToLua1.S7):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua1.S7 = append(exportToLua1.S7, int16(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua1.S7); i++ {
		luaT.Append(lua.LNumber(exportToLua1.S7[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua1S8 ExportToLua1.S8 get set
func registerExportToLua1S8(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.S8 = exportToLua1.S8[len(exportToLua1.S8):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua1.S8 = append(exportToLua1.S8, uint32(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua1.S8); i++ {
		luaT.Append(lua.LNumber(exportToLua1.S8[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua1S9 ExportToLua1.S9 get set
func registerExportToLua1S9(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.S9 = exportToLua1.S9[len(exportToLua1.S9):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua1.S9 = append(exportToLua1.S9, int32(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua1.S9); i++ {
		luaT.Append(lua.LNumber(exportToLua1.S9[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua1S10 ExportToLua1.S10 get set
func registerExportToLua1S10(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.S10 = exportToLua1.S10[len(exportToLua1.S10):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua1.S10 = append(exportToLua1.S10, uint(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua1.S10); i++ {
		luaT.Append(lua.LNumber(exportToLua1.S10[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua1S11 ExportToLua1.S11 get set
func registerExportToLua1S11(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.S11 = exportToLua1.S11[len(exportToLua1.S11):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua1.S11 = append(exportToLua1.S11, int(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua1.S11); i++ {
		luaT.Append(lua.LNumber(exportToLua1.S11[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua1S12 ExportToLua1.S12 get set
func registerExportToLua1S12(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.S12 = exportToLua1.S12[len(exportToLua1.S12):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua1.S12 = append(exportToLua1.S12, uint64(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua1.S12); i++ {
		luaT.Append(lua.LNumber(exportToLua1.S12[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua1S13 ExportToLua1.S13 get set
func registerExportToLua1S13(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.S13 = exportToLua1.S13[len(exportToLua1.S13):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua1.S13 = append(exportToLua1.S13, int64(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua1.S13); i++ {
		luaT.Append(lua.LNumber(exportToLua1.S13[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua1S14 ExportToLua1.S14 get set
func registerExportToLua1S14(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.S14 = exportToLua1.S14[len(exportToLua1.S14):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua1.S14 = append(exportToLua1.S14, float32(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua1.S14); i++ {
		luaT.Append(lua.LNumber(exportToLua1.S14[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua1S15 ExportToLua1.S15 get set
func registerExportToLua1S15(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.S15 = exportToLua1.S15[len(exportToLua1.S15):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua1.S15 = append(exportToLua1.S15, float64(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua1.S15); i++ {
		luaT.Append(lua.LNumber(exportToLua1.S15[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua1S17 ExportToLua1.S17 get set
func registerExportToLua1S17(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.S17 = exportToLua1.S17[len(exportToLua1.S17):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua1.S17 = append(exportToLua1.S17, luaT.RawGetInt(i+1).(*lua.LUserData).Value.(*example1.MyStruct))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua1.S17); i++ {
		luaT.Append(GenUserDataExample1MyStruct(L, exportToLua1.S17[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua1S18 ExportToLua1.S18 get set
func registerExportToLua1S18(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.S18 = exportToLua1.S18[len(exportToLua1.S18):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua1.S18 = append(exportToLua1.S18, luaT.RawGetInt(i+1).(*lua.LUserData).Value.(*example2.ExportToLua2))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua1.S18); i++ {
		luaT.Append(example2Tolua.GenUserDataExample2ExportToLua2(L, exportToLua1.S18[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua1A1 ExportToLua1.A1 get set
func registerExportToLua1A1(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A1 = int8(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua1.A1))
	return 1
}

// registerExportToLua1A2 ExportToLua1.A2 get set
func registerExportToLua1A2(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A2 = uint8(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua1.A2))
	return 1
}

// registerExportToLua1A3 ExportToLua1.A3 get set
func registerExportToLua1A3(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A3 = int16(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua1.A3))
	return 1
}

// registerExportToLua1A4 ExportToLua1.A4 get set
func registerExportToLua1A4(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A4 = int16(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua1.A4))
	return 1
}

// registerExportToLua1A5 ExportToLua1.A5 get set
func registerExportToLua1A5(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A5 = L.CheckInt(2)
		return 0
	}
	L.Push(lua.LNumber(exportToLua1.A5))
	return 1
}

// registerExportToLua1A6 ExportToLua1.A6 get set
func registerExportToLua1A6(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A6 = uint(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua1.A6))
	return 1
}

// registerExportToLua1A7 ExportToLua1.A7 get set
func registerExportToLua1A7(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A7 = int32(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua1.A7))
	return 1
}

// registerExportToLua1A8 ExportToLua1.A8 get set
func registerExportToLua1A8(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A8 = int32(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua1.A8))
	return 1
}

// registerExportToLua1A9 ExportToLua1.A9 get set
func registerExportToLua1A9(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A9 = L.CheckInt64(2)
		return 0
	}
	L.Push(lua.LNumber(exportToLua1.A9))
	return 1
}

// registerExportToLua1A10 ExportToLua1.A10 get set
func registerExportToLua1A10(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A10 = L.CheckInt64(2)
		return 0
	}
	L.Push(lua.LNumber(exportToLua1.A10))
	return 1
}

// registerExportToLua1A11 ExportToLua1.A11 get set
func registerExportToLua1A11(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A11 = float32(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua1.A11))
	return 1
}

// registerExportToLua1A12 ExportToLua1.A12 get set
func registerExportToLua1A12(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A12 = float64(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua1.A12))
	return 1
}

// registerExportToLua1A13 ExportToLua1.A13 get set
func registerExportToLua1A13(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A13 = byte(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua1.A13))
	return 1
}

// registerExportToLua1A14 ExportToLua1.A14 get set
func registerExportToLua1A14(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A14 = luaHelper.Check(L, exportToLua1.A14, 2)
		return 0
	}
	luaHelper.Push(L, exportToLua1.A14)
	return 1
}

// registerExportToLua1A15 ExportToLua1.A15 get set
func registerExportToLua1A15(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A15 = make(map[string]*example1.MyStruct)
		t := L.CheckTable(2)
		t.ForEach(func(key lua.LValue, value lua.LValue) {
			exportToLua1.A15[string(key.(lua.LString))] = value.(*lua.LUserData).Value.(*example1.MyStruct)
		})

		return 0
	}
	t := L.NewTable()
	for k, v := range exportToLua1.A15 {
		t.RawSet(lua.LString(k), GenUserDataExample1MyStruct(L, v))
	}

	return 1
}

// registerExportToLua1A16 ExportToLua1.A16 get set
func registerExportToLua1A16(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A16 = make(map[string]int)
		t := L.CheckTable(2)
		t.ForEach(func(key lua.LValue, value lua.LValue) {
			exportToLua1.A16[string(key.(lua.LString))] = int(value.(lua.LNumber))
		})

		return 0
	}
	t := L.NewTable()
	for k, v := range exportToLua1.A16 {
		t.RawSet(lua.LString(k), lua.LNumber(v))
	}

	return 1
}

// registerExportToLua1A17 ExportToLua1.A17 get set
func registerExportToLua1A17(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A17 = example2Tolua.CheckExample2ExportToLua2ToLua(L, 2)
		return 0
	}
	example2Tolua.PushExample2ExportToLua2ToLua(L, exportToLua1.A17)
	return 1
}

// registerExportToLua1A18 ExportToLua1.A18 get set
func registerExportToLua1A18(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A18 = example2.MyInt(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua1.A18))
	return 1
}

// registerExportToLua1A19 ExportToLua1.A19 get set
func registerExportToLua1A19(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A19 = example1.MyString(L.CheckString(2))
		return 0
	}
	L.Push(lua.LString(exportToLua1.A19))
	return 1
}

// registerExportToLua1A20 ExportToLua1.A20 get set
func registerExportToLua1A20(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A20 = CheckExample1MyStructToLua(L, 2)
		return 0
	}
	PushExample1MyStructToLua(L, exportToLua1.A20)
	return 1
}

// registerExportToLua1A21 ExportToLua1.A21 get set
func registerExportToLua1A21(L *lua.LState) int {
	exportToLua1 := CheckExample1ExportToLua1ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua1.A21 = example1.MyInt(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua1.A21))
	return 1
}
func GenUserDataExample1MyStruct(L *lua.LState, myStruct *example1.MyStruct) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = myStruct
	L.SetMetatable(ud, L.GetTypeMetatable("MyStruct"))
	return ud
}
func PushExample1MyStructToLua(L *lua.LState, myStruct *example1.MyStruct) {
	L.Push(GenUserDataExample1MyStruct(L, myStruct))
}
func CheckExample1MyStructToLua(L *lua.LState, index int) *example1.MyStruct {
	ud := L.CheckUserData(index)
	if v, ok := ud.Value.(*example1.MyStruct); ok {
		return v
	}
	L.ArgError(index, "example1.MyStruct expected")
	return nil
}
func newExample1MyStructToLua(L *lua.LState) int {
	PushExample1MyStructToLua(L, &example1.MyStruct{})
	return 1
}

// registerMyStructFloat MyStruct.Float get set
func registerMyStructFloat(L *lua.LState) int {
	myStruct := CheckExample1MyStructToLua(L, 1)
	if L.GetTop() == 2 {
		myStruct.Float = float32(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(myStruct.Float))
	return 1
}
