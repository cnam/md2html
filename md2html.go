package main

import (
	"github.com/codegangsta/cli"
	"os"
	"github.com/cnam/md2html/generator"
)

const APP_VER  = "0.1"

func main() {
	app := cli.NewApp()
	app.Name = "md2html"
	app.Email = "support@leanlabs.io"
	app.Usage = "Github generator html pages from markdown wiki"
	app.Version = APP_VER
	app.Action = generator.GenerateDoc
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input, i",
			Usage: "Direcotory with markdown files",
		},
		cli.StringFlag{
			Name:  "output, o",
			Value: "documentation",
			Usage: "Directory for output files",
		},
		cli.StringFlag{
			Name:  "template, t",
			Value: "templates/documentation.tpl",
			Usage: "Template for generated documentation",
		},
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

