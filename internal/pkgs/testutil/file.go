package testutil

import (
	"io"
	"os"
	"testing"
)

func ReadFileAsString(t *testing.T, path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("error while reading file with path = %s, %s", path, err)
	}
	return string(content)
}

func ReadFileAsReader(t *testing.T, path string) io.ReadCloser {
	content, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Errorf("error while reading file with path = %s, %s", path, err)
	}
	return content
}
