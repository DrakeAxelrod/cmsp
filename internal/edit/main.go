package edit

import (
	"io/ioutil"
	"log"
	"os"
)

// function that edits like sed -i
func EditFile(filename string, editFunc func([]byte) []byte) error {
	// read the file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	// get file permissions
	fi, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	// edit the file
	data = editFunc(data)

	// write the file
	err = ioutil.WriteFile(filename, data, fi.Mode())
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
