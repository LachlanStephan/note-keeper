package main

func (app *application) doSetUp() error {
	err := app.fileSystem.CreateRootDir()
	if err != nil {
		return err
	}

	err = app.fileSystem.CreateConfigFile()
	if err != nil {
		return err
	}

	return nil
}
