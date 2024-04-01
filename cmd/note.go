package main

import "os"

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
