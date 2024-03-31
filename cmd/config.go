package main

import (
	"encoding/json"
	"fmt"
)

type ConfigValues struct {
	LastOpened string `json:"lastOpened"`
}

func (app *application) updateLastOpened(value string) error {
	fmt.Print(app.configPaths.configFilePath)

	values := &ConfigValues{LastOpened: value}
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	f, err := getFile(app.configPaths.configFilePath)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	f.Close()

	return nil
}

func (app *application) getLastOpened() (string, error) {
	return "", nil
}
