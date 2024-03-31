package main

import (
	"fmt"

	time "github.com/LachlanStephan/note-keeper/internal/time"
)

func run(app *application) error {
	rootDir, err := getHomeDir()
	if err != nil {
		return err
	}
	app.configPaths.rootDirPath = rootDir
	app.configPaths.configFilePath = app.getConfigFilePath()
	app.configPaths.noteFilePath = app.getNoteFilePath()

	exists, err := fileExists(app.configPaths.rootDirPath)
	if err != nil {
		return err
	}
	if !exists {
		fmt.Print("Config does not exists... scaffolding application\n")
		app.createScaffold()
	}

	err = app.updateLastOpened(time.SetTimes().FormattedDate)
	if err != nil {
		return err
	}

	/**
	if we are to now open the note - we need to check the config file for last opened

	then get the current time - if it === last opened -> open the note
	else
	write a new date heading then open the note
	*/

	return nil
}
