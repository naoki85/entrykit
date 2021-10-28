package main

import (
	"github.com/naoki85/entrykit"
	_ "github.com/naoki85/entrykit/codep"
)

var cmd = "codep"

func main() {
	entrykit.Cmds[cmd](
		entrykit.NewConfig(cmd, true))
}
