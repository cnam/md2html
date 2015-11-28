package generator

import (
	"html/template"
	"io/ioutil"
	"github.com/cnam/md2html/parser"
)

// NewSidebar created new sidebar
func NewSidebar(dir string) (template.HTML, error) {
	s, err := generateSidebar(dir);

	if err != nil {
		return "", err
	}

	return s, nil
}

func generateSidebar(mdSidebar string) (template.HTML, error) {
	var sidebar template.HTML
	prs,err := parser.New(mdSidebar)
	if err != nil {
		return "", err
	}

	file, err := ioutil.ReadFile(mdSidebar);

	if err != nil {
		return "", err
	}

	sidebar = template.HTML(prs.Parse(file))

	return sidebar, nil
}
