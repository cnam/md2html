package generator

import (
	"os"
	"io/ioutil"
	"strings"
	"html/template"
	"fmt"
	"errors"
	"github.com/cnam/md2html/parser"
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
	Seo      *Seo
}

type Seo struct {
	Title       string
	Description string
	Keywords    string
}

// NewPage create new page
func (d *Dir) NewPage(f os.FileInfo) (*Page, error) {
	prs, err := parser.New(f.Name())
	if err != nil || strings.HasPrefix(f.Name(), "_") {
		return nil, errors.New(fmt.Sprintf("Not allowed file format %s\n", f.Name()))
	}

	cont, err := ioutil.ReadFile(getPath(d.mdDir, f.Name()))

	if (err != nil) {
		return nil, err
	}

	title := prs.GetTitle(f.Name())
	html := prs.Parse(cont)

	p := &Page{}
	p.Title = title
	p.Seo = &Seo{
		Title: "",
		Description: "",
		Keywords: "",
	}
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