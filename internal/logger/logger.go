package logger

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "[Wh-Api] ", log.Ldate|log.Ltime|log.Lshortfile)
