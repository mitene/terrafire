package terrafire

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestUnzip(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("can't create temp directory: %s", err)
	}
	defer os.RemoveAll(tempDir)

	file, err := os.Open("./sample/terraform.zip")
	if err != nil {
		t.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	client := &GithubClientImpl{}
	err = client.extract(file, "subdir", tempDir)
	if err != nil {
		t.Fatalf("unzip fail: %s", err)
	}

	_, err = os.Stat(filepath.Join(tempDir, "test.tf"))
	if err != nil {
		t.Fatalf("file not found: %s", err)
	}
}
