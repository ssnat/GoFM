package modules

import (
	"fmt"
	"time"
)

type ILogger struct{}

var Logger ILogger

func init() {
	Logger = ILogger{}
}

func (logger *ILogger) Println(logType string, a ...any) {
	date := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] %s\n", date, logType, fmt.Sprint(a...))
}

func (logger *ILogger) Info(a ...interface{}) {
	logger.Println("info", a...)
}

func (logger *ILogger) Debug(a ...any) {
	config := GetConfig()
	if config.Debug {
		logger.Println("debug", a...)
	}
}

func (logger *ILogger) Error(a ...any) {
	logger.Println("error", a...)
}
