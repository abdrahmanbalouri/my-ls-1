package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"comned/comned"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {

		var flagg string
		fmt.Print(">--- ")
		comn, _ := reader.ReadString('\n')
		comn = strings.TrimSpace(comn)
		if comn == "" {
			continue
		}

		args := strings.Fields(comn)
		if len(args) > 3 {
			fmt.Println("iwant just 3")
			continue
		}
		if args[0] != "ls" {
			fmt.Println("command not exist")
			continue
		}

		path := "."
		if len(args) > 1 && !strings.HasPrefix(args[1], "-") {
			path = args[1]
			args = args[1:]

			b, err := os.Stat(path)
			if err != nil {
				fmt.Println("makaynch", b)
				continue
			} else if b.IsDir() {
				fmt.Println()
			} else {
				if len(args) == 1 {
					//kk:=trim(path)
					fmt.Println(b.Name())
					continue
				} else if len(args) > 1 {
					flag := args[1]
					if len(flag) == 1 {
						fmt.Println("not a flag  plz  change ")
						continue
					}
					if flag[0] != '-' {
						fmt.Println("not flag")
						continue
					}
					newflag := flag[1:]
					if !bolle(newflag) {

						fmt.Println("machi flag")
						continue
					}
					flagg := sort(newflag)

					for i := 0; i < len(flagg); i++ {
						switch {
						case flagg[i] == 'l':
							fmt.Printf("%s %10d %s %s\n",
								b.Mode(),
								b.Size(),
								b.ModTime().Format(time.RFC822),
								b.Name())
						case flagg[i] == 't':
							fmt.Println(b.Name())
						case flagg[i] == 'R':
							fmt.Println(b.Name())
						case flagg[i] == 'r':
							fmt.Println(b.Name())
						case flagg[i] == 'a':
							fmt.Println(b.Name())




						}

					}
					continue

				}

			}
		}

		flags := ""
		if len(args) > 1 {
			flags = strings.Join(args[1:], "")
		}
		if flags != "" {
			flagg = sort(flags[1:])
		}

		if !bolle(flags) {
			fmt.Println("machi flag")
			continue
		}

		fmt.Println(flagg)
		if len(args) == 1 {
			comned.ListFiles(path, false)
		}

		for _, i := range flagg {

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
