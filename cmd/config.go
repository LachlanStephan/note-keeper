package main

import (
	"fmt"

	time "github.com/LachlanStephan/note-keeper/internal/time"
)

func (app *application) UpdateLastOpenedIfNeeded() error {
	if app.fileSystem.ConfigFile == nil {
		err := app.fileSystem.UpdateLastOpened()
		if err != nil {
			return err
		}
		return nil
	}

	lastOpened, err := app.fileSystem.GetLastOpened()
	if err != nil {
		return err
	}

	fmt.Print("Last opened " + lastOpened)

	currTime := time.SetTimes().FormattedDate

	if lastOpened == currTime {
		return nil
	}

	err = app.fileSystem.UpdateLastOpened()
	if err != nil {
		return err
	}

	return nil
}
