package generator

import (
	"github.com/shurcooL/github_flavored_markdown"
	"os"
	"io/ioutil"
	"strings"
	"html/template"
	"fmt"
	"errors"
)

//Page represent page for generate
type Page struct {
	Title    string
	Url      string
	Path     string
	Items    []*Page
	Body     template.HTML
	Sidebar  template.HTML
	Template string
}

// NewPage create new page
func (d *Dir) NewPage(f os.FileInfo) (*Page, error) {
	if !strings.HasSuffix(f.Name(), ".md") || strings.HasPrefix(f.Name(), "_") {
		return nil, errors.New(fmt.Sprintf("Not allowed file format %s\n", f.Name()))
	}

	markdown, err := ioutil.ReadFile(getPath(d.mdDir, f.Name()))

	if (err != nil) {
		return nil, err
	}

	title := strings.Replace(f.Name(), ".md", "", 1)

	html := string(github_flavored_markdown.Markdown(markdown))
	html = strings.Replace(html, "README.md", "index.html", -1)
	html = strings.Replace(html, ".md", ".html", -1)

	p := &Page{}
	p.Title = title
	p.Body = template.HTML(html)
	p.Path = getPath(d.htmlDir, getUrl(p.Title) + ".html")
	p.Url = getPath(d.path, getUrl(p.Title) + ".html")
	p.Template = d.template

	return p, nil
}

// getUrl returns generated url
func getUrl(title string) string {
	url := title
	if title == "README" {
		url = "index"
	}

	return url
}

// save saving current page to filesystem
func (p *Page) save(d *Dir) error {
	p.Sidebar = d.sidebar
	p.Items = d.pages
	file, err := os.Create(p.Path)

	if (err != nil) {
		return err
	}

	fmt.Printf("Create new page: %s\n \tby link:%s\n", p.Title, p.Path)

	return p.render(file)
}

// render rendering current page template
func (p *Page) render(f *os.File) error {
	t, err := template.ParseFiles(p.Template)

	if (err != nil) {
		return err
	}

	return t.Execute(f, p)
}