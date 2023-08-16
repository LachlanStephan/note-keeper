package commands

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type StartCommand struct {
	RootDirs []string
}

func (s *StartCommand) Start() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	path := homeDir + "/note-keeper"
	_, err = os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("%s\n", err)
			createRootDir(path)
			return
		}

		fmt.Print("Something went wrong...\n")
		return
	}

	// get root dirs and return them
	startApp()
}

func createRootDir(path string) {
	fmt.Printf("Creating dir 'note-keeper' at %s\n", path)
	if err := os.Mkdir(path, os.ModePerm); err != nil {
		fmt.Printf("Unable to create root dir %s\n", err)
		return
	}

	fmt.Printf("Created dir: %s\n", path)

	// createRootDirs and return them
}

func createRootDirs() {
	fmt.Printf("Please enter the name(s) for the top level dirs you would like...\n")

	buf := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	rootDir, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Print("Unable to accept this input")
	}

	fmt.Print(rootDir)

	// path := "~/.config/note-keeper/" + rootDir
	// _, err := os.Create("~./config/note-keeper")
	// if err != nil {
	// 	fmt.Printf("Unable to create dir at: ~./config/note-keeper\n")
	// }
}

func startApp() {
	fmt.Print("Starting app...\n")
}
