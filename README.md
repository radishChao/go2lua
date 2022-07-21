# go2lua
# go代码注入到lua的代码生成帮助工具
编译项目后 执行命令即可


命令如下                                           
-i, --dir string          需要注入lua的go文件目录\
-e, --export string       导出的注入lua的go文件目录      
-g, --goPackage string    需要注入lua的go文件包名        
-h, --help                help for this command          
-l, --luaPackage string   导出的注入lua的go文件包名[可选]

目前只支持文件中 type xxx struct{}的导出 其他类型的全局变量的导出还未开发

# 注释说明

### @lua-package
包注释放在文件最开头的位置\
@lua-package 引入的外部包名 外部包生成的lua注入代码包名 外部包名全路径\

示例 @lua-package example2 example2Tolua github.com/radishChao/go2lua/example/tolua/example2

### @lua 
标记要导出的成员变量 在成员变量处注释

### 字段标签 `lua:"get set"  luaType:"struct"`
标记该字段需要导出 
get,set 导出该字段的get、set函数  
luaType 当字段是在外部包定义 需要声明类型 支持 struct和基础类型

命令用法示例:
go2lua -i github.com/radishChao/go2lua/example/go  -e github.com/radishChao/go2lua/example/tolua -g github.com/radishChao/go2lua/example/go

[示例代码目录](example)

lua第三方包
github.com/yuin/gopher-lua


