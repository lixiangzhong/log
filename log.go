package log

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
)

var (
	logger      = log.New(os.Stdout, color.New(color.FgHiBlue).Sprintf("[Info]"), log.LstdFlags)
	errorlogger = log.New(os.Stdout, color.New(color.FgRed).Sprintf("[Error]"), log.LstdFlags|log.Lshortfile)
	debuglogger = log.New(os.Stdout, color.New(color.FgHiYellow).Sprintf("[Debug]"), log.LstdFlags|log.Lshortfile)
)

type LoggerSetter interface {
	SetFlags(int)
	SetOutput(io.Writer)
	SetPrefix(string)
}

type ConfigFunc func(info, err, debug LoggerSetter)

func Set(f ConfigFunc) {
	f(logger, errorlogger, debuglogger)
}

func Println(v ...interface{}) {
	logger.Output(2, fmt.Sprint(v...))
}

func Printf(format string, v ...interface{}) {
	logger.Output(2, fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
	for i, e := range v {
		if err, ok := e.(error); ok {
			v[i] = fmt.Sprintf("%+v", errors.New(err.Error()))
		}
	}
	errorlogger.Output(2, fmt.Sprint(fmt.Sprintln(v...)))
}

func Errorf(format string, v ...interface{}) {
	errorlogger.Output(2, fmt.Sprintf(format, v...))
}

func Debug(v ...interface{}) {
	for i, e := range v {
		if err, ok := e.(error); ok {
			v[i] = fmt.Sprintf("%+v", err)
		}
	}
	debuglogger.Output(2, fmt.Sprint(fmt.Sprintln(v...)))
}

func Fatal(v ...interface{}) {
	errorlogger.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}
