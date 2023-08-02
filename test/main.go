package main

import (
	"fmt"
	"time"
)

var (
	readPath  = "/Users/hnijad/Desktop/lab/read-test/"
	writePath = "/Users/hnijad/Desktop/lab/write/"
)

func main() {
	startTime := time.Now()
	//readFile(readPath + "text-0.txt")
	localFileSystemReadTest()
	//localFileSystemTest()
	//generateRandomTextFiles(100000)
	//sosTest()
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v\n", duration.Seconds())
}
