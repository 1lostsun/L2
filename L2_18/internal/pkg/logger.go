package logger

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"time"
)

// LogLevel : Аллиас над интом для констант уровней логгирования
type LogLevel int

// Константы уровней логгирования
const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

var levelColors = map[LogLevel]*color.Color{
	DEBUG: color.New(color.FgHiCyan),
	INFO:  color.New(color.FgHiGreen),
	WARN:  color.New(color.FgHiYellow),
	ERROR: color.New(color.FgHiRed, color.Bold),
}

var currentLevel = INFO

// SetLevel : Установка уровня логгирования
func SetLevel(l LogLevel) {
	currentLevel = l
}

func log(l LogLevel, format string, args ...interface{}) {
	if l < currentLevel {
		return
	}
	timestamp := time.Now().Format("15:04:05")
	msg := fmt.Sprintf(format, args...)
	prefix := fmt.Sprintf("[%s] %-5s", timestamp, l.String())
	if c, ok := levelColors[l]; ok {
		c.Fprintf(os.Stdout, "%s %s\n", prefix, msg)
	} else {
		fmt.Fprintf(os.Stdout, "%s %s\n", prefix, msg)
	}
}

func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// Debug : Логгирование в режиме дебага
func Debug(format string, args ...interface{}) { log(DEBUG, format, args...) }

// Info : Логгирование информации
func Info(format string, args ...interface{}) { log(INFO, format, args...) }

// Warn : Логгирование опасности
func Warn(format string, args ...interface{}) { log(WARN, format, args...) }

// Error : Логгирование ошибок
func Error(format string, args ...interface{}) { log(ERROR, format, args...) }
