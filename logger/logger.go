package logger

import (
	"fmt"
	"time"

	"github.com/op/go-logging"
)

var (
	logger    *logging.Logger
	logBuffer []struct {
		time  string
		level logging.Level
		log   string
	}
)

func InitLogger(level logging.Level) {
	newLogger := logging.MustGetLogger("scratch-x-ui")
	logger = newLogger
}

func Info(args ...interface{}) {
	logger.Info(args...)
	addToBuffer("INFO", fmt.Sprint(args...))
}

func addToBuffer(level string, newLog string) {
	t := time.Now()
	if len(logBuffer) >= 10240 {
		logBuffer = logBuffer[1:]
	}

	logLevel, _ := logging.LogLevel(level)
	logBuffer = append(logBuffer, struct {
		time  string
		level logging.Level
		log   string
	}{
		time:  t.Format("2006/01/02 15:04:05"),
		level: logLevel,
		log:   newLog,
	})
}

func Warningf(format string, args ...interface{}) {
	logger.Warningf(format, args...)
	addToBuffer("WARNING", fmt.Sprintf(format, args...))
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
	addToBuffer("INFO", fmt.Sprintf(format, args...))
}

func Warning(args ...interface{}) {
	logger.Warning(args...)
	addToBuffer("WARNING", fmt.Sprint(args...))
}
