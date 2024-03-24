package main

import (
	"log"
	"os"

	fileSystem "github.com/LachlanStephan/note-keeper/internal/filesystem"
)

type application struct {
	errorLog   *log.Logger
	infoLog    *log.Logger
	fileSystem *fileSystem.FileSys
}

var (
	infoLog  = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)
)

func main() {
	app := &application{
		errorLog:   errorLog,
		infoLog:    infoLog,
		fileSystem: &fileSystem.FileSys{},
	}

	err := run(app)
	if err != nil {
		errorLog.Fatal(err)
	}
}
