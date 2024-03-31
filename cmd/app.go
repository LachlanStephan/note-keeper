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

	// check first probably
	currentTime := time.SetTimes().FormattedDate
	lastOpened, err := app.getLastOpened()
	if err != nil {
		return err
	}

	if currentTime != lastOpened {
		err = app.updateLastOpened(time.SetTimes().FormattedDate)
		if err != nil {
			return err
		}
		err = app.writeNewHeader(currentTime)
		if err != nil {
			return err
		}
	}

	return nil
}
