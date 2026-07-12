package assets

import (
	"embed"
	"io/fs"
)

//go:embed templates layout.html
var templateFS embed.FS

//go:embed static
var staticFS embed.FS

// GetTemplateFS returns embedded templates filesystem
func GetTemplateFS() embed.FS {
	return templateFS
}

// GetStaticFS returns embedded static files filesystem
func GetStaticFS() embed.FS {
	return staticFS
}

// StaticFile reads a file from the embedded static filesystem
func StaticFile(path string) ([]byte, error) {
	return fs.ReadFile(staticFS, "static/"+path)
}
