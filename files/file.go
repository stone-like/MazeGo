package files

import (
	"os"
	"path/filepath"
)

func GetFilePath(fileName string) string {
	path, _ := os.Getwd()
	path = filepath.Join(path, "files")
	return filepath.Join(path, fileName)
}
