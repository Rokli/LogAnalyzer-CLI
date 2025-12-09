package main

import "fmt"

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
	for key, value := range m.count {
		fmt.Println(key, ":", value)
	}
}
