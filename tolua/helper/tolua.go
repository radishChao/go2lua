package helper

import (
	lua "github.com/yuin/gopher-lua"
	"reflect"
)

func Push(L *lua.LState, i interface{}) {
	L.Push(InterfaceToLuaValue(L, i))
}

func Check(L *lua.LState, i interface{}, index int) interface{} {
	return LuaValueToInterface(L.CheckAny(index), i)
}

// LuaValueToInterface luaValue转换为interface
func LuaValueToInterface(lValue lua.LValue, i ...interface{}) interface{} {
	switch lValue.Type() {
	case lua.LTNil:
		return nil
	case lua.LTBool:
		return bool(lValue.(lua.LBool))
	case lua.LTChannel:
		return lValue.(lua.LChannel)
	case lua.LTNumber:
		return luaNumberToInterfaceWithType(lValue.(lua.LNumber), reflect.TypeOf(i))
	case lua.LTString:
		return string(lValue.(lua.LString))
	case lua.LTTable:
		return luaTableToInterface(lValue.(*lua.LTable))
	case lua.LTUserData:
		return lValue.(*lua.LUserData).Value
	default:
		return nil
	}
}

func luaNumberToInterfaceWithType(n lua.LNumber, t reflect.Type) interface{} {
	switch t.Kind() {
	case reflect.Int:
		return int(n)
	case reflect.Uint:
		return uint(n)
	case reflect.Int8:
		return int8(n)
	case reflect.Uint8:
		return uint8(n)
	case reflect.Int16:
		return int16(n)
	case reflect.Uint16:
		return uint16(n)
	case reflect.Int32:
		return int32(n)
	case reflect.Uint32:
		return uint32(n)
	case reflect.Int64:
		return int64(n)
	case reflect.Uint64:
		return uint64(n)
	case reflect.Float32:
		return float32(n)
	case reflect.Float64:
		return float64(n)
	}
	return 0
}

// lua table转换
// lua中的table可能是 go中的map或者slice
func luaTableToInterface(table *lua.LTable) interface{} {
	var slice []interface{}
	var m map[interface{}]interface{}
	isSlice := table.Len() != 0
	if isSlice {
		slice = []interface{}{}
	} else {
		m = map[interface{}]interface{}{}
	}
	table.ForEach(func(key lua.LValue, Value lua.LValue) {
		if isSlice {
			slice = append(slice, LuaValueToInterface(Value))
		} else {
			m[LuaValueToInterface(key)] = LuaValueToInterface(Value)
		}
	})

	if isSlice {
		return slice
	}
	if len(m) == 0 {
		return []interface{}{}
	}
	return m
}

// InterfaceToLuaValue interface转化为luaValue
func InterfaceToLuaValue(L *lua.LState, i interface{}) lua.LValue {
	kind := reflect.ValueOf(i).Kind()
	switch kind {
	case reflect.Int:
		return lua.LNumber(i.(int))
	case reflect.Uint:
		return lua.LNumber(i.(uint))
	case reflect.String:
		return lua.LString(i.(string))
	case reflect.Map:
		return goMapToLuaTable(L, i)
	case reflect.Slice:
		return goSliceToLuaTable(L, i)
	case reflect.Pointer:
		return interfaceToLuaUserData(L, i)
	case reflect.Int8:
		return lua.LNumber(i.(int8))
	case reflect.Uint8:
		return lua.LNumber(i.(uint8))
	case reflect.Int16:
		return lua.LNumber(i.(int16))
	case reflect.Uint16:
		return lua.LNumber(i.(uint16))
	case reflect.Int32:
		return lua.LNumber(i.(int32))
	case reflect.Uint32:
		return lua.LNumber(i.(uint32))
	case reflect.Int64:
		return lua.LNumber(i.(int64))
	case reflect.Uint64:
		return lua.LNumber(i.(uint16))
	case reflect.Float32:
		return lua.LNumber(i.(float32))
	case reflect.Float64:
		return lua.LNumber(i.(float64))
	}
	return lua.LNil
}

// go map转换为lua的table
func goMapToLuaTable(L *lua.LState, i interface{}) *lua.LTable {
	t := L.NewTable()
	iter := reflect.ValueOf(i).MapRange()
	for iter.Next() {
		t.RawSet(InterfaceToLuaValue(L, iter.Key().Interface()), InterfaceToLuaValue(L, iter.Value().Interface()))
	}
	return t
}

// go slice转换为lua的table
func goSliceToLuaTable(L *lua.LState, i interface{}) *lua.LTable {
	t := L.NewTable()
	v := reflect.ValueOf(i)
	size := v.Len()
	for j := 0; j < size; j++ {
		t.Append(InterfaceToLuaValue(L, v.Index(j).Interface()))
	}
	return t
}

func interfaceToLuaUserData(L *lua.LState, i interface{}) lua.LValue {
	m, ok := metatables[L]
	if !ok {
		return lua.LNil
	}
	typeMetatable, ok := m[reflect.TypeOf(i)]
	if !ok {
		return lua.LNil
	}
	pLValue := L.NewUserData()
	pLValue.Value = i
	L.SetMetatable(pLValue, L.GetTypeMetatable(typeMetatable))
	return pLValue
}
