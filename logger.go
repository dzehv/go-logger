/*
  Golang logging package with numeric levels

  funcs:
      SetDebugLvl(int) --  set debug level
      SetMsgPrefix(string) -- set all log messages prefix
      UseStderr(bool) -- if true log to stderr instead of stdout
      Debug(int, string, ...interface{}) -- log debug into including caller func name and line number
      Info(...)
      Warning(...)
      Error(...)
      Fatal(...)
      Panic(...) -- log panic message and than panic() with trace
*/

package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
)

// global debug level
var debugLevel int

// log message prefix string
var msgPrefix string

// our pid to log
var pids string = strconv.Itoa(os.Getpid())

// init
func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile /* | log.LUTC */)
}

// set/change debug level
func SetDebugLvl(level int) {
	debugLevel = level
}

// set message prefix,
// will go right after log header (ts, pid, etc) and before actual message
func SetMsgPrefix(prefix string) {
	msgPrefix = prefix
}

// toggle output to stderr or stdout
func UseStderr(use bool) {
	if use {
		log.SetOutput(os.Stderr)
	} else {
		log.SetOutput(os.Stdout)
	}
}

// print actual log message
func printMsg(lvl string, msg string, more ...interface{}) {
	log.Printf("pid:"+pids+" ["+lvl+"]"+msgPrefix+" "+msg, more...)
}

// print debug with some simple trace
func Debug(level int, msg string, more ...interface{}) {
	if level > debugLevel {
		return
	}

	pc, _, no, _ := runtime.Caller(1)
	de := runtime.FuncForPC(pc)
	var dlev = fmt.Sprintf("D%d", level)
	var format = fmt.Sprintf("@%s:%d %s", de.Name(), no, msg)

	printMsg(dlev, format, more...)
}

// print info message
func Info(msg string, more ...interface{}) {
	printMsg("I", msg, more...)
}

// print error message
func Error(msg string, more ...interface{}) {
	printMsg("E", msg, more...)
}

// print warning message
func Warning(msg string, more ...interface{}) {
	printMsg("W", msg, more...)
}

// print fatal message
func Fatal(msg string, more ...interface{}) {
	printMsg("F", msg, more...)
}

// print Panic message and... well, PANIC!
func Panic(msg string, more ...interface{}) {
	printMsg("P", msg, more...)
	panic("more")
}
