package util

import (
	"io/ioutil"
	"path/filepath"
)

func LoopFiles(directory string) ([]string, error) {
	directory = filepath.ToSlash(directory + string(filepath.Separator))
	//获取文件或目录相关信息
	fileInfoList, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}
	var files []string
	for i := range fileInfoList {
		if fileInfoList[i].IsDir() {
			fs, err := LoopFiles(directory + fileInfoList[i].Name())
			if err != nil {
				return nil, err
			}
			for j := 0; j < len(fs); j++ {
				files = append(files, filepath.ToSlash(fileInfoList[i].Name()+string(filepath.Separator)+fs[j]))
			}

		} else {
			files = append(files, fileInfoList[i].Name())
		}
	}
	return files, nil
}
