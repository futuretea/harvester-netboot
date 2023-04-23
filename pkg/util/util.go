package util

import (
	"io/ioutil"
)

func CreateFileWithContent(filename string, content string) error {
	data := []byte(content)

	return ioutil.WriteFile(filename, data, 0644)
}
