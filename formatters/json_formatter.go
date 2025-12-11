package formatters

import (
	"encoding/json"

	"github.com/Rokli/LogAnalyzer-CLI/types"
)

func ToJSON(logs []types.LogEntry) ([]byte, error) {
	return json.MarshalIndent(logs, "", "  ")
}
