// Package example 说明
// @lua-package example2 example2Tolua github.com/radishChao/go2lua/example/tolua/example2
package example

import (
	example2 "github.com/radishChao/go2lua/example/go/example2"
)

type (
	// ExportToLua3
	// @lua
	// get a value
	ExportToLua3 struct {
		S1  []bool                   `lua:"get set" `
		S2  []byte                   `lua:"get set" `
		S3  []string                 `lua:"get set" `
		S4  []uint8                  `lua:"get set" `
		S5  []int8                   `lua:"get set" `
		S6  []uint16                 `lua:"get set" `
		S7  []int16                  `lua:"get set" `
		S8  []uint32                 `lua:"get set" `
		S9  []int32                  `lua:"get set" `
		S10 []uint                   `lua:"get set" `
		S11 []int                    `lua:"get set" `
		S12 []uint64                 `lua:"get set" `
		S13 []int64                  `lua:"get set" `
		S14 []float32                `lua:"get set" `
		S15 []float64                `lua:"get set" `
		S17 []*MyStruct              `lua:"get set"`
		S18 []*example2.ExportToLua2 `lua:"get set"  luaType:"struct"`
		A1  int8                     `lua:"get set"`
		A2  uint8                    `lua:"get set"`
		A3  int16                    `lua:"get set"`
		A4  int16                    `lua:"get set"`
		A5  int                      `lua:"get set"`
		A6  uint                     `lua:"get set"`
		A7  int32                    `lua:"get set"`
		A8  int32                    `lua:"get set"`
		A9  int64                    `lua:"get set"`
		A10 int64                    `lua:"get set"`
		A11 float32                  `lua:"get set"`
		A12 float64                  `lua:"get set"`
		A13 byte                     `lua:"get set"`
		A14 MyBool                   `lua:"get set"`
		A15 map[string]*MyStruct     `lua:"get set"`
		A16 map[string]int           `lua:"get set"`
		A17 *example2.ExportToLua2   `lua:"get set" luaType:"struct"`
		A18 example2.MyInt           `lua:"get set " luaType:"int32"`
		A19 MyString                 `lua:"get set " json:"name"`
		A20 *MyStruct                `lua:"get set"`
		A21 MyInt                    `lua:"get set "`
	}
	MyBool interface {
	}
	MyStrings []example2.MyInt
	MyInt     uint64
	MyString  string
	// MyStruct
	// @lua
	MyStruct struct {
		Float float32 `lua:"set get"`
	}
)

const (
	Export1 = "export1"
	Export2 = 1
)

var (
	Var1 example2.MyInt
	Var2 int
)
