package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var fp *os.File
	//fmt.Println(reflect.TypeOf(fp))

	if len(os.Args) != 2 {
		fmt.Println("The number of arguments specified is incorrect.")
		os.Exit(1)
	} else {
		folderList := os.Args[1] // Folder list(*.txt)
		var err error
		ext := filepath.Ext(folderList)
		if ext != ".txt" {
			fmt.Println("Please specify *.txt.")
			os.Exit(1)
		}
		fp, err = os.Open(folderList)
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		folderName := scanner.Text()
		if err := os.Mkdir(folderName, 0777); err != nil {
			fmt.Println("Error:")
			fmt.Println(err)
		}
	}

	fmt.Println("\nDone!")
}
