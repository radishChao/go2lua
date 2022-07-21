package main

import (
	"github.com/radishChao/go2lua/example/tolua/example1"
	"github.com/radishChao/go2lua/example/tolua/example2"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()
	err := example2.RegisterExampleToLua(L)
	if err != nil {
		panic(err)
		return
	}
	err = example1.RegisterExampleToLua(L)
	if err != nil {
		panic(err)
		return
	}

	err = L.DoFile("github.com/radishChao/go2lua/example/script/main.lua")
	if err != nil {
		panic(err)
		return
	}

}
