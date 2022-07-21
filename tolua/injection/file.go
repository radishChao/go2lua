package injection

import (
	"fmt"
	"github.com/radishChao/go2lua/env"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type (
	packageImport struct {
		packageName string
		short       string
	}
	luaFile struct {
		// 生成的go文件名
		targetGoFileName string
		// 生成go文件包名
		targetPackage string
		// 注入到lua的table名
		luaPackage string
		// go源文件包名
		sourcePackage string
		// 导出目录
		exportDir string
		// go源文件路径
		goSourceFilePath string
		// 需要导入的包
		imports []packageImport
		// 需要注册的结构体
		registerStruct []*luaStruct
		// 需要导出的函数
		funcs []*luaFunc
		// 生成的代码
		coder *codeWriter
		// 注释
		comment *luaComment
		// 需要导入的外部包
		packages map[string]*luaFile
	}

	goCaller interface {
		getName() string
	}
)

const (
	gopherLuaShort   = "lua"
	toLuaHelperShort = "luaHelper"
	luaStateShort    = "L"
)

var (
	imports = []packageImport{
		{
			packageName: "\"github.com/yuin/gopher-lua\"",
			short:       gopherLuaShort,
		},
		{
			packageName: "\"github.com/radishChao/go2lua/tolua/helper\"",
			short:       toLuaHelperShort,
		},
	}
)

func newLuaFile() *luaFile {
	return &luaFile{
		coder:    newCodeWriter(),
		packages: map[string]*luaFile{},
	}
}

func (lf *luaFile) addRegisterStruct(ls *luaStruct) {
	if ls == nil {
		return
	}
	lf.registerStruct = append(lf.registerStruct, ls)
}

// 解析type中定义的struct数据
func (lf *luaFile) parseStructFromType(t *ast.TypeSpec) {
	comment := parseLuaComment(t.Doc)
	if !comment.isExport {
		fmt.Printf("结构体:%s未解析到@lua注释,忽略导出\n", t.Name.Name)
		return
	}
	ls := newLuaStruct(comment, lf)
	ls.parseStructFromType(t)
	lf.addRegisterStruct(ls)
}

// 解析文件的type定义
func (lf *luaFile) parseAstDecl(decls []ast.Decl) {
	for i := 0; i < len(decls); i++ {
		decl := decls[i]
		switch decl.(type) {
		case *ast.GenDecl:
			genDecl := decl.(*ast.GenDecl)
			for j := 0; j < len(genDecl.Specs); j++ {
				spec := genDecl.Specs[j]
				switch spec.(type) {
				case *ast.TypeSpec:
					t := spec.(*ast.TypeSpec)
					switch t.Type.(type) {
					case *ast.StructType:
						lf.parseStructFromType(t)
					default:
						//fmt.Println("暂不支持", t.Name, t)
					}
				case *ast.ImportSpec:
					it := spec.(*ast.ImportSpec)
					lf.imports = append(lf.imports, packageImport{
						packageName: it.Path.Value,
						short:       it.Name.Name,
					})
					lf.packages[it.Name.Name] = &luaFile{
						sourcePackage: it.Name.Name,
						imports: []packageImport{{
							packageName: it.Path.Value,
							short:       it.Name.Name,
						}},
					}

				}
			}

		}
	}
}

// 生成注释
func (lf *luaFile) genComment(comment ...*ast.CommentGroup) {
	lf.coder.writeString("// Package ")
	lf.coder.writeString(lf.targetPackage)
	lf.coder.writeString(" gen register toLua code with tag ")
	lf.coder.writeString(toLuaComment)
	lf.coder.writeString(" from file ")
	lf.coder.writeString(lf.goSourceFilePath)
	for i := 0; i < len(lf.comment.comment); i++ {
		lf.coder.writeLineEnd()
		lf.coder.writeLine(lf.comment.comment[i])
	}
	lf.coder.writeLineEnd()
}

// 生成package
func (lf *luaFile) genPackage() {
	lf.coder.writeString("package ")
	lf.coder.writeLine(lf.targetPackage)
}

// 生成import
func (lf *luaFile) genImport() {
	lf.coder.writeLine("import (")
	var importLine packageImport
	for i := 0; i < len(lf.imports); i++ {
		importLine = lf.imports[i]
		short := importLine.short
		if short != "" {
			short += " "
		}
		lf.coder.writeString(short)
		lf.coder.writeString(importLine.packageName)
		lf.coder.writeLineEnd()
	}

	for _, t := range lf.comment.packages {
		if t.short == lf.targetPackage {
			continue
		}
		short := t.short
		if short != "" {
			short += " "
		}
		lf.coder.writeString(short)
		lf.coder.writeString(t.packageName)
		lf.coder.writeLineEnd()
	}
	lf.coder.writeLine(")")
}

//生成var
func (lf *luaFile) genVar() {
	lf.coder.writeLine("var (")
	for i := 0; i < len(lf.registerStruct); i++ {
		lf.registerStruct[i].genVar(lf.coder)
	}

	lf.coder.writeLine(")")
}

// 生成文件
func (lf *luaFile) genFile() error {
	if lf.coder.len() == 0 {
		return nil
	}
	fileName := lf.targetGoFileName
	ext := filepath.Ext(fileName)
	fileName = fileName[0:strings.Index(fileName, ext)]
	exportFilePath := lf.exportDir + string(filepath.Separator) + fileName + "_toLua" + ext
	err := os.MkdirAll(lf.exportDir, os.ModePerm)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(exportFilePath, []byte(lf.coder.code()), 0666)
	if err != nil {
		return err
	}
	if env.CheckEnvGo() == nil {
		cmd := exec.Command("gofmt", "-l", "-w", exportFilePath)
		if err = cmd.Start(); err != nil { // 运行命令
			env.GetError().Printf("format file %s error：%v\n", exportFilePath, err)
		}
	}
	return nil
}

func (lf *luaFile) parseFile(filePath, exportDir, goModProject, targetPackage string, ignorePackage bool) (*ast.File, error) {
	fSet := token.NewFileSet()
	f, err := parser.ParseFile(fSet, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	lf.exportDir = filepath.Clean(exportDir)
	_, lf.targetGoFileName = filepath.Split(filePath)
	lf.goSourceFilePath = filePath
	lf.sourcePackage = f.Name.Name
	if !ignorePackage {
		lf.luaPackage = lf.sourcePackage
	}
	lf.imports = append(lf.imports, imports...)
	if filepath.Dir(filePath) != exportDir {
		// 导出的目录和.go源文件不在一个目录下
		// 需要导入.go源文件所在的包
		lf.imports = append(lf.imports, packageImport{
			packageName: fmt.Sprintf("\"%s\"", goModProject),
			short:       "",
		})
		if targetPackage == "" {
			// 注入lua的go文件包名为空使用导出目录的目录名
			targetPackage = filepath.Base(exportDir)
		}
		lf.targetPackage = targetPackage
	} else {
		// 在同一个目录下 使用go源文件的包名
		lf.targetPackage = lf.sourcePackage
	}
	return f, nil
}

// 生成注入代码
func (lf *luaFile) genRegisterCode() {
	fileName := lf.targetGoFileName[0:strings.LastIndex(lf.targetGoFileName, ".")]
	funcName := fmt.Sprintf("Register%s%sToLua", strings.ToUpper(string(fileName[0])), fileName[1:])
	lf.coder.writeString("// ")
	lf.coder.writeString(funcName)
	lf.coder.writeLine(" call this func register ")
	lf.coder.writeString("func ")
	lf.coder.writeString(funcName)
	lf.coder.writeString("(")
	lf.coder.writeString(luaStateShort)
	lf.coder.writeString(" *")
	lf.coder.writeString(gopherLuaShort)
	lf.coder.writeLine(".LState) error {")
	lf.coder.writeLine("var err error")
	for i := 0; i < len(lf.registerStruct); i++ {
		lf.registerStruct[i].genRegisterCode(lf.coder)
	}
	lf.coder.writeLine("return nil")
	lf.coder.writeLine(rbrace)
}
func (lf *luaFile) genBlockStmt() {
	for i := 0; i < len(lf.registerStruct); i++ {
		lf.registerStruct[i].genBlockStmt(lf.coder)
	}
}

// 解析文件
func (lf *luaFile) gen(filePath, exportDir, goModProject, targetPackage string, ignorePackage bool) error {
	f, err := lf.parseFile(filePath, exportDir, goModProject, targetPackage, ignorePackage)
	if err != nil {
		return err
	}
	lf.comment = parseLuaComment(f.Doc)
	lf.parseAstDecl(f.Decls)
	if len(lf.registerStruct) == 0 && len(lf.funcs) == 0 {
		return fmt.Errorf("文件:%s没有可以导出的数据", filePath)
	}
	// 生成注释
	lf.genComment(f.Doc)
	// 生成package
	lf.genPackage()
	// 生成import
	lf.genImport()
	// 生成var
	lf.genVar()
	// 生成注入代码
	lf.genRegisterCode()
	// 生成代码
	lf.genBlockStmt()
	return lf.genFile()
}
