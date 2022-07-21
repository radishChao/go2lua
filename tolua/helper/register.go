package helper

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"reflect"
)

var (
	metatables map[*lua.LState]map[reflect.Type]string
)

// 注册go包
func registerPackage(L *lua.LState, goPackageName string) *lua.LTable {
	goPackage := L.GetGlobal(goPackageName)
	if goPackage != lua.LNil && goPackage.Type() != lua.LTTable {
		return nil
	}
	if goPackage == lua.LNil {
		goPackage = L.NewTable()
		L.SetGlobal(goPackageName, goPackage)
	}
	return goPackage.(*lua.LTable)
}

// RegisterStruct 注入结构体
func RegisterStruct(L *lua.LState, goPackageName string, typeMetatable string, i interface{}, constructor lua.LGFunction, methods map[string]lua.LGFunction) error {
	if L.GetTypeMetatable(typeMetatable) != lua.LNil {
		return fmt.Errorf("%s registed", typeMetatable)
	}

	mt := L.NewTypeMetatable(typeMetatable)
	if goPackageName != "" {
		registerPackage(L, goPackageName).RawSetString(typeMetatable, mt)
	} else {
		L.SetGlobal(typeMetatable, mt)
	}
	L.SetField(mt, "new", L.NewFunction(constructor))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), methods))

	if len(metatables) == 0 {
		metatables = map[*lua.LState]map[reflect.Type]string{}
	}
	m, ok := metatables[L]
	if !ok {
		m = map[reflect.Type]string{}
		metatables[L] = m
	}
	metatables[L][reflect.TypeOf(i)] = typeMetatable
	return nil
}

func RemoveLuaState(L *lua.LState) {
	delete(metatables, L)
}
