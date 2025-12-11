package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Rokli/LogAnalyzer-CLI/internal/formatters"
	"github.com/Rokli/LogAnalyzer-CLI/internal/types"
)

func readFile(filename string) []types.LogEntry {
	var parseFile []types.LogEntry
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		parseFile = append(parseFile, types.ParseLine(string(scanner.Text())))
		count++
	}

	fmt.Println("Файл прочитан")
	fmt.Println("Всего строк:", count)
	fmt.Println()

	return parseFile
}

func printOutput(output []types.LogEntry) {
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
	var cfg types.Config = types.ParseFlags()

	err := types.ValidateConfig(cfg)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var analyzeFile []types.LogEntry = readFile(cfg.File)

	if cfg.Help {
		fmt.Println(types.Help())
		return
	}

	if cfg.Limit != 0 {
		analyzeFile = types.GetLimitStr(analyzeFile, cfg.Limit)
	}

	if cfg.Stats {
		var output map[string]int = types.GetStats(analyzeFile)
		for key, value := range output {
			fmt.Println(key, ":", value)
		}
		fmt.Println()
	}

	if cfg.Level != "" {
		analyzeFile = types.GetFilterByLevel(analyzeFile, cfg.Level)
	}

	if cfg.Search != "" {
		analyzeFile = types.GetFindSubStr(analyzeFile, cfg.Search)
	}

	if cfg.Search != "" {
		if cfg.Search == "json" {
			data, err := formatters.ToJSON(analyzeFile)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			types.CreateFileJSON(data)
		} else if cfg.Search == "csv" {
			data, err := formatters.ToCSV(analyzeFile)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			types.CreateFileCSV(data)
		}
	}

	printOutput(analyzeFile)
}
