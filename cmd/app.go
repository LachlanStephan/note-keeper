package main

func run(app *application) error {
	path, err := app.fileSystem.GetHomeDir()
	if err != nil {
		return err
	}

	app.fileSystem.RootDir = path + "/note-keeper"

	_, err = app.fileSystem.FileExists(app.fileSystem.RootDir)

	if err != nil {
		return err
	}

	err = app.runTUI()
	return err
}
