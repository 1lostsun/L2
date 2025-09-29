package logger

import (
	"github.com/fatih/color"
	"log"
	"strings"
)

// Logger : Структура логгера
type Logger struct{}

// New : Конструктор структуры логгера
func New() *Logger {
	return &Logger{}
}

// Log : Функция логгирования
func (logger *Logger) Log(newLine bool, message ...string) {
	msg := strings.Join(message, " ")
	Color := color.New(color.FgHiWhite).Add(color.Bold)

	if newLine {
		_, err := Color.Println(msg)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, err := Color.Print(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}
