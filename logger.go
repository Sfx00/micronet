papackage main

import (
	"log"
	"os"
)

type Logger struct{
	infoLogger *log.Logger
	errorLogger *log.Logger
	verbose bool
}


func NewLogger(verbose bool) *Logger{
	var logger Logger

	logger.infoLogger = log.New(os.Stdout, "INFO: ",log.LstdFlags)
	logger.errorLogger = log.New(os.Stderr, "ERROR: ", log.LstdFlags)
	logger.verbose = verbose
	return  &logger
}

func (l *Logger) Info(msg string){
	l.infoLogger.Println(msg)
}

func (l *Logger) Error(err error){
	l.errorLogger.Println(err)
}

func (l *Logger) Verbose(msg string){
	if(!l.verbose){
		return
	}else{
		l.infoLogger.Println(msg)
	}
}
