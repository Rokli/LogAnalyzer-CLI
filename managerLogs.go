package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type managerLogs struct {
	arrayLogs []LogEntry
	count     map[string]int
}

func newManagerLogs() *managerLogs {
	var manager managerLogs

	manager.count = map[string]int{"INFO": 0, "ERROR": 0, "WARN": 0}

	return &manager
}

func (m managerLogs) logAnalyze() {
	for _, log := range m.arrayLogs {
		if log.Level == "INFO" {
			m.count["INFO"]++
		}

		if log.Level == "ERROR" {
			m.count["ERROR"]++
		}

		if log.Level == "WARN" {
			m.count["WARN"]++
		}
	}
}

func (m managerLogs) printStatCount() {
	m.logAnalyze()
	for key, value := range m.count {
		fmt.Println(key, ":", value)
	}
}

func (m managerLogs) filterByLevel(level string) []LogEntry {
	var levels []LogEntry

	for _, value := range m.arrayLogs {
		if value.Level == level {
			levels = append(levels, value)
		}
	}
	return levels
}

func (m managerLogs) printFilterByLevel(level string) {
	var levels []LogEntry = m.filterByLevel(level)

	for _, value := range levels {
		fmt.Println(value.Timestamp, " ", value.Level, " ", value.Message)
	}
}

func (m managerLogs) findSubStr(str string) []LogEntry {
	var findMessage []LogEntry

	for _, value := range m.arrayLogs {
		if strings.Contains(value.Message, str) {
			findMessage = append(findMessage, value)
		}
	}
	return findMessage
}

func (m managerLogs) printSubStr(str string) {
	var subStr []LogEntry = m.findSubStr(str)

	for _, value := range subStr {
		fmt.Println(value.Timestamp, " ", value.Level, " ", value.Message)
	}
}

func (m managerLogs) printJson() {
	b, err := json.MarshalIndent(m.arrayLogs, "\n", " ")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := os.Create("output.json")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	file.Write(b)
}

func (m managerLogs) printCsv() {
	var data [][]string

	for _, value := range m.arrayLogs {
		var tmp []string
		tmp = append(tmp, value.Timestamp)
		tmp = append(tmp, value.Level)
		tmp = append(tmp, value.Message)

		data = append(data, tmp)
	}

	file, err := os.Create("output.csv")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	for _, record := range data {
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	writer.Flush()
}

func (m managerLogs) printLimitStr(number int) {
	for i := 0; i < number; i++ {
		fmt.Println(
			m.arrayLogs[i].Timestamp,
			" ",
			m.arrayLogs[i].Level,
			" ",
			m.arrayLogs[i].Message,
		)
	}
}
