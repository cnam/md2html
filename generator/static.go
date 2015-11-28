package generator

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
	"errors"
)

//Static file represents need copy
type StaticFile struct {
	Body []byte
	Name string
}

// NewStatic creating new static file
func (d *Dir) NewStatic(f os.FileInfo) (*StaticFile, error) {
	body, err := ioutil.ReadFile(getPath(d.mdDir, f.Name()))

	if strings.HasPrefix(f.Name(), "_") {
		return nil, errors.New(fmt.Sprintf("Not allowed file %s\n", f.Name()))
	}

	if (err != nil) {
		return nil, err
	}

	return &StaticFile{Body:body, Name:f.Name()}, nil
}

// write writing content to file
func (f *StaticFile) write(d *Dir) error {
	fi, err := os.Create(getPath(d.htmlDir, f.Name))

	if err != nil {
		return err
	}

	_, err = fi.Write(f.Body)

	if err != nil {
		return err
	}

	fmt.Printf("Create new static file: %s\n", getPath(d.htmlDir, f.Name))

	return nil

}