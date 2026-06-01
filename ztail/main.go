package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 4 || os.Args[1] != "-c" {
		return
	}
	count := 0
	for _, c := range os.Args[2] {
		count = count*10 + int(c-'0')
	}
	files := os.Args[3:]
	exitCode := 0
	for i, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			exitCode = 1
			continue
		}
		if len(files) > 1 {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", file)
		}
		if count < len(data) {
			data = data[len(data)-count:]
		}
		fmt.Print(string(data))
	}
	if exitCode != 0 {
		fmt.Fprintf(os.Stderr, "")
		os.Exit(1)
	}
}
