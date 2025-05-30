package logger

import (
	"log"
	"os"
)

// Структура логгера
// С отдельными логгерами для разных событий:
// - информационных
// - ошибок
// - запросов
type Logger struct {
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
	RequestLog *log.Logger
}

// New  создаёт новый экземпляр Logger
func New() *Logger {
	return &Logger{
		InfoLog:    log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog:   log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		RequestLog: log.New(os.Stdout, "REQUEST\t", log.Ldate|log.Ltime),
	}
}
