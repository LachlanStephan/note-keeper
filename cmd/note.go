package main

import (
	"os"
)

func (app *application) writeNewHeader(header string) error {
	f, err := getFile(app.configPaths.noteFilePath, os.O_APPEND)
	if err != nil {
		return err
	}

	_, err = f.WriteString(header)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) openNote() error {
	err := os.Chdir(app.configPaths.rootDirPath)
	if err != nil {
		return err
	}

	app.cmd.Name = editor
	app.cmd.Args = []string{noteFileName}
	err = app.cmd.Execute()
	if err != nil {
		return err
	}

	return nil
}
