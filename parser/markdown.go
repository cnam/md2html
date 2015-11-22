package parser

import (
	"strings"
	"github.com/shurcooL/github_flavored_markdown"
)

type MdParser struct {

}

func NewMdParser() *MdParser {
	return &MdParser{}
}

func (prs *MdParser) GetTitle(name string) string {
	return strings.Replace(strings.Replace(name, ".md", "", 1), "_", " ", -1)
}

func (prs *MdParser) Parse(d []byte) string {
	html := string(github_flavored_markdown.Markdown(d))
	html = strings.Replace(html, "README.md", "index.html", -1)
	html = strings.Replace(html, ".md", ".html", -1)

	return html
}