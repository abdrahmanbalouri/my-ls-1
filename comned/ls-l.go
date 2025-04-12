package comned

import (
	"fmt"
	"os"
	"time"
)

func ListFilesBydone(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			fmt.Println("Error getting file info:", err)
			continue
		}

		fmt.Printf("%s %10d %s %s\n",
			info.Mode(),
			info.Size(),
			info.ModTime().Format(time.RFC822),
			info.Name())
	}
}
