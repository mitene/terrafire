package main

import (
	"fmt"
	"io/ioutil"

	"github.com/mitene/terrafire"
)

func main() {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		return
	}
	fmt.Println(tempDir)
	terrafire.GetSource("maychannel-dev", "terraform", "master", tempDir)
}
