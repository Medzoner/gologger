package logger_test

import (
	"errors"
	"fmt"
	"github.com/Medzoner/gologger"
	"os"
	"testing"
)

func TestLoggerSuccess(t *testing.T) {
	pwd, _ := os.Getwd()
	loggerTest := logger.Logger{LogPath: pwd + "/var/log/test.log"}

	t.Run("Unit: test logger log success", func(t *testing.T) {
		err := loggerTest.Log("good test")
		if err != nil {
			fmt.Println("error on log")
		}
	})
}

func TestLoggerFailed(t *testing.T) {
	pwd, _ := os.Getwd()
	loggerTest := logger.Logger{LogPath: pwd + "/fake/test.log"}

	t.Run("Unit: test logger log failed", func(t *testing.T) {
		err := loggerTest.Log("test return error")
		if err != nil {
			fmt.Println("no return error")
		}
	})
}

func TestLoggerErrorSuccess(t *testing.T) {
	pwd, _ := os.Getwd()
	loggerTest := logger.Logger{LogPath: pwd + "/var/log/test.log"}

	t.Run("Unit: test logger log error success", func(t *testing.T) {
		_ = loggerTest.Error(errors.New("this is a log error"))
	})
}

func TestLoggerErrorFailed(t *testing.T) {
	pwd, _ := os.Getwd()
	loggerTest := logger.Logger{LogPath: pwd + "/fake/test.log"}

	t.Run("Unit: test logger log error failed", func(t *testing.T) {
		_ = loggerTest.Error(errors.New("this is a log error"))
	})
}