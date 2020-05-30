package internal

import "testing"

func TestVersion(t *testing.T) {
	version := GetVersion()
	if version != "0.1" {
		t.Fatal("failed test")
	}
}
