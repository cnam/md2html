package generator

import (
	"html/template"
	"io/ioutil"
	"github.com/shurcooL/github_flavored_markdown"
)

func NewSidebar(dir string) (template.HTML, error) {
	mdSidebar := getPath(dir, "_Sidebar.md")
	s, err := generateSidebar(mdSidebar);

	if err != nil {
		return "", err
	}

	return s, nil
}

func generateSidebar(mdSidebar string) (template.HTML, error) {
	var sidebar template.HTML
	file, err := ioutil.ReadFile(mdSidebar);

	if err != nil {
		return "", err
	}
	sidebar = template.HTML(github_flavored_markdown.Markdown(file))

	return sidebar, nil
}
