package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func RemoveFolder(folderPath string) error {
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		return err
	}

	err = os.RemoveAll(folderPath)
	return err
}

func CreateFolder(folderPath string) error {
	// Create folder is it doesn't exists
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, 0755)
		if err != nil {
			return err
		}
	}
	return err
}

func CreateFile(filePath string) error {

	filePath, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}

	// Attempt to create a file
	res := filepath.Dir(filePath)
	err = CreateFolder(res)
	if err != nil {
		return err
	}

	// Create the file
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		// err = os.WriteFile(filePath, []byte("Hello World!"), os.ModeAppend.Perm())
		_, err = os.Create(filePath)
		fmt.Println(err)
		// if os.IsNotExist(err) {
		// 	return nil
		// }
		return err
	}

	return nil
}
