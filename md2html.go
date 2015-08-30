package main

import (
	"github.com/codegangsta/cli"
	"os"
	"github.com/cnam/md2html/cmd"
)

const APP_VER  = "0.1"

func main() {
	app := cli.NewApp()
	app.Name = "md2html"
	app.Email = "support@leanlabs.io"
	app.Usage = "Github generator html pages from markdown wiki"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.GenerateCmd,
	}
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "V",
			Email: "support@leanlabs.io",
		},
		cli.Author{
			Name:  "cnam",
			Email: "cnam812@gmail.com",
		},
	}
	app.Run(os.Args)
}

