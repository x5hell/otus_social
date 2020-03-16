package file

import (
	"component/handler"
	"fmt"
	"io/ioutil"
	"os"
)

func GetContent(filePath string) (content string, err error) {
	fileBuffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		handler.ErrorLog(err)
		return "", err
	}
	return string(fileBuffer), nil
}

func PutContent(filePath string, content string, fileAccess int) error {
	file, err := os.OpenFile(filePath, fileAccess, 0644)
	if err != nil {
		handler.ErrorLog(err)
		return err
	}
	_, err = fmt.Fprintln(file, content)
	if err != nil {
		handler.ErrorLog(err)
		return err
	}
	err = file.Close()
	if err != nil {
		handler.ErrorLog(err)
		return err
	}
	return nil
}

func WriteList(filePath string, stringList []string) (err error) {
	for num, value := range stringList {
		if num == 0 {
			err = PutContent(filePath, value, os.O_CREATE|os.O_TRUNC|os.O_WRONLY)
		} else {
			err = PutContent(filePath, value, os.O_APPEND|os.O_WRONLY)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
