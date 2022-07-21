package example2

type (
	// ExportToLua2
	// @lua
	// get a value
	ExportToLua2 struct {
		// 名字
		I MyInt `lua:"set get"`
	}
	MyInt int32
)
