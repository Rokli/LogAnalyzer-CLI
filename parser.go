package main

import (
	"strings"
)

type LogEntry struct {
	Timestamp string
	Level     string
	Message   string
}

func parseLine(line string) LogEntry {
	array := strings.Split(line, " ")
	parse := LogEntry{array[0] + " " + array[1], array[2], strings.Join(array[3:], " ")}
	return parse
}
