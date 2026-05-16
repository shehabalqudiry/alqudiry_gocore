package utils

import (
	"os"
	"path/filepath"
)

func CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func WriteFile(path string, content string) error {

	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(
		path,
		[]byte(content),
		0644,
	)
}

func PathJoin(elem ...string) string {
	return filepath.Join(elem...)
}
