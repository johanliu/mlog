package mlog

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

var DefaultLevel int = LevelInfo
var DefaultOutput io.Writer = os.Stdout
var DefaultFlag int = log.Ldate | log.Ltime | log.Lshortfile

const (
	LevelError = iota
	LevelWarning
	LevelInfo
	LevelDebug
)

var levelName = map[string]int{
	"ERROR":   LevelError,
	"WARNING": LevelWarning,
	"INFO":    LevelInfo,
	"DEBUG":   LevelDebug,
}

type Logger struct {
	level int
	err   *log.Logger
	war   *log.Logger
	inf   *log.Logger
	deb   *log.Logger
}

func (l *Logger) Error(err error) {
	if LevelError > l.level {
		return
	}

	s := err.Error()
	l.err.Output(2, s)
	panic(s)
}

func (l *Logger) Warning(format string, v ...interface{}) {
	if LevelWarning > l.level {
		return
	}

	l.war.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(format string, v ...interface{}) {
	if LevelInfo > l.level {
		return
	}

	l.inf.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(format string, v ...interface{}) {
	if LevelDebug > l.level {
		return
	}

	l.deb.Output(2, fmt.Sprintf(format, v...))
}

var mutex = new(sync.RWMutex)

func (l *Logger) SetLevel(level int) {
	mutex.Lock()
	defer mutex.Unlock()
	l.level = level
}

func (l *Logger) SetLevelByName(level string) {
	mutex.Lock()
	defer mutex.Unlock()
	l.level = levelName[strings.ToUpper(level)]
}

func (l *Logger) Level() int {
	mutex.RLock()
	defer mutex.RUnlock()
	return l.level
}

func NewLogger() *Logger {
	l := &Logger{
		level: DefaultLevel,
	}

	l.err = log.New(DefaultOutput, "[ERROR]: ", DefaultFlag)
	l.war = log.New(DefaultOutput, "[WARNING]: ", DefaultFlag)
	l.inf = log.New(DefaultOutput, "[INFO]: ", DefaultFlag)
	l.deb = log.New(DefaultOutput, "[DEBUG]: ", DefaultFlag)

	return l
}
