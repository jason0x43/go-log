package log

import (
	"io"
	"log"
)

const (
	_ = iota
	LEVEL_ERROR
	LEVEL_WARN
	LEVEL_INFO
	LEVEL_DEBUG
	LEVEL_TRACE
)

var Enabled = true
var Level = LEVEL_TRACE

func Log(level int, format string, v ...interface{}) {
	if !Enabled || Level < level {
		return
	}
	log.Printf(format, v...)
}

func Trace(format string, v ...interface{}) {
	Log(LEVEL_TRACE, format, v...)
}

func Debug(format string, v ...interface{}) {
	Log(LEVEL_DEBUG, format, v...)
}

func Info(format string, v ...interface{}) {
	Log(LEVEL_INFO, format, v...)
}

func Warn(format string, v ...interface{}) {
	Log(LEVEL_WARN, format, v...)
}

func Error(format string, v ...interface{}) {
	Log(LEVEL_ERROR, format, v...)
}

// passthrough functions

func Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

func Fataln(format string, v ...interface{}) {
	log.Fatalln(v...)
}

func Flags() int {
	return log.Flags()
}

func Panic(v ...interface{}) {
	log.Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	log.Panicf(format, v...)
}

func Panicln(v ...interface{}) {
	log.Panicln(v...)
}

func Prefix() string {
	return log.Prefix()
}

func Print(v ...interface{}) {
	if !Enabled {
		return
	}
	log.Print(v...)
}

func Printf(format string, v ...interface{}) {
	if !Enabled {
		return
	}
	log.Printf(format, v...)
}

func Println(v ...interface{}) {
	if !Enabled {
		return
	}
	log.Println(v...)
}

func SetFlags(flag int) {
	log.SetFlags(flag)
}

func SetOutput(w io.Writer) {
	log.SetOutput(w)
}

func SetPrefix(prefix string) {
	log.SetPrefix(prefix)
}
