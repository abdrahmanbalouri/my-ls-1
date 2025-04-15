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
         path:="."
		var flage []string
		fmt.Print(">--- ")
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
		if len(args)==1&& args[0]=="ls"{
			comned.ListFiles(path, false)
			continue

		}

		
		if len(args) > 1 && !strings.HasPrefix(args[1], "-") {
			path = args[1]
			args = args[1:]

			b, err := os.Stat(path)
			if err != nil {
				fmt.Println("makaynch", b)
				continue
			} else if b.IsDir() {
				if len(args) == 1 {
					// kk:=trim(path)
					comned.ListFiles(path, false)
					continue
				} else if len(args) > 1 {
					flag := args[1:]
					kk := strings.Join(flag, " ")
					flage = strings.Fields(kk)
				}
				l:=true
				for i := 0; i < len(flage); i++ {
					if flage[i][0] != '-'|| len(flage[i])==1 && flage[i]=="-" {
						fmt.Println("not a flg")
						l=false
						break
					}
				}
				if !l{
					continue
				}
				if !bolle(flage) {
					fmt.Println("not  a flage")
					continue

				}
				mratab := sort(flage)
				for _, i := range mratab {
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

					}
				}
				continue
			} else {
				if len(args) == 1 {
					// kk:=trim(path)
					fmt.Println(b.Name())
					continue
				} else if len(args) > 1 {
					flag := args[1:]
					kk := strings.Join(flag, " ")
					flage = strings.Fields(kk)
					fmt.Println(flage)
				}
				l:=true
				for i := 0; i < len(flage); i++ {
					if flage[i][0] != '-'|| len(flage[i])==1 && flage[i]=="-"  {
						fmt.Println("not a flg")
						l=false
						break
						
					}
				}
				if !l{
					continue
				}
				if !bolle(flage) {
					fmt.Println("not  a flage")
					continue

				}
				mratab := sort(flage)
				for _, i := range mratab {
					switch {
					case i == 'l':
						fmt.Printf("%s %10d %s %s\n",
							b.Mode(),
							b.Size(),
							b.ModTime().Format(time.RFC822),
							b.Name())
					case i == 'R':
						fmt.Println(b.Name())
					case i == 'a':
						fmt.Println(b.Name())

					case i == 'r':
						fmt.Println(b.Name())

					case i == 't':
						fmt.Println(b.Name())

					}
				}
				continue

			}

		}
		flag:=args[1:]
		kk := strings.Join(flag, " ")
	    flage = strings.Fields(kk)
		fmt.Println(flage)
		l:=true
		for i := 0; i < len(flage); i++ {
			if flage[i][0] != '-'|| len(flage[i])==1 && flage[i]=="-"  {
				fmt.Println("not a flg")
				l=false
				break
			}
		}
		if !l{
			continue

		}
		if !bolle(flage) {
			fmt.Println("not  a flage")
			continue

		}
		mratab := sort(flage)
		for _, i := range mratab {
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

			}
		}

		
	}
}
