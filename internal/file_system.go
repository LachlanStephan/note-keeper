package internal

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
	fmt.Printf("Creating dir 'note-keeper' at %s\n", path)
	if err := os.Mkdir(path, os.ModePerm); err != nil {
		return err
	}

	fmt.Printf("Created dir: %s\n", path)
	return nil
}

func (f *FileSys) GetHomeDir() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
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
