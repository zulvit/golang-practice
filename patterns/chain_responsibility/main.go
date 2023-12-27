package main

import "fmt"

// Logger интерфейс с методом Log, который принимает сообщение и его уровень.
type Logger interface {
	Log(level string, message string) bool
}

// Создадим конкретные обработчики для сообщений разного уровня

type InfoLogger struct {
	next Logger
}

type WarningLogger struct {
	next Logger
}

type ErrorLogger struct {
	next Logger
}

// Реализуем логеры

func (l *InfoLogger) Log(level, message string) bool {
	if level == "INFO" {
		fmt.Printf("INFO: %s\n", message)
		return true
	} else if l.next != nil {
		return l.next.Log(level, message)
	}
	return false
}

func (l *WarningLogger) Log(level, message string) bool {
	if level == "WARNING" {
		fmt.Printf("Warning: %s\n", message)
		return true
	} else if l.next != nil {
		return l.next.Log(level, message)
	}
	return false
}

func (l *ErrorLogger) Log(level, message string) bool {
	if level == "ERROR" {
		fmt.Printf("Error: %s\n", message)
		return true
	} else if l.next != nil {
		return l.next.Log(level, message)
	}
	return false
}

func main() {
	infoLogger := &InfoLogger{}
	warningLogger := &WarningLogger{}
	errorLogger := &ErrorLogger{}

	infoLogger.next = warningLogger
	warningLogger.next = errorLogger

	infoLogger.Log("INFO", "Это информационное сообщение")
	infoLogger.Log("WARNING", "Это предупреждение")
	infoLogger.Log("ERROR", "Это сообщение об ошибке")
}
