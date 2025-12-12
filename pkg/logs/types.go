package logs

type LogEntry struct {
	Timestamp string `json:"timestamp" csv:"timestamp"`
	Level     string `json:"level" csv:"level"`
	Message   string `json:"message" csv:"message"`
}
