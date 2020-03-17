package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"time"
)

type HLog struct {
	tag string
	out io.Writer
	log *log.Logger
}

type LOG_LEVEL uint
type OUT_FOMATTER uint
const (
	TRACE LOG_LEVEL = 0
	INFO LOG_LEVEL = 1
	WARN LOG_LEVEL = 2
	ERROR LOG_LEVEL = 3
	FATAL LOG_LEVEL = 4
	PANIC LOG_LEVEL = 5

	TEXT_FORMAT OUT_FOMATTER = 0
	JSON_FORMAT OUT_FOMATTER = 1
)

var (
	_level LOG_LEVEL = INFO
	_formatter = TEXT_FORMAT
)

func New(tag string, dir string) *HLog {
	logrus.Panicf()
	var out io.Writer
	if len(dir) > 0 {
		fileName := time.Now().Format("2006-01-02") + ".log"
		file, err := os.OpenFile(dir+ string(os.PathSeparator) + fileName, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		out = io.MultiWriter(file, os.Stdout)
	} else {
		out = os.Stdout
	}

	return &HLog{ tag, out, log.New(out, tag + " ", log.Ldate | log.Ltime | log.Lshortfile)}
}

func SetLevel(level LOG_LEVEL) {
	_level = level
}

func SetFormatter(fomatter OUT_FOMATTER) {
	_formatter = fomatter
}

func (l *HLog) Trace(msg string) {
	l.print(TRACE, msg,nil)
}

func (l *HLog) Tracef(format string, a ...interface{}) {
	l.print(TRACE, fmt.Sprintf(format, a...), nil)
}

func (l *HLog) Info(msg string) {
	l.print(INFO, msg, nil)
}

func (l *HLog) Infof(format string, a ...interface{}) {
	l.print(INFO, fmt.Sprintf(format, a...), nil)
}

func (l *HLog) Warn(msg string) {
	l.print(WARN, msg, nil)
}

func (l *HLog) Warnf(format string, a ...interface{}) {
	l.print(WARN, fmt.Sprintf(format, a...), nil)
}

func (l *HLog) Error(msg string) {
	l.print(ERROR, msg, nil)
}

func (l *HLog) Errorf(format string, a ...interface{}) {
	l.print(ERROR, fmt.Sprintf(format, a...), nil)
}

func (l *HLog) Fatal(msg string) {
	l.print(FATAL, msg, nil)
}

func (l *HLog) Fatalf(format string, a ...interface{}) {
	l.print(FATAL, fmt.Sprintf(format, a...), nil)
}

func (l *HLog) Panic(msg string, err error) {
	l.print(PANIC, msg, err)
}

func (l *HLog) Panicf(format string, a ...interface{}) {
	l.print(PANIC, fmt.Sprintf(format, a...), nil)
}

func (l *HLog) print(level LOG_LEVEL, msg string, err error) {
	if level < _level {
		return
	}

	switch level {
		case TRACE:
			l.log.Printf("[TRACE]: %s", msg)
		case INFO:
			l.log.Printf("[INFO]: %s", msg)
		case WARN:
			l.log.Printf("[WARN]: %s", msg)
		case ERROR:
			l.log.Printf("[ERROR]: %s", msg)
		case FATAL:
			l.log.Fatalf("[FATAL]: %s", msg)
		case PANIC:
			var msgTmp string
			if err == nil {
				msgTmp = fmt.Sprintf("[PANIC]: %s", msg)
			} else {
				msgTmp = fmt.Sprintf("[PANIC]: %s\n\t%s", msg, err)
			}
			l.log.Panicf(msgTmp)
	}
}