package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	LogPath string
}

func (l *Logger) Log(msg string) error {
	return logFile(msg, l.LogPath)
}

func (l *Logger) Error(msg error) error {
	return errorFile(msg, l.LogPath)
}

func logFile(msg string, fileLog string) error {
	file, err := os.OpenFile(fileLog, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	log.SetOutput(file)
	log.Println(msg)
	return nil
}

func errorFile(msg error, fileLog string) error {
	file, err := os.OpenFile(fileLog, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Error(err)
		return err
	}
	log.SetOutput(file)
	log.Error(msg)
	return nil
}
