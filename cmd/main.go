package main

import (
	"log"
	"os"

	exec "github.com/LachlanStephan/note-keeper/internal/exec"
)

type configPaths struct {
	noteFilePath   string
	configFilePath string
	rootDirPath    string
}

type application struct {
	errorLog    *log.Logger
	infoLog     *log.Logger
	configPaths *configPaths
	cmd         *exec.Command
}

var (
	infoLog  = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)
)

func main() {
	app := &application{
		errorLog:    errorLog,
		infoLog:     infoLog,
		configPaths: &configPaths{},
		cmd:         &exec.Command{},
	}

	err := run(app)
	if err != nil {
		errorLog.Fatal(err)
	}
}
