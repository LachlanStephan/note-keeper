package fileSystem

import (
	"errors"
	"fmt"
	"os"
)

type FileSys struct {
	RootDir  string
	RootDirs []string
}

func (f *FileSys) CreateRootDir() error {
	path := f.RootDir
	fmt.Printf("Creating dir 'note-keeper' at '%s'", path)
	if err := os.Mkdir(path, os.ModePerm); err != nil {
		fmt.Printf("Unable to create file at '%s'", path)
		return err
	}

	return nil
}

func (f *FileSys) GetHomeDir() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Unable to get home dir: %s", path)
		return "", err
	}

	return path, nil
}

func (f *FileSys) SetRootDirs() error {
	// get all dirs under f.RootDir
	// and set them in f.RootDirs
	return nil
}

func (f *FileSys) FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (f *FileSys) CreateConfigFile() error {
	file := "config.json"
	filePath := f.RootDir + "/" + file
	_, err := os.Create(filePath)
	fmt.Printf("Creating file '%s' at '%s'", file, filePath)
	if err != nil {
		return err
	}

	return nil
}
