package main

import (
	"os"
	"path/filepath"
)

func getFilesList(inputDir string) []string {
	out := []string{}

	filepath.Walk(inputDir, func(path string, f os.FileInfo, err error) error {

		if !f.IsDir() {
			out = append(out, path)
		}

		return nil
	})

	return out
}
