package utils

import (
	"path/filepath"
	"runtime"
)

func GetCurrentDirectory() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filename)
}
