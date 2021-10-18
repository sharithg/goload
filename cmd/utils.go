package cmd

import "os"

// check if a given path exists
func PathExits(path string) (bool, error) {
	fileOrDir, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	switch mode := fileOrDir.Mode(); {
	case mode.IsDir():
		// do directory stuff
		return true, nil
	case mode.IsRegular():
		return false, nil
	}
	return false, err
}
