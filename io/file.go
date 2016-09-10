package io

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

// WriteToFile writes string to a file in the specified path
func WriteToFile(content, path string) {
	fp, _ := filepath.Abs(path)
	bcontent := []byte(content)

	err := ioutil.WriteFile(fp, bcontent, 0755)

	if err != nil {
		log.Println(err)
	}
}
