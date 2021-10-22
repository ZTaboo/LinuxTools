package main

import (
	"LinuxTools/cmd"
	"os"
)

func main() {
	app := cmd.Cmd()
	_ = app.Run(os.Args)
}
