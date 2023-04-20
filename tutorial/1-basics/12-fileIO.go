package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func writeToFile(fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	iPrimeArr := []int{2, 3, 5, 7, 11}
	var sPrimeArr []string
	for _, num := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(num))
	}
	for _, sNum := range sPrimeArr {
		_, err := f.WriteString(sNum + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	f, err = os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		fmt.Println("Prime 1:", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}
}

func appendToFile(fileName string, info string) {
	// Append to file
	/*
		Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified
		O_RDONLY : open the file read-only
		O_WRONLY : open the file write-only
		O_RDWR   : open the file read-write
		These can be or'ed
		O_APPEND : append data to the file when writing
		O_CREATE : create a new file if none exists
		O_EXCL   : used with O_CREATE, file must not exist
		O_SYNC   : open for synchronous I/O
		O_TRUNC  : truncate regular writable file when opened
	*/

	_, err := os.Stat(fileName)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File doesn't exist")
		return
	}

	// filename, Openfile mode, permissions of new file
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	if _, err := f.WriteString(info); err != nil {
		log.Fatal(err)
	}
}

func main() {
	writeToFile("data.txt")
	appendToFile("data.txt", "13")
}
