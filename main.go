package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type manageCli struct {
	mLogs *managerLogs
}

func newManageCli() manageCli {
	var manage manageCli
	manage.mLogs = newManagerLogs()
	return manage
}

func (m manageCli) help() string {
	return "Эта утилита может работать с логами и парсить их в разные форматы данных(JSON/CSV)"
}

func (m manageCli) readFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		m.mLogs.arrayLogs = append(m.mLogs.arrayLogs, parseLine(string(scanner.Text())))
		count++
	}

	fmt.Println("Файл прочитан")
	fmt.Println("Всего строк:", count)
	fmt.Println()
}

func (m manageCli) printStatFile() {
	m.mLogs.printStatCount()
}

func (m manageCli) printFilterByLevel(command string) {
	fmt.Println("Найдено:")
	m.mLogs.printFilterByLevel(strings.Split(command, "=")[1])
}

func (m manageCli) printSubStr(command string) {
	fmt.Println("Найдено:")
	m.mLogs.printSubStr(strings.Split(command, "=")[1])
}

func (m manageCli) printOutput(command string) {
	if strings.Contains(command, "json") {
		m.mLogs.printJson()
	} else if strings.Contains(command, "csv") {
		m.mLogs.printCsv()
	}
	fmt.Println("Файл создан")
}

func main() {
	manage := newManageCli()
	command := os.Args[1:]

	if command[0] == "--help" {
		fmt.Println(manage.help())
	}

	if command[0] == "--file-read" {
		manage.readFile(command[1])
		if len(command) > 2 {
			if command[2] == "-stats" {
				manage.printStatFile()
			}

			if strings.Contains(command[2], "-level") {
				manage.printFilterByLevel(command[2])
			}

			if strings.Contains(command[2], "-search") {
				manage.printSubStr(command[2])
			}

			if strings.Contains(command[2], "-output") {
				manage.printOutput(command[2])
			}
		}

	}
}
