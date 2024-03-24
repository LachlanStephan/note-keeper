package main

func run(app *application) error {
	path, err := app.fileSystem.GetHomeDir()
	if err != nil {
		app.errorLog.Print("Failed to get the home directory")
		return err
	}

	app.fileSystem.RootDir = path + "/note-keeper"

	exists, err := app.fileSystem.FileExists(app.fileSystem.RootDir)
	if err != nil {
		return err
	}
	if !exists {
		app.doSetUp()
	}

	//err = app.runTUI()
	//return err
	return nil
}
