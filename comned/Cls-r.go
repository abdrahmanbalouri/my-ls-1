package comned

import (
	"fmt"
	"os"
	"path/filepath"
)

func ListFilesRecursive(root string) {
	filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		fmt.Println(path)
		return nil
	})
}
