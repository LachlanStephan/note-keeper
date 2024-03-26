package fileSystem

import (
	"errors"
	"fmt"
	"os"
)

type FileSys struct {
	RootDir      string
	NoteFile     *os.File
	NoteFilePath string
	ConfigFile   *os.File
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

func (f *FileSys) SetHomeDir() error {
	path, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Unable to get home dir: %s", path)
		return err
	}

	f.RootDir = path + "/note-keeper"
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
	fileName := "config.json"
	filePath := f.RootDir + "/" + fileName
	file, err := os.Create(filePath)
	fmt.Printf("Creating file '%s' at '%s'", fileName, filePath)
	if err != nil {
		return err
	}

	f.ConfigFile = file
	return nil
}

func (f *FileSys) CreateNoteFile() error {
	path := f.RootDir + "/note.md"

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	f.NoteFile = file
	f.NoteFilePath = path
	return nil
}
