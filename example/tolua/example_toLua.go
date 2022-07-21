// Package tolua gen register toLua code with tag @lua from filegithub.com/radishChao/go2lua/example/go/example.go
// Package example 说明

package tolua

import (
	"github.com/radishChao/go2lua/example/go"
	example2 "github.com/radishChao/go2lua/example/go/example2"
	example2Tolua "github.com/radishChao/go2lua/example/tolua/example2"
	luaHelper "github.com/radishChao/go2lua/tolua/helper"
	lua "github.com/yuin/gopher-lua"
)

var (
	mapExportToLua3ToLuaMethods = map[string]lua.LGFunction{
		"S1":  registerExportToLua3S1,
		"S2":  registerExportToLua3S2,
		"S3":  registerExportToLua3S3,
		"S4":  registerExportToLua3S4,
		"S5":  registerExportToLua3S5,
		"S6":  registerExportToLua3S6,
		"S7":  registerExportToLua3S7,
		"S8":  registerExportToLua3S8,
		"S9":  registerExportToLua3S9,
		"S10": registerExportToLua3S10,
		"S11": registerExportToLua3S11,
		"S12": registerExportToLua3S12,
		"S13": registerExportToLua3S13,
		"S14": registerExportToLua3S14,
		"S15": registerExportToLua3S15,
		"S17": registerExportToLua3S17,
		"S18": registerExportToLua3S18,
		"A1":  registerExportToLua3A1,
		"A2":  registerExportToLua3A2,
		"A3":  registerExportToLua3A3,
		"A4":  registerExportToLua3A4,
		"A5":  registerExportToLua3A5,
		"A6":  registerExportToLua3A6,
		"A7":  registerExportToLua3A7,
		"A8":  registerExportToLua3A8,
		"A9":  registerExportToLua3A9,
		"A10": registerExportToLua3A10,
		"A11": registerExportToLua3A11,
		"A12": registerExportToLua3A12,
		"A13": registerExportToLua3A13,
		"A14": registerExportToLua3A14,
		"A15": registerExportToLua3A15,
		"A16": registerExportToLua3A16,
		"A17": registerExportToLua3A17,
		"A18": registerExportToLua3A18,
		"A19": registerExportToLua3A19,
		"A20": registerExportToLua3A20,
		"A21": registerExportToLua3A21,
	}
	mapMyStructToLuaMethods = map[string]lua.LGFunction{
		"Float": registerMyStructFloat,
	}
)

// RegisterExampleToLua call this func register
func RegisterExampleToLua(L *lua.LState) error {
	var err error
	err = luaHelper.RegisterStruct(L, "example", "ExportToLua3", (*example.ExportToLua3)(nil), newExampleExportToLua3ToLua, mapExportToLua3ToLuaMethods)
	if err != nil {
		return err
	}
	err = luaHelper.RegisterStruct(L, "example", "MyStruct", (*example.MyStruct)(nil), newExampleMyStructToLua, mapMyStructToLuaMethods)
	if err != nil {
		return err
	}
	return nil
}
func GenUserDataExampleExportToLua3(L *lua.LState, exportToLua3 *example.ExportToLua3) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = exportToLua3
	L.SetMetatable(ud, L.GetTypeMetatable("ExportToLua3"))
	return ud
}
func PushExampleExportToLua3ToLua(L *lua.LState, exportToLua3 *example.ExportToLua3) {
	L.Push(GenUserDataExampleExportToLua3(L, exportToLua3))
}
func CheckExampleExportToLua3ToLua(L *lua.LState, index int) *example.ExportToLua3 {
	ud := L.CheckUserData(index)
	if v, ok := ud.Value.(*example.ExportToLua3); ok {
		return v
	}
	L.ArgError(index, "example.ExportToLua3 expected")
	return nil
}
func newExampleExportToLua3ToLua(L *lua.LState) int {
	PushExampleExportToLua3ToLua(L, &example.ExportToLua3{})
	return 1
}

// registerExportToLua3S1 ExportToLua3.S1 get set
func registerExportToLua3S1(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.S1 = exportToLua3.S1[len(exportToLua3.S1):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua3.S1 = append(exportToLua3.S1, bool(luaT.RawGetInt(i+1).(lua.LBool)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua3.S1); i++ {
		luaT.Append(lua.LBool(exportToLua3.S1[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua3S2 ExportToLua3.S2 get set
func registerExportToLua3S2(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.S2 = exportToLua3.S2[len(exportToLua3.S2):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua3.S2 = append(exportToLua3.S2, byte(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua3.S2); i++ {
		luaT.Append(lua.LNumber(exportToLua3.S2[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua3S3 ExportToLua3.S3 get set
func registerExportToLua3S3(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.S3 = exportToLua3.S3[len(exportToLua3.S3):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua3.S3 = append(exportToLua3.S3, string(luaT.RawGetInt(i+1).(lua.LString)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua3.S3); i++ {
		luaT.Append(lua.LString(exportToLua3.S3[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua3S4 ExportToLua3.S4 get set
func registerExportToLua3S4(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.S4 = exportToLua3.S4[len(exportToLua3.S4):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua3.S4 = append(exportToLua3.S4, uint8(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua3.S4); i++ {
		luaT.Append(lua.LNumber(exportToLua3.S4[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua3S5 ExportToLua3.S5 get set
func registerExportToLua3S5(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.S5 = exportToLua3.S5[len(exportToLua3.S5):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua3.S5 = append(exportToLua3.S5, int8(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua3.S5); i++ {
		luaT.Append(lua.LNumber(exportToLua3.S5[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua3S6 ExportToLua3.S6 get set
func registerExportToLua3S6(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.S6 = exportToLua3.S6[len(exportToLua3.S6):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua3.S6 = append(exportToLua3.S6, uint16(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua3.S6); i++ {
		luaT.Append(lua.LNumber(exportToLua3.S6[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua3S7 ExportToLua3.S7 get set
func registerExportToLua3S7(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.S7 = exportToLua3.S7[len(exportToLua3.S7):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua3.S7 = append(exportToLua3.S7, int16(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua3.S7); i++ {
		luaT.Append(lua.LNumber(exportToLua3.S7[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua3S8 ExportToLua3.S8 get set
func registerExportToLua3S8(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.S8 = exportToLua3.S8[len(exportToLua3.S8):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua3.S8 = append(exportToLua3.S8, uint32(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua3.S8); i++ {
		luaT.Append(lua.LNumber(exportToLua3.S8[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua3S9 ExportToLua3.S9 get set
func registerExportToLua3S9(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.S9 = exportToLua3.S9[len(exportToLua3.S9):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua3.S9 = append(exportToLua3.S9, int32(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua3.S9); i++ {
		luaT.Append(lua.LNumber(exportToLua3.S9[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua3S10 ExportToLua3.S10 get set
func registerExportToLua3S10(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.S10 = exportToLua3.S10[len(exportToLua3.S10):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua3.S10 = append(exportToLua3.S10, uint(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua3.S10); i++ {
		luaT.Append(lua.LNumber(exportToLua3.S10[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua3S11 ExportToLua3.S11 get set
func registerExportToLua3S11(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.S11 = exportToLua3.S11[len(exportToLua3.S11):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua3.S11 = append(exportToLua3.S11, int(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua3.S11); i++ {
		luaT.Append(lua.LNumber(exportToLua3.S11[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua3S12 ExportToLua3.S12 get set
func registerExportToLua3S12(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.S12 = exportToLua3.S12[len(exportToLua3.S12):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua3.S12 = append(exportToLua3.S12, uint64(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua3.S12); i++ {
		luaT.Append(lua.LNumber(exportToLua3.S12[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua3S13 ExportToLua3.S13 get set
func registerExportToLua3S13(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.S13 = exportToLua3.S13[len(exportToLua3.S13):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua3.S13 = append(exportToLua3.S13, int64(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua3.S13); i++ {
		luaT.Append(lua.LNumber(exportToLua3.S13[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua3S14 ExportToLua3.S14 get set
func registerExportToLua3S14(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.S14 = exportToLua3.S14[len(exportToLua3.S14):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua3.S14 = append(exportToLua3.S14, float32(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua3.S14); i++ {
		luaT.Append(lua.LNumber(exportToLua3.S14[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua3S15 ExportToLua3.S15 get set
func registerExportToLua3S15(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.S15 = exportToLua3.S15[len(exportToLua3.S15):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua3.S15 = append(exportToLua3.S15, float64(luaT.RawGetInt(i+1).(lua.LNumber)))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua3.S15); i++ {
		luaT.Append(lua.LNumber(exportToLua3.S15[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua3S17 ExportToLua3.S17 get set
func registerExportToLua3S17(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.S17 = exportToLua3.S17[len(exportToLua3.S17):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua3.S17 = append(exportToLua3.S17, luaT.RawGetInt(i+1).(*lua.LUserData).Value.(*example.MyStruct))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua3.S17); i++ {
		luaT.Append(GenUserDataExampleMyStruct(L, exportToLua3.S17[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua3S18 ExportToLua3.S18 get set
func registerExportToLua3S18(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.S18 = exportToLua3.S18[len(exportToLua3.S18):1]
		luaT := L.CheckTable(2)
		luaTLen := luaT.Len()
		for i := 0; i < luaTLen; i++ {
			exportToLua3.S18 = append(exportToLua3.S18, luaT.RawGetInt(i+1).(*lua.LUserData).Value.(*example2.ExportToLua2))
		}

		return 0
	}
	luaT := L.NewTable()
	for i := 0; i < len(exportToLua3.S18); i++ {
		luaT.Append(example2Tolua.GenUserDataExample2ExportToLua2(L, exportToLua3.S18[i]))
	}
	L.Push(luaT)

	return 1
}

// registerExportToLua3A1 ExportToLua3.A1 get set
func registerExportToLua3A1(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A1 = int8(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua3.A1))
	return 1
}

// registerExportToLua3A2 ExportToLua3.A2 get set
func registerExportToLua3A2(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A2 = uint8(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua3.A2))
	return 1
}

// registerExportToLua3A3 ExportToLua3.A3 get set
func registerExportToLua3A3(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A3 = int16(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua3.A3))
	return 1
}

// registerExportToLua3A4 ExportToLua3.A4 get set
func registerExportToLua3A4(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A4 = int16(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua3.A4))
	return 1
}

// registerExportToLua3A5 ExportToLua3.A5 get set
func registerExportToLua3A5(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A5 = L.CheckInt(2)
		return 0
	}
	L.Push(lua.LNumber(exportToLua3.A5))
	return 1
}

// registerExportToLua3A6 ExportToLua3.A6 get set
func registerExportToLua3A6(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A6 = uint(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua3.A6))
	return 1
}

// registerExportToLua3A7 ExportToLua3.A7 get set
func registerExportToLua3A7(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A7 = int32(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua3.A7))
	return 1
}

// registerExportToLua3A8 ExportToLua3.A8 get set
func registerExportToLua3A8(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A8 = int32(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua3.A8))
	return 1
}

// registerExportToLua3A9 ExportToLua3.A9 get set
func registerExportToLua3A9(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A9 = L.CheckInt64(2)
		return 0
	}
	L.Push(lua.LNumber(exportToLua3.A9))
	return 1
}

// registerExportToLua3A10 ExportToLua3.A10 get set
func registerExportToLua3A10(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A10 = L.CheckInt64(2)
		return 0
	}
	L.Push(lua.LNumber(exportToLua3.A10))
	return 1
}

// registerExportToLua3A11 ExportToLua3.A11 get set
func registerExportToLua3A11(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A11 = float32(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua3.A11))
	return 1
}

// registerExportToLua3A12 ExportToLua3.A12 get set
func registerExportToLua3A12(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A12 = float64(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua3.A12))
	return 1
}

// registerExportToLua3A13 ExportToLua3.A13 get set
func registerExportToLua3A13(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A13 = byte(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua3.A13))
	return 1
}

// registerExportToLua3A14 ExportToLua3.A14 get set
func registerExportToLua3A14(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A14 = luaHelper.Check(L, exportToLua3.A14, 2)
		return 0
	}
	luaHelper.Push(L, exportToLua3.A14)
	return 1
}

// registerExportToLua3A15 ExportToLua3.A15 get set
func registerExportToLua3A15(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A15 = make(map[string]*example.MyStruct)
		t := L.CheckTable(2)
		t.ForEach(func(key lua.LValue, value lua.LValue) {
			exportToLua3.A15[string(key.(lua.LString))] = value.(*lua.LUserData).Value.(*example.MyStruct)
		})

		return 0
	}
	t := L.NewTable()
	for k, v := range exportToLua3.A15 {
		t.RawSet(lua.LString(k), GenUserDataExampleMyStruct(L, v))
	}

	return 1
}

// registerExportToLua3A16 ExportToLua3.A16 get set
func registerExportToLua3A16(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A16 = make(map[string]int)
		t := L.CheckTable(2)
		t.ForEach(func(key lua.LValue, value lua.LValue) {
			exportToLua3.A16[string(key.(lua.LString))] = int(value.(lua.LNumber))
		})

		return 0
	}
	t := L.NewTable()
	for k, v := range exportToLua3.A16 {
		t.RawSet(lua.LString(k), lua.LNumber(v))
	}

	return 1
}

// registerExportToLua3A17 ExportToLua3.A17 get set
func registerExportToLua3A17(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A17 = example2Tolua.CheckExample2ExportToLua2ToLua(L, 2)
		return 0
	}
	example2Tolua.PushExample2ExportToLua2ToLua(L, exportToLua3.A17)
	return 1
}

// registerExportToLua3A18 ExportToLua3.A18 get set
func registerExportToLua3A18(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A18 = example2.MyInt(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua3.A18))
	return 1
}

// registerExportToLua3A19 ExportToLua3.A19 get set
func registerExportToLua3A19(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A19 = example.MyString(L.CheckString(2))
		return 0
	}
	L.Push(lua.LString(exportToLua3.A19))
	return 1
}

// registerExportToLua3A20 ExportToLua3.A20 get set
func registerExportToLua3A20(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A20 = CheckExampleMyStructToLua(L, 2)
		return 0
	}
	PushExampleMyStructToLua(L, exportToLua3.A20)
	return 1
}

// registerExportToLua3A21 ExportToLua3.A21 get set
func registerExportToLua3A21(L *lua.LState) int {
	exportToLua3 := CheckExampleExportToLua3ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua3.A21 = example.MyInt(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua3.A21))
	return 1
}
func GenUserDataExampleMyStruct(L *lua.LState, myStruct *example.MyStruct) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = myStruct
	L.SetMetatable(ud, L.GetTypeMetatable("MyStruct"))
	return ud
}
func PushExampleMyStructToLua(L *lua.LState, myStruct *example.MyStruct) {
	L.Push(GenUserDataExampleMyStruct(L, myStruct))
}
func CheckExampleMyStructToLua(L *lua.LState, index int) *example.MyStruct {
	ud := L.CheckUserData(index)
	if v, ok := ud.Value.(*example.MyStruct); ok {
		return v
	}
	L.ArgError(index, "example.MyStruct expected")
	return nil
}
func newExampleMyStructToLua(L *lua.LState) int {
	PushExampleMyStructToLua(L, &example.MyStruct{})
	return 1
}

// registerMyStructFloat MyStruct.Float get set
func registerMyStructFloat(L *lua.LState) int {
	myStruct := CheckExampleMyStructToLua(L, 1)
	if L.GetTop() == 2 {
		myStruct.Float = float32(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(myStruct.Float))
	return 1
}
