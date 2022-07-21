package tolua

import (
	"fmt"
	"github.com/radishChao/go2lua/tolua/injection"
)

type (
	parseResult struct {
		filePath string
		err      error
	}
)

func parse(filePath, exportDir, goModProject, targetPackage string, ch chan parseResult) {
	go func() {
		err := injection.Parse(filePath, exportDir, goModProject, targetPackage, false)
		ch <- parseResult{
			filePath: filePath,
			err:      err,
		}
	}()
}

func Parse(exportDir, goModProject, targetPackage string, filePath ...string) {
	size := len(filePath)
	if size == 0 {
		return
	}
	ch := make(chan parseResult)
	for i := 0; i < size; i++ {
		parse(filePath[i], exportDir, goModProject, targetPackage, ch)
	}
	i := 0
	for i < size {
		result := <-ch
		i++
		if result.err != nil {
			fmt.Printf("gen:%s,fail:%s\n", result.filePath, result.err.Error())
			continue
		}
		fmt.Printf("gen:%s success\n", result.filePath)

	}

}
