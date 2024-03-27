package main

func run(app *application) error {
	err := app.fileSystem.SetHomeDir()
	if err != nil {
		app.errorLog.Print("Failed to get the home directory")
		return err
	}

	exists, err := app.fileSystem.FileExists(app.fileSystem.RootDir)
	if err != nil {
		return err
	}
	if !exists {
		app.createScaffold()
	}

	err = app.UpdateLastOpenedIfNeeded()
	if err != nil {
		return err
	}

	/*
		TODO
			create a key/value in config for tracking the date the file was last opened - can use this to see if we need a new date or not
	*/

	return nil
}
