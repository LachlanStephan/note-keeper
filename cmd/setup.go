package main

func (app *application) createScaffold() error {
	err := app.fileSystem.CreateRootDir()
	if err != nil {
		return err
	}

	err = app.fileSystem.CreateConfigFile()
	if err != nil {
		return err
	}

	err = app.fileSystem.CreateNoteFile()
	if err != nil {
		return err
	}

	return nil
}
