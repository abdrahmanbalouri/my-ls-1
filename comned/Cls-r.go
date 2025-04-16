package comned

import (
	"fmt"
	"os"
)

func ListFilesRecursive(root string) {
	entries, err := os.ReadDir(root)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, entry := range entries {
		path := root + "/" + entry.Name()
		fmt.Println(path)

		if entry.IsDir() {
			ListFilesRecursive(path) // Call recursively for subdirectories
		}
	}
}
