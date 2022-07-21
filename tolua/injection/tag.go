package injection

import (
	"bytes"
	"go/ast"
)

type (
	luaTag struct {
		isExport bool
		get      bool
		set      bool
	}
	luaTypeTag struct {
		typ string
	}
)

const (
	toluaTag           = "lua"
	toluaTagGet        = "get"
	toluaTagSet        = "set"
	toluaTagStart      = ':'
	toluaTagSplit      = ' '
	toluaTagStartQuote = '"'

	toluaTypeTag = "luaType"
)

// 解析字段的luaType tag
func parseFieldLuaTypeTag(tag *ast.BasicLit) *luaTypeTag {
	if tag == nil {
		return nil
	}
	text := tag.Value[1 : len(tag.Value)-1]
	var bStart bool
	var value bytes.Buffer
	for i := 0; i < len(text); i++ {
		c := text[i]
		if c == toluaTagStart || c == toluaTagSplit || c == toluaTagStartQuote {
			if bStart && c == toluaTagStartQuote {
				return &luaTypeTag{
					typ: value.String(),
				}
			}
			value.Reset()
			continue
		}
		value.WriteByte(c)
		if value.String() == toluaTypeTag {
			bStart = true
			i += 2
			value.Reset()
			continue
		}

	}
	return nil
}

// 解析字段的lua tag
func parseFieldLuaTag(tag *ast.BasicLit) *luaTag {
	fTag := &luaTag{
		isExport: false,
		get:      false,
		set:      false,
	}
	if tag == nil {
		return fTag
	}
	text := tag.Value[1 : len(tag.Value)-1]
	var (
		oneTag       []uint8
		tagFlag      []uint8
		bLuaTagStart bool
		quotes       int
	)
	for i := 0; i < len(text); i++ {
		c := text[i]
		oneTag = append(oneTag, c)
		if string(oneTag) == toluaTag && !bLuaTagStart {
			fTag.isExport = true
			bLuaTagStart = true
			tagFlag = tagFlag[len(tagFlag):]
			oneTag = oneTag[len(oneTag):]
			continue
		}
		if bLuaTagStart {
			if c == toluaTagSplit || c == toluaTagStart || c == toluaTagStartQuote {
				if c == toluaTagStart {
					quotes++
				}
				if c == toluaTagSplit {
					tagFlag = tagFlag[len(tagFlag):]
				}
				continue
			}

			tagFlag = append(tagFlag, c)
			if quotes == 2 {
				break
			}

			if len(tagFlag) == 3 {
				if string(tagFlag) == toluaTagGet {
					fTag.get = true
					tagFlag = tagFlag[len(tagFlag):]
					continue
				}
				if string(tagFlag) == toluaTagSet {
					fTag.set = true
					tagFlag = tagFlag[len(tagFlag):]
					continue
				}
			}

		}
	}
	return fTag
}
