package main

import (
	"encoding/json"
	"os"
)

type ConfigValues struct {
	LastOpened string `json:"lastOpened"`
}

func (app *application) updateLastOpened(value string) error {
	values := &ConfigValues{LastOpened: value}
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	f, err := getFile(app.configPaths.configFilePath, os.O_CREATE)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) getLastOpened() (string, error) {
	b, err := readFile(app.configPaths.configFilePath)
	if err != nil {
		return "", err
	}

	var data ConfigValues
	err = json.Unmarshal(b, &data)
	if err != nil {
		return "", err
	}

	return data.LastOpened, nil
}
