package injection

import (
	"go/ast"
	"strings"
)

type (
	luaComment struct {
		get      bool
		set      bool
		isExport bool
		// 需要导入的外部包
		packages map[string]*packageImport
		// 外部定义的字段
		fields map[string]*luaStructOutField
		// 注释
		comment []string
	}
)

const (
	toLuaComment        = "@lua"
	toLuaCommentPackage = "@lua-package"
	toLuaCommentType    = "@lua-type"
)

func parseLuaComment(commentGroup *ast.CommentGroup) *luaComment {
	c := &luaComment{
		get:      false,
		set:      false,
		isExport: false,
		packages: map[string]*packageImport{},
		fields:   map[string]*luaStructOutField{},
		comment:  nil,
	}
	c.parse(commentGroup)
	return c
}

func (ltc *luaComment) parseTolua(comment string) {
	ltc.isExport = true
	comments := strings.Split(comment, string(toluaTagSplit))
	for i := 0; i < len(comments); i++ {
		if comments[i] == toluaTagGet {
			ltc.get = true
		} else if comments[i] == toluaTag {
			ltc.set = true
		}
	}
}

// 解析@lua-package
func (ltc *luaComment) parseToluaPackage(comment string) {
	comments := strings.Split(comment, string(toluaTagSplit))
	if len(comments) >= 3 {
		short := comments[0]
		tolua := comments[1]
		pFull := comments[2]
		ltc.packages[short] = &packageImport{
			packageName: "\"" + pFull + "\"",
			short:       tolua,
		}
	}
}

// 解析@lua-type
func (ltc *luaComment) parseToluaType(comment string) {
	comments := strings.Split(comment, string(toluaTagSplit))
	commentsSize := len(comments)
	if commentsSize >= 2 {
		ltc.fields[comments[0]] = &luaStructOutField{
			typeName: comments[1],
		}
	}
}

// 解析注释
func (ltc *luaComment) parse(commentGroup *ast.CommentGroup) {
	if commentGroup == nil {
		return
	}
	for i := 0; i < len(commentGroup.List); i++ {
		comment := commentGroup.List[i].Text
		toluaPackageIndex := strings.Index(comment, toLuaCommentPackage)
		toluaIndex := strings.Index(comment, toLuaComment)
		toluaFieldIndex := strings.Index(comment, toLuaCommentType)
		if toluaIndex != -1 && toluaFieldIndex == -1 && toluaPackageIndex == -1 {
			// 可导出
			ltc.parseTolua(strings.TrimSpace(comment[toluaIndex+len(toLuaComment):]))
		} else if toluaPackageIndex != -1 {
			ltc.parseToluaPackage(strings.TrimSpace(comment[toluaPackageIndex+len(toLuaCommentPackage):]))
		} else if toluaFieldIndex != -1 {
			ltc.parseToluaType(strings.TrimSpace(comment[toluaFieldIndex+len(toLuaCommentType):]))
		} else {
			ltc.comment = append(ltc.comment, comment)
		}
	}

	//if toluaPackageIndex == -1 {
	//	return
	//}
	//comment = strings.TrimSpace(comment[toluaPackageIndex+len(toLuaCommentPackage):])
	//comment = strings.TrimSpace(comment)
	//comments := strings.Split(comment, string(toluaTagSplit))
	//if len(comments) >= 3 {
	//	short := comments[0]
	//	pFull := comments[1]
	//}
	//var tag []uint8
	//for i := 0; i < len(comment); i++ {
	//	c := comment[i]
	//	if c == toLuaCommentSplit {
	//		tag = tag[len(tag):]
	//		continue
	//	}
	//	tag = append(tag, c)
	//	if len(tag) == 3 {
	//		s := string(tag)
	//		if s == toLuaCommentGet {
	//			ltc.get = true
	//		} else if s == toLuaCommentSet {
	//			ltc.set = true
	//		}
	//	}
	//}
}
