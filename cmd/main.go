package main

import (
	"log"
	"os"

	commands "github.com/LachlanStephan/note-keeper/commands"
)

type Cmd struct {
	Name   string
	Args   []string
	Stdin  *os.File
	Stdout *os.File
	Stderr *os.File
	start  *commands.StartCommand
}

type Application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	cmd      *Cmd
}

var (
	infoLog  = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)
)

func main() {
	app := &Application{
		infoLog:  infoLog,
		errorLog: errorLog,
		cmd:      &Cmd{},
	}

	app, err := newApp(app)
	if err != nil {
		errorLog.Fatal("Could not start application")
	}
}

func newApp(app *Application) (*Application, error) {
	app.cmd.start.Start()
	return app, nil
}
