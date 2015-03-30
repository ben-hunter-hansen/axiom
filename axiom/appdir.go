package axiom

import (
	"os"
	"path/filepath"
	"strings"
)

// Models the applications directory.
// 1. Name of root view directory
// 2. Name of root resource directory
// 3. Map of view directories
// 	  k: directory (named after controller)
//    v: full path to view file
// 4. Map of resource directories
// 	  k: directory
//    v: full path to resource
// 5. Valid file extensions
type AppDirectory struct {
	ViewRoot     string
	ResourceRoot string
	Views        map[string][]string
	Resources    map[string][]string
	Extensions   []string
}

// Traverses and maps view directory
func (a *AppDirectory) viewWalker(path string, fi os.FileInfo, err error) error {
	if fi.IsDir() && path != a.ViewRoot {
		a.Views[fi.Name()] = make([]string, 0)
	} else {
		pathIgnoreRoot := strings.TrimPrefix(path, a.ViewRoot+"/")
		for k, _ := range a.Views {
			if strings.HasPrefix(pathIgnoreRoot, k) {
				a.Views[k] = append(a.Views[k], path)
			}
		}
	}
	return nil
}

// Traverses and maps resource directory
func (a *AppDirectory) resWalker(path string, fi os.FileInfo, err error) error {
	if fi.IsDir() && path != a.ResourceRoot {
		a.Resources[fi.Name()] = make([]string, 0)
	} else {
		pathIgnoreRoot := strings.TrimPrefix(path, a.ResourceRoot+"/")
		for k, _ := range a.Resources {
			if strings.HasPrefix(pathIgnoreRoot, k) {
				a.Resources[k] = append(a.Resources[k], path)
			}
		}
	}
	return nil
}

// Attempts to locate a view path under controller directory key
func (a *AppDirectory) GetView(ctrl string, v string) string {
	for _, path := range a.Views[ctrl] {
		for _, ext := range a.Extensions {
			suf := v + ext
			if strings.HasSuffix(path, suf) {
				return path
			}
		}
	}
	return a.ViewRoot + "error/404.html"
}

// Maps out application directory
func (a *AppDirectory) Update() {
	a.Views = make(map[string][]string)
	a.Resources = make(map[string][]string)
	err := filepath.Walk(a.ViewRoot, a.viewWalker)
	err = filepath.Walk(a.ResourceRoot, a.resWalker)
	if err != nil {
		Logger.Error("%s\n", err.Error())
	}
}
