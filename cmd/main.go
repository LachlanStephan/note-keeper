package main

import (
	"log"
	"os"
)

type configPaths struct {
	noteFilePath   string
	noteFile       *os.File
	configFilePath string
	configFile     *os.File
	rootDirPath    string
}

type application struct {
	errorLog    *log.Logger
	infoLog     *log.Logger
	configPaths *configPaths
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
	}

	err := run(app)
	if err != nil {
		errorLog.Fatal(err)
	}
}
