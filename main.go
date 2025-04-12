package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"comned/comned"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		comn, _ := reader.ReadString('\n')
		comn = strings.TrimSpace(comn)
		if comn == "" {
			continue
		}

		args := strings.Fields(comn)
		if args[0] != "ls" {
			fmt.Println("command not exist")
			continue
		}

		path := "."
		if len(args) > 1 && !strings.HasPrefix(args[1], "-") {
			path = args[1]
			args = args[1:]
		}

		flags := ""
		if len(args) > 1 {
			flags = strings.Join(args[1:], "")
		}
		if len(args) == 1 {
			comned.ListFiles(path, false)
		}

		for _, i := range flags {
			if i != '-' {
				switch {
				case i == 'l':
					comned.ListFilesBydone(path)
				case i == 'R':
					comned.ListFilesRecursive(path)
				case i == 'a':
					comned.ListFiles(path, true)
				case i == 'r':
					comned.ListFilesReverse(path)
				case i == 't':
					comned.ListFilesByTime(path)
				default:
					fmt.Println("command not exist")
				}
			}
		}

	}
}
