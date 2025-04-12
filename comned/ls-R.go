package comned

import (
	"fmt"
	"os"
	"sort"
)

func ListFilesReverse(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	sort.Sort(sort.Reverse(sort.StringSlice(fileNames)))

	for _, name := range fileNames {
		fmt.Println(name)
	}
}
