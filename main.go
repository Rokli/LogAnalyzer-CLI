package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(filename string) []LogEntry {
	var parseFile []LogEntry
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		parseFile = append(parseFile, parseLine(string(scanner.Text())))
		count++
	}

	fmt.Println("Файл прочитан")
	fmt.Println("Всего строк:", count)
	fmt.Println()

	return parseFile
}

func printOutput(output []LogEntry) {
	for _, value := range output {
		fmt.Println(
			value.Timestamp,
			" ",
			value.Level,
			" ",
			value.Message,
		)
	}
}

func main() {
	var cfg Config = parseFlags()

	err := validateConfig(cfg)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var analyzeFile []LogEntry = readFile(cfg.File)

	if cfg.Help {
		fmt.Println(help())
		return
	}

	if cfg.Limit != 0 {
		analyzeFile = GetLimitStr(analyzeFile, cfg.Limit)
	}

	if cfg.Stats {
		var output map[string]int = GetStats(analyzeFile)
		for key, value := range output {
			fmt.Println(key, ":", value)
		}
		fmt.Println()
	}

	if cfg.Level != "" {
		analyzeFile = GetFilterByLevel(analyzeFile, cfg.Level)
	}

	if cfg.Search != "" {
		analyzeFile = GetFindSubStr(analyzeFile, cfg.Search)
	}

	if cfg.Search != "" {
		if cfg.Search == "json" {
			ToJSON(analyzeFile)
		} else if cfg.Search == "csv" {
			ToCSV(analyzeFile)
		}
	}

	printOutput(analyzeFile)
}
