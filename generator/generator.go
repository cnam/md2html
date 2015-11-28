package generator

import (
	"github.com/codegangsta/cli"
	"fmt"
)

// GenerateDoc generating new documentation
func GenerateDoc(c *cli.Context) {
	md := c.String("input")
	html := c.String("output")
	t := c.String("template")
	path := c.String("path")

	if md == "" {
		cli.ShowAppHelp(c)
		return
	}

	fmt.Println("Begin generate")

	parent := &Dir{sidebar:""};
	dir, err := NewDir(md, html, t, path)
	if (err != nil) {
		fmt.Printf("Error read dir %s\n \t%s\n", dir.mdDir, err.Error())
	}
	err = dir.read()

	if err != nil {
	  fmt.Printf("Error read dir %s\n \t%s\n", dir.mdDir, err.Error())
	}
	err = dir.write(parent)

	if err != nil {
		fmt.Printf("Error write dir %s\n", dir.htmlDir)
	}

	fmt.Println("End generate")
}

