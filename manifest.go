package dsjas

import "encoding/json"

const (
	// File name for the DSJAS plugin manifest.
	ManifestFile = ".dsjasManifest"
)

// Plugin type definitions. Used for format extension manifest file types.
const (
	PluginTheme = iota
	PluginModule
	// Currently unused.
	PluginExtension
)

// PluginType holds the integer representation of a plugin type.
type PluginType uint8

// Returns the correctly stringified PluginType.
func (t PluginType) MarshalJSON() ([]byte, error) {
	return []byte("\"" + t.String() + "\""), nil
}

func (t PluginType) String() string {
	switch t {
	case PluginTheme:
		return "theme"
	case PluginModule:
		return "module"
	case PluginExtension:
		return "extension"
	default:
		return "unknown"
	}
}

// PluginManifest is a JSON representation of the .dsjasManifest file.
type PluginManifest struct {
	Name string     `json:"name"`
	Type PluginType `json:"extension-type"`
}

// Returns the stringified JSON representation of the manifest file.
func (p PluginManifest) String() string {
	out, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		panic(err)
	}

	return string(out)
}
