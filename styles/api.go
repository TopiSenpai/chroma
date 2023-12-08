package styles

import (
	"embed"
	"io/fs"
	"path/filepath"
	"sort"

	"github.com/topi314/chroma/v2"
	"github.com/topi314/chroma/v2/base16"
)

//go:embed embedded
var embedded embed.FS

// Registry of Styles.
var Registry = func() map[string]*chroma.Style {
	registry := map[string]*chroma.Style{}
	// Register all embedded styles.
	embeddedSub, err := fs.Sub(embedded, "embedded")
	if err != nil {
		panic(err)
	}
	styles, err := LoadFromFS(embeddedSub)
	if err != nil {
		panic(err)
	}
	for _, style := range styles {
		registry[style.Name] = style
	}
	return registry
}()

// Fallback style. Reassign to change the default fallback style.
var Fallback = Registry["swapoff"]

func LoadFromFS(fsys fs.FS) ([]*chroma.Style, error) {
	files, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return nil, err
	}

	styles := make([]*chroma.Style, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		r, err := fsys.Open(file.Name())
		if err != nil {
			return nil, err
		}

		var style *chroma.Style
		switch filepath.Ext(file.Name()) {
		case ".xml":
			style, err = chroma.NewXMLStyle(r)
		case ".yaml", ".yml":
			style, err = base16.NewStyle(r)
		}
		if err != nil {
			return nil, err
		}

		styles = append(styles, style)
	}

	return styles, nil
}

// Register a chroma.Style.
func Register(style *chroma.Style) *chroma.Style {
	Registry[style.Name] = style
	return style
}

// Names of all available styles.
func Names() []string {
	out := []string{}
	for name := range Registry {
		out = append(out, name)
	}
	sort.Strings(out)
	return out
}

// Get named style, or Fallback.
func Get(name string) *chroma.Style {
	if style, ok := Registry[name]; ok {
		return style
	}
	return Fallback
}
