package auto

import (
	"html/template"
	"path/filepath"

	"github.com/elazarl/go-bindata-assetfs"
)

// Load loads the templates from the embedded file map. This function will not
// compile if go generate is not executed before.
func LoadTemplates() *template.Template {
	dir, _ := AssetDir("templates")
	tmpl := template.New("_")
	for _, name := range dir {
		path := filepath.Join("templates", name)
		src := MustAsset(path)
		tmpl = template.Must(
			tmpl.New(name).Parse(string(src)),
		)
	}

	return tmpl
}

// AssetFS returns the AssetFS object
func AssetFS(dir string) *assetfs.AssetFS {
	return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: dir}
}
