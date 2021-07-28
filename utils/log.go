package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*logrus.Logger
	writer *lumberjack.Logger
}

func CreateLog(dir, name string, level logrus.Level, stdout bool) (*Logger, error) {
	var (
		err error
		out *Logger
	)
	if dir, err = getLogDir(dir); err != nil {
		return nil, err
	}
	logPath := fmt.Sprintf("%v%v.log", dir, getLogName(name))
	if out, err = NewLogger(logPath, level, stdout); err != nil {
		return nil, err
	}
	return out, nil
}

func NewLogger(path string, level logrus.Level, stdout bool) (*Logger, error) {
	out := &Logger{
		writer: &lumberjack.Logger{
			Filename:   path,
			MaxSize:    50, // megabytes
			MaxBackups: 3,
			MaxAge:     1, // days
			Compress:   true,
		},
	}

	out.Logger = logrus.New()
	out.Logger.Level = level
	out.Logger.ReportCaller = false
	out.Logger.Formatter = new(logrus.JSONFormatter)
	out.Logger.Out = out.writer
	if stdout {
		out.Logger.Out = io.MultiWriter(os.Stdout, out.writer)
	}

	return out, nil
}

func getLogDir(dir string) (string, error) {
	if len(dir) == 0 {
		dir = "/logs/"
	}
	if !strings.HasPrefix(dir, "/") {
		dir = "/" + dir
	}
	if !strings.HasSuffix(dir, "/") {
		dir = dir + "/"
	}
	file, err := filepath.Abs(os.Args[0])
	if err != nil {
		return "", err
	}
	logDir := filepath.Dir(filepath.Dir(file)+"..") + dir
	if !PathExist(logDir) {
		if err = os.MkdirAll(logDir, os.ModePerm); err != nil {
			return "", err
		}
	}
	return logDir, nil
}

func getLogName(name string) string {
	if len(name) > 0 {
		return name
	}
	return filepath.Base(os.Args[0])
}
