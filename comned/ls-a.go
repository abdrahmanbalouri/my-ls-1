package comned

import (
	"fmt"
	"os"
)

func ListFiles(path string, showHidden bool) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if !showHidden && file.Name()[0] == '.' {
			continue
		}
		fmt.Println(file.Name())
	}
}
