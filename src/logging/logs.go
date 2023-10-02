package logging

import (
	"log"
	"os"
)

// Setting loggers for package main
// Check this utilization: https://www.honeybadger.io/blog/golang-logging/
var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	DebugLogger   *log.Logger
	ErrorLogger   *log.Logger
)

// Setup happens in init function
func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	WarningLogger = log.New(os.Stdout, "WARN: ", log.LstdFlags|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "ERROR: ", log.LstdFlags|log.Lshortfile)
	DebugLogger = log.New(os.Stdout, "DEBUG: ", log.LstdFlags)
}
