package main

import (
	"fmt"
	"io"
	"os"
)

type manage struct{}

func (m manage) help() string {
	return "Эта утилита может работать с логами и парсить их в разные форматы данных(JSON/CSV)"
}

func (m manage) readFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	data := make([]byte, 64)

	count := 0
	for {
		_, err := file.Read(data)
		count++
		if err == io.EOF {
			break
		}
	}
	fmt.Println("Всего строк:", count)
}

func main() {
	manage := new(manage)
	command := os.Args[1:]

	if command[0] == "--help" {
		fmt.Println(manage.help())
	}

	if command[0] == "--file-read" {
		manage.readFile(command[1])
	}
}
