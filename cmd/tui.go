package main

func (app *application) runTUI() error {
	err := app.fileSystem.SetRootDirs()
	if err != nil {
		return err
	}

	// create TUI here

	//
	return nil
}
