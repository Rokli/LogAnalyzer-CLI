package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type manage struct {
	filePath string
}

func (m manage) help() string {
	return "Эта утилита может работать с логами и парсить их в разные форматы данных(JSON/CSV)"
}

func (m manage) pathFile(path string) {
	m.filePath = path
	fmt.Println("added path:", path)
}

func main() {
	fmt.Println(os.Args[1:])

	manage := new(manage)

	for true {
		fmt.Print("> ")
		// fmt.Scan(&command)
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		commandArray := strings.Split(command, " ")

		// fmt.Println(commandArray)

		if commandArray[0] == "help" {
			fmt.Println(manage.help())
		}

		if commandArray[0] == "file-path" {
			manage.pathFile(commandArray[1])
		}

		if commandArray[0] == "exit" {
			break
		}
	}

}
