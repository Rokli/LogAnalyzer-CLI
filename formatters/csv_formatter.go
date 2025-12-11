package formatters

import (
	"github.com/Rokli/LogAnalyzer-CLI/types"
)

func ToCSV(analyzeFile []types.LogEntry) ([][]string, error) {
	var rows [][]string
	for _, log := range analyzeFile {
		rows = append(rows, []string{log.Timestamp, log.Level, log.Message})
	}
	return rows, nil
}
