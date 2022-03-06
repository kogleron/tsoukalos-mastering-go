package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logFilePath := getLogFilePath()
	fmt.Println("Log file: " + logFilePath)

	file := openFile(logFilePath)
	defer file.Close()

	logger := log.New(file, "customLogger ", log.LstdFlags|log.Lshortfile)
	logger.Println("Hello there")
	logger.Println("Another entry")
}

func openFile(logFilePath string) *os.File {
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panic(err)
	}
	return file
}

func getLogFilePath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	return dir + "/log.txt"
}
