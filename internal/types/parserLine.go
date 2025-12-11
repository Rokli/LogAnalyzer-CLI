package types

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type LogEntry struct {
	Timestamp string
	Level     string
	Message   string
}

func ParseLine(line string) LogEntry {
	array := strings.Split(line, " ")
	parse := LogEntry{array[0] + " " + array[1], array[2], strings.Join(array[3:], " ")}
	return parse
}

func GetStats(analyzeFile []LogEntry) map[string]int {
	var count map[string]int = map[string]int{"INFO": 0, "ERROR": 0, "WARN": 0}
	for _, log := range analyzeFile {
		if log.Level == "INFO" {
			count["INFO"]++
		}

		if log.Level == "ERROR" {
			count["ERROR"]++
		}

		if log.Level == "WARN" {
			count["WARN"]++
		}
	}

	return count
}

func GetFilterByLevel(analyzeFile []LogEntry, level string) []LogEntry {
	var levels []LogEntry

	for _, value := range analyzeFile {
		if value.Level == level {
			levels = append(levels, value)
		}
	}
	return levels
}

func GetFindSubStr(analyzeFile []LogEntry, subStr string) []LogEntry {
	var findMessage []LogEntry

	for _, value := range analyzeFile {
		if strings.Contains(value.Message, subStr) {
			findMessage = append(findMessage, value)
		}
	}
	return findMessage
}

func CreateFileJSON(data []byte) {
	file, err := os.Create("output.json")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	file.Write(data)
}

func CreateFileCSV(data [][]string) {

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

func GetLimitStr(analyzeFile []LogEntry, number int) []LogEntry {
	var limitAnalyzeFile []LogEntry
	for i := 0; i < number; i++ {
		limitAnalyzeFile = append(limitAnalyzeFile, analyzeFile[i])
	}
	return limitAnalyzeFile
}

func Help() string {
	return "Эта утилита может работать с логами и парсить их в разные форматы данных(JSON/CSV)"
}
