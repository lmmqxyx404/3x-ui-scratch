package xray

type LogWriter struct {
	lastLine string
}

func NewLogWriter() *LogWriter {
	return &LogWriter{}
}
