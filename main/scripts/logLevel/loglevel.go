package logLevel

import (
	"log"
	"os"
)

const (
	flag       = log.Ldate | log.Ltime
	preDebug   = "[DEBUG] "
	preInfo    = "[INFO] "
	preWarning = "[WARNING] "
	preError   = "[ERROR] "
)

var (
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
)

func init() {
	debugLogger = log.New(os.Stderr, preDebug, flag)
	infoLogger = log.New(os.Stderr, preInfo, flag)
	warningLogger = log.New(os.Stderr, preWarning, flag)
	errorLogger = log.New(os.Stderr, preError, flag)
}

func Debugf(format string, v ...interface{}) {
	debugLogger.Printf(format, v...)
}

func Infof(format string, v ...interface{}) {
	infoLogger.Printf(format, v...)
}

func Warningf(format string, v ...interface{}) {
	warningLogger.Printf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	errorLogger.Printf(format, v...)
}
