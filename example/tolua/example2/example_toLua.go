// Package example2 gen register toLua code with tag @lua from file github.com/radishChao/go2lua/example/go/example2/example.go
package example2

import (
	"github.com/radishChao/go2lua/example/go/example2"
	luaHelper "github.com/radishChao/go2lua/tolua/helper"
	lua "github.com/yuin/gopher-lua"
)

var (
	mapExportToLua2ToLuaMethods = map[string]lua.LGFunction{
		"I": registerExportToLua2I,
	}
)

// RegisterExampleToLua call this func register
func RegisterExampleToLua(L *lua.LState) error {
	var err error
	err = luaHelper.RegisterStruct(L, "example2", "ExportToLua2", (*example2.ExportToLua2)(nil), newExample2ExportToLua2ToLua, mapExportToLua2ToLuaMethods)
	if err != nil {
		return err
	}
	return nil
}
func GenUserDataExample2ExportToLua2(L *lua.LState, exportToLua2 *example2.ExportToLua2) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = exportToLua2
	L.SetMetatable(ud, L.GetTypeMetatable("ExportToLua2"))
	return ud
}
func PushExample2ExportToLua2ToLua(L *lua.LState, exportToLua2 *example2.ExportToLua2) {
	L.Push(GenUserDataExample2ExportToLua2(L, exportToLua2))
}
func CheckExample2ExportToLua2ToLua(L *lua.LState, index int) *example2.ExportToLua2 {
	ud := L.CheckUserData(index)
	if v, ok := ud.Value.(*example2.ExportToLua2); ok {
		return v
	}
	L.ArgError(index, "example2.ExportToLua2 expected")
	return nil
}
func newExample2ExportToLua2ToLua(L *lua.LState) int {
	PushExample2ExportToLua2ToLua(L, &example2.ExportToLua2{})
	return 1
}

// registerExportToLua2I ExportToLua2.I get set
func registerExportToLua2I(L *lua.LState) int {
	exportToLua2 := CheckExample2ExportToLua2ToLua(L, 1)
	if L.GetTop() == 2 {
		exportToLua2.I = example2.MyInt(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(exportToLua2.I))
	return 1
}
