package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func localFileSystemTest() {
	dirEntries, err := os.ReadDir(readPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		os.Exit(1)
	}
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			continue
		}

		filePath := filepath.Join(readPath, dirEntry.Name())

		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", filePath, err)
			continue
		}

		_ = os.WriteFile(writePath+dirEntry.Name(), data, 0644)
	}
}

func readFile(path string) {
	cont, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	_ = len(cont) / 10
}

func localFileSystemReadTest() {
	limit := 100000
	for i := 0; i < limit; i++ {
		name := fmt.Sprintf("text-%v.txt", i)
		filePath := filepath.Join(readPath, name)
		readFile(filePath)
		//fmt.Println("reading", i)
	}
}
