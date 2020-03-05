package terrafire

import (
	"io/ioutil"
	"testing"
)

func TestUnzip(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("can't create temp directory: %s", err)
	}
	err = Unzip("./sample/terraform.zip", tempDir)
	if err != nil {
		t.Fatalf("unzip fail: %s", err)
	}
}