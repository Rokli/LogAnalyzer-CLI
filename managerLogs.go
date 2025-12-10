package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

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

func ToJSON(analyzeFile []LogEntry) {
	b, err := json.MarshalIndent(analyzeFile, "\n", " ")

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

func ToCSV(analyzeFile []LogEntry) {
	var data [][]string

	for _, value := range analyzeFile {
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

func GetLimitStr(analyzeFile []LogEntry, number int) []LogEntry {
	var limitAnalyzeFile []LogEntry
	for i := 0; i < number; i++ {
		limitAnalyzeFile = append(limitAnalyzeFile, analyzeFile[i])
	}
	return limitAnalyzeFile
}

func help() string {
	return "Эта утилита может работать с логами и парсить их в разные форматы данных(JSON/CSV)"
}
