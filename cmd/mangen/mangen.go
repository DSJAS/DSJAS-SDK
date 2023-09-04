// Mangen generates a DSJAS manifest file for an extension with the given name
// and type. If no custom details are provided, details are guessed from the
// surroundings.
package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	dsjas "github.com/DSJAS/DSJAS-SDK"
)

var (
	Name   = flag.String("name", "", "The name under which the extension will install [default: current directory name]")
	Theme  = flag.Bool("theme", false, "Generate a manifest file for a theme")
	Module = flag.Bool("module", false, "Generate a manifest file module")
	Save   = flag.Bool("save", false, "Save the generated file to the current directory")
)

func main() {
	flag.Parse()
	if *Name == "" {
		var err error
		*Name, err = os.Getwd()
		if err != nil {
			fmt.Println("name autogeneration:", err)
			os.Exit(1)
		}

		*Name = path.Base(*Name)
	}

	var mode dsjas.PluginType
	if *Theme {
		mode = dsjas.PluginTheme
	} else if *Module {
		mode = dsjas.PluginModule
	} else {
		fmt.Println("mangen: invalid manifest configuration: need to generate for either theme or module")
		os.Exit(1)
	}

	m := dsjas.PluginManifest{Name: *Name, Type: mode}.String()
	if *Save {
		f, err := os.Create(dsjas.ManifestFile)
		if err != nil {
			fmt.Println("save manifest: ", err)
			os.Exit(1)
		}
		defer f.Close()

		f.WriteString(m)
	} else {
		fmt.Println(m)
	}
}
