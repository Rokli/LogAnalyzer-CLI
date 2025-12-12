package types

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/Rokli/LogAnalyzer-CLI/pkg/logs"
)

func ParseLine(line string) (logs.LogEntry, error) {
	array := strings.Split(line, " ")
	if len(array) < 3 {
		return logs.LogEntry{}, fmt.Errorf("invalid log format: %s", line)
	}

	timestamp := array[0] + " " + array[1]
	level := array[2]
	message := strings.Join(array[3:], " ")

	return logs.LogEntry{Timestamp: timestamp, Level: level, Message: message}, nil
}

func GetStats(analyzeFile []logs.LogEntry) map[string]int {
	count := map[string]int{
		logs.LevelInfo:  0,
		logs.LevelError: 0,
		logs.LevelWarn:  0,
	}
	for _, log := range analyzeFile {
		if val, exists := count[log.Level]; exists {
			count[log.Level] = val + 1
		}
	}
	return count
}

func GetFilterByLevel(analyzeFile []logs.LogEntry, level string) []logs.LogEntry {
	var levels []logs.LogEntry

	for _, value := range analyzeFile {
		if value.Level == level {
			levels = append(levels, value)
		}
	}
	return levels
}

func GetFindSubStr(analyzeFile []logs.LogEntry, subStr string) []logs.LogEntry {
	var findMessage []logs.LogEntry

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

func GetLimitStr(analyzeFile []logs.LogEntry, number int) []logs.LogEntry {
	var limitAnalyzeFile []logs.LogEntry
	for i := 0; i < number; i++ {
		limitAnalyzeFile = append(limitAnalyzeFile, analyzeFile[i])
	}
	return limitAnalyzeFile
}

func Help() string {
	return "Эта утилита может работать с логами и парсить их в разные форматы данных(JSON/CSV)"
}
