package axiom

import (
	"log"
	"os"
)

// A simple logger
type Log struct {
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

var Logger = NewLogger()

// Returns a new logger
func NewLogger() *Log {
	return &Log{
		Info:    log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime),
		Warning: log.New(os.Stdout, "WARN ", log.Ldate|log.Ltime),
		Error:   log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime),
	}
}
