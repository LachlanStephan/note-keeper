package fileSystem

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	time "github.com/LachlanStephan/note-keeper/internal/time"
)

type FileSys struct {
	RootDir      string
	NoteFile     *os.File // needs to be in app struct
	NoteFilePath string   // needs to be in app struct
	ConfigFile   *os.File // needs to be in app struct
	ConfigPath   string   // needs to be in app struct
}

type ConfigData struct {
	LastOpened string
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
	f.ConfigPath = filePath
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

func (f *FileSys) GetLastOpened() (string, error) {
	configData := ConfigData{}
	fileBytes, _ := os.ReadFile(f.ConfigPath)
	err := json.Unmarshal(fileBytes, &configData)
	if err != nil {
		return "", err
	}

	return configData.LastOpened, nil
}

func (f *FileSys) UpdateLastOpened() error {
	fmt.Print("update last opened")
	data := ConfigData{
		LastOpened: time.SetTimes().FormattedDate,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	fmt.Print("past here 1")
	fmt.Print("Path" + f.ConfigPath)

	err = os.WriteFile(f.ConfigPath, b, os.ModePerm)
	if err != nil {
		return err
	}

	fmt.Print("past here 2")

	return nil
}
