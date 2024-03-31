package main

func (app *application) createScaffold() error {
	err := createRootDir(app.configPaths.rootDirPath)
	if err != nil {
		return err
	}

	paths := [2]string{
		app.configPaths.configFilePath,
		app.configPaths.noteFilePath,
	}

	for _, v := range paths {
		err = createFile(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *application) getConfigFilePath() string {
	return app.configPaths.rootDirPath + "/" + configFileName
}

func (app *application) getNoteFilePath() string {
	return app.configPaths.rootDirPath + "/" + noteFileName
}
