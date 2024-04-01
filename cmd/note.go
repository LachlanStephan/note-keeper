package main

import (
	"os"
	"os/exec"
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

	cmd := exec.Command(editor, noteFileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
