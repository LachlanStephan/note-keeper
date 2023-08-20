package main

import (
	"log"
	"os"

	internal "github.com/LachlanStephan/note-keeper/internal"
)

type application struct {
	errorLog   *log.Logger
	infoLog    *log.Logger
	fileSystem *internal.FileSys
}

var (
	infoLog  = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)
)

func main() {
	app := &application{
		errorLog:   errorLog,
		infoLog:    infoLog,
		fileSystem: &internal.FileSys{},
	}

	err := run(app)
	errorLog.Fatal(err)
}
