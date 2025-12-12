package formatters

import (
	"encoding/json"

	"github.com/Rokli/LogAnalyzer-CLI/pkg/logs"
)

func ToJSON(logs []logs.LogEntry) ([]byte, error) {
	return json.MarshalIndent(logs, "", "  ")
}
