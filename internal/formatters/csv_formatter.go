package formatters

import (
	"github.com/Rokli/LogAnalyzer-CLI/pkg/logs"
)

func ToCSV(analyzeFile []logs.LogEntry) ([][]string, error) {
	var rows [][]string
	for _, log := range analyzeFile {
		rows = append(rows, []string{log.Timestamp, log.Level, log.Message})
	}
	return rows, nil
}
