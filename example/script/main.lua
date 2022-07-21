local v1 = example1.ExportToLua1.new()


local myS = example1.MyStruct.new()
myS:Float(123)
v1:A20(myS)


local v2 = example2.ExportToLua2.new()
v2:I(3)
v1:A17(v2)


print(v1:A17():I(),v1:A20():Float())