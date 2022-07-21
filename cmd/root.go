/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/radishChao/go2lua/env"
	"github.com/radishChao/go2lua/tolua"
	"github.com/radishChao/go2lua/util"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

type (
	fileGroup struct {
		files       []string
		goPath      string
		packageName string
		exportDir   string
	}
)

// rootCmd represents the base command when called without any subcommands
var (
	importDir    string
	exportDir    string
	rootGoPath   string
	targetGoPath string

	rootCmd = &cobra.Command{
		Long: `go2lua
lua register tool
https://github.com/radishChao/go2lua`,
		Run: func(cmd *cobra.Command, args []string) {
			if importDir == "" {
				env.GetError().Println("must be set arg dir[-d]")
				return
			}
			importDir = filepath.Clean(importDir)
			if exportDir == "" {
				env.GetError().Println("must be set arg export[-e]")
				return
			}
			exportDir = filepath.Clean(exportDir)
			if rootGoPath == "" {
				env.GetError().Println("must be set arg package[-p]")
				return
			}
			var groups []*fileGroup
			rootGoPath = filepath.Clean(rootGoPath)
			if importDir != "" {
				var err error
				groups, err = parseDir(importDir)
				if err != nil {
					env.GetError().Printf("parse importDir error :%v\n", err)
					return
				}
			}

			gLen := len(groups)
			if gLen == 0 {
				env.GetError().Println("nothing to export register files")
				return
			}
			for i := 0; i < gLen; i++ {
				group := groups[i]
				tolua.Parse(group.exportDir, group.goPath, targetGoPath, group.files...)
			}
		},
	}
)

func parseDir(importDir string) ([]*fileGroup, error) {
	files, err := util.LoopFiles(importDir)
	if err != nil {
		return nil, err
	}
	groupMap := map[string]*fileGroup{}

	for i := 0; i < len(files); i++ {
		f := files[i]
		groupId := filepath.Dir(f)
		group, ok := groupMap[groupId]
		if !ok {
			group = &fileGroup{
				files:     nil,
				goPath:    filepath.ToSlash(rootGoPath),
				exportDir: exportDir,
			}
			if groupId != "." {
				group.goPath = filepath.ToSlash(rootGoPath + string(filepath.Separator) + groupId)
				group.exportDir = exportDir + string(filepath.Separator) + groupId
			}
			groupMap[groupId] = group
		}
		group.files = append(group.files, filepath.ToSlash(importDir+string(filepath.Separator)+f))
	}

	var groups []*fileGroup
	for _, group := range groupMap {
		groups = append(groups, group)
	}
	return groups, nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go2lua.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&importDir, "dir", "i", "", "需要注入lua的go文件目录")
	rootCmd.Flags().StringVarP(&exportDir, "export", "e", "", "导出的注入lua的go文件目录")
	rootCmd.Flags().StringVarP(&rootGoPath, "goPackage", "g", "", "需要注入lua的go文件包名")
	rootCmd.Flags().StringVarP(&targetGoPath, "luaPackage", "l", "", "导出的注入lua的go文件包名[可选]")

}
