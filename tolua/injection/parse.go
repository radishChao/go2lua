package injection

// Parse 解析go文件
func Parse(filePath, exportDir, goModProject, targetPackage string, ignorePackage bool) error {
	lf := newLuaFile()
	err := lf.gen(filePath, exportDir, goModProject, targetPackage, ignorePackage)
	if err != nil {
		return err
	}
	return nil
}

func newLuaFieldByTyName(name string) luaField {
	var t luaField
	switch name {
	case typeString:
		t = newLuaFieldString()
	case typeInt8:
		t = newLuaFieldInt8()
	case typeUInt8:
		t = newLuaFieldUint8()
	case typeBool:
		t = newLuaFieldBool()
	case typeByte:
		t = newLuaFieldUint8()
	case typeInt16:
		t = newLuaFieldInt16()
	case typeUInt16:
		t = newLuaFieldUint16()
	case typeInt:
		t = newLuaFieldInt()
	case typeUInt:
		t = newLuaFieldUint()
	case typeInt32:
		t = newLuaFieldInt32()
	case typeUInt32:
		t = newLuaFieldUint32()
	case typeInt64:
		t = newLuaFieldInt64()
	case typeUInt64:
		t = newLuaFieldUint64()
	case typeFloat32:
		t = newLuaFieldFloat32()
	case typeFloat64:
		t = newLuaFieldFloat64()
	}
	return t
}
