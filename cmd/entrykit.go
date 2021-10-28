package main

import (
	"fmt"
	"os"

	"github.com/naoki85/entrykit"

	_ "github.com/naoki85/entrykit/codep"
	_ "github.com/naoki85/entrykit/prehook"
	_ "github.com/naoki85/entrykit/render"
	_ "github.com/naoki85/entrykit/switch"
)

var Version string

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-v", "--version":
			fmt.Println(Version)
			os.Exit(0)
		case "--symlink":
			entrykit.Symlink()
			os.Exit(0)
		}
	}
	entrykit.RunCmd()
}
