package internal

import (
	"testing"
)

func TestFileRead(t *testing.T) {
	content, err := FileRead("../docs/template.ghp")
	if err != nil {
		t.Fatalf("FileRead() returned an error: %v", err)
	}

	if content == "" {
		t.Fatal("FileRead() returned empty content")
	}
}
