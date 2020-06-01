package utils

import (
	"io/ioutil"
	"os"
)

var TempRoot string

func init() {
	dir, err := ioutil.TempDir("", "terrafire-")
	if err != nil {
		panic(err)
	}
	TempRoot = dir
}

func TempDir() (string, error) {
	return ioutil.TempDir(TempRoot, "")
}

func TempClean(path string) {
	err := os.RemoveAll(path)
	LogError(err)
}

func TempFile() (*os.File, error) {
	return ioutil.TempFile(TempRoot, "")
}

func RemoveTempRoot() error {
	return os.RemoveAll(TempRoot)
}
