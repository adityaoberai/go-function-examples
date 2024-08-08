package services

import (
	"os"
	"path/filepath"

	"github.com/open-runtimes/types-for-go/v4/openruntimes"
)

func GetStaticFile(Context openruntimes.Context, fileName string) string {
	// Get the directory of the current executable
	ex, err := os.Executable()
	if err != nil {
		Context.Error(err)
	}
	exPath := filepath.Dir(ex)

	// Build the path to the static folder
	staticFolder := filepath.Join(exPath, "./static")

	// Build the full path to the file
	filePath := filepath.Join(staticFolder, fileName)

	// Read the file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		Context.Error(err)
	}

	return string(content)
}
