package cmd

import (
    "github.com/codegangsta/cli"
    "github.com/cnam/md2html/generator"
)

var GenerateCmd = cli.Command{
    Name:  "daemon",
    Usage: "Generate documentation",
    Flags: []cli.Flag{
        cli.StringFlag{
            Name:  "i",
            Value: "wiki",
            Usage: "Direcotory with markdown files",
        },
        cli.StringFlag{
            Name:  "o",
            Value: "documentation",
            Usage: "Directory for output files",
        },
        cli.StringFlag{
            Name:  "t",
            Value: "documentation.tpl",
            Usage: "Template for generated documentation",
        },
    },
    Action: generator.GenerateDoc,
}

