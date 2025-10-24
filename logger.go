package main

// Logger implementation will be added step-by-step.
import ("log"
		"os"
)
type Logger struct{
	infoLogger *log.Logger
	errorLogger *log.Logger
	verbose bool
}

func NewLogger(verbose bool) *Logger{
	
}