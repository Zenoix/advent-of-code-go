package utils

import (
	"path/filepath"
	"runtime"
)

func GetCurrentDir() (rootPath string) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Could not find current directory")
	}

	rootPath = filepath.Join(filepath.Dir(filename), "..")
	return
}
