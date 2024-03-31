package main

import (
	"errors"
	"fmt"
	"os"
)

func createFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	fmt.Printf("Creating file at '%s'\n", path)

	f.Close()

	return nil
}

func createRootDir(path string) error {
	if err := os.Mkdir(path, os.ModePerm); err != nil {
		return err
	}

	fmt.Printf("Creating directory at '%s'\n", path)

	return nil
}

func getHomeDir() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Unable to get home dir: %s\n", path)
		return "", err
	}

	return path + "/" + appRootDir, nil
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func getFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}
