package terrafire

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestUnzip(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "")
	defer os.RemoveAll(tempDir)

	if err != nil {
		t.Fatalf("can't create temp directory: %s", err)
	}
	err = Unzip("./sample/terraform.zip", tempDir)
	if err != nil {
		t.Fatalf("unzip fail: %s", err)
	}
	_, err = os.Stat(filepath.Join(tempDir, "terraform", "test.tf"))
	if err != nil {
		t.Fatalf("file not found: %s", err)
	}
}
