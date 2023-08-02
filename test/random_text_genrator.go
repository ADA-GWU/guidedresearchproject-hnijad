package main

import (
	"fmt"
	"math/rand"
	"os"
)

func getRandomString() string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	l := rand.Int()%1000 + 1
	l = 1000
	b := make([]byte, l)
	for i := 0; i < l; i++ {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func generateRandomTextFiles(cnt int) {
	for i := 0; i < cnt; i++ {
		name := fmt.Sprintf("text-%d.txt", i)
		file, err := os.Create(readPath + name)
		if err != nil {
			fmt.Println("Could not create text file")
			return
		}
		_, err = file.Write([]byte(getRandomString()))
		if err != nil {
			fmt.Println("Could not write to the file")
			return
		}
		_ = file.Close()
	}
}
