package main

import (
	"log"
	"os"
	"strings"
	"fmt"
	"io/ioutil"
)

type TypeLogger struct {
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	LogFile *log.Logger
}

func NewLogger(source string) *TypeLogger {
	var pathLogFile string
	prod := os.Getenv("PROD")
	if strings.Compare(prod, "1") == 0 {
		pathLogFile = fmt.Sprintf("./logs_vol/%s.log", source)
	} else {
		pathLogFile = fmt.Sprintf("./app_logs/%s.log", source)
	}

	file, err := os.OpenFile(pathLogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open logging file", ":", err)
	}

	l := &TypeLogger{
		Trace: log.New(ioutil.Discard,
			"TRACE: ",
			log.Ldate|log.Ltime|log.Lshortfile),
		Info: log.New(os.Stdout,
			"INFO: ",
			log.Ldate|log.Ltime|log.Lshortfile),
		Warning: log.New(os.Stdout,
			"WARNING: ",
			log.Ldate|log.Ltime|log.Lshortfile),
		Error: log.New(os.Stderr,
			"ERROR: ",
			log.Ldate|log.Ltime|log.Lshortfile),
	}

	l.Info.SetOutput(file)
	l.Trace.SetOutput(file)
	l.Info.SetOutput(file)
	l.Warning.SetOutput(file)
	l.Error.SetOutput(file)

	return l
}
