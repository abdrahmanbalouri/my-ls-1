package comned

import (
	"fmt"
	"os"
	"sort"
	"time"
)

type FileInfoStruct struct {
	Name    string
	ModTime time.Time
}

func ListFilesByTime(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	var fileInfos []FileInfoStruct

	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			fmt.Println("Error getting file info:", err)
			continue
		}
		fileInfos = append(fileInfos, FileInfoStruct{Name: file.Name(), ModTime: info.ModTime()})
	}

	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].ModTime.After(fileInfos[j].ModTime)
	})

	for _, file := range fileInfos {
		fmt.Println(file.ModTime.Format(time.RFC822), "-", file.Name)
	}
}
