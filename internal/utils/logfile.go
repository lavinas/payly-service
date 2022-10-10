package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/lavinas/payly-service/internal/core/ports"
)

type logFile struct {
	path      string
	component string
	file      *os.File
	date      string
	info      bool
}

func NewlogFile(c ports.Config, component string) *logFile {
	p, b := getVars(c)
	f, d := initFile(p, component)
	return &logFile{path: p, file: f, component: component, date: d, info: b}
}

func (l *logFile) GetFile() *os.File {
	return l.file
}

func (l *logFile) Info(message string) {
	if l.info {
		write(l, message)
	}
}

func (l *logFile) Error(message string) {
	write(l, message)
}

func write(l *logFile, message string) {
	shiftFile(l)
	d := time.Now().Format("2006/01/02 - 15:04:05")
	t := "[LOG] " + d + " | " + message + "\n"
	l.file.Write([]byte(t))
	l.file.Sync()
}

func initFile(path string, component string) (*os.File, string) {
	d := time.Now().Format("2006-01-02")
	fp := path + "/" + d + "-" + component + ".log"
	f, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic("creating log file error")
	}
	return f, d
}

func getVars(c ports.Config) (string, bool) {
	p, err := c.GetField("log", "path")
	if err != nil {
		panic("No path log configured")
	}
	i, err := c.GetField("log", "info")
	if err != nil {
		panic("No path info configured")
	}
	b, err := strconv.ParseBool(i)
	if err != nil {
		panic("Path info is not boolean type")
	}
	return p, b
}

func shiftFile(l *logFile) {
	d := time.Now().Format("2006-01-02")
	if d != l.date {
		l.file.Close()
		l.file, l.date = initFile(l.path, l.component)
	}
}
