package generator

import (
    "github.com/shurcooL/github_flavored_markdown"
    "os"
    "io/ioutil"
    "log"
    "strings"
    "html/template"
    "html"
    "github.com/codegangsta/cli"
)

type Page struct {
    Title string
    Url string
    Body string
    Items []Page
    Sidebar string
    Template string
}

func GenerateDoc(c *cli.Context) {
    var pages []Page
    var page Page
    var sidebar string
    mdDir := c.String("i")
    htmlDir := c.String("o")
    template := c.String("t")
    MdSidebar := c.String("md") + "/_Sidebar.md"

    files, err := ioutil.ReadDir(mdDir)

    if (err != nil) {
        log.Panic(err);
    }

    for _, f := range files {
        if strings.HasSuffix(f.Name(), ".md") && ! strings.HasPrefix(f.Name(), "_") {
            markdown, err := ioutil.ReadFile(mdDir + "/"+f.Name())
            log.Printf("%s generated", f.Name());

            if (err != nil) {
                log.Panic(err);
            }

            title := strings.Replace(f.Name(), ".md", "", 1)

            page = Page{}
            page.Title = title
            page.Url = htmlDir + "/" + title + ".html"
            page.Body = html.UnescapeString(string(github_flavored_markdown.Markdown(markdown)))

            pages = append(pages, page)
        }
    }

    file, err := ioutil.ReadFile(MdSidebar);

    if err == nil {
        sidebar = html.UnescapeString(string(github_flavored_markdown.Markdown(file)))
    }

    for _, p := range pages {
        p.Template = template
        p.Items = pages
        p.Sidebar = sidebar
        p.save()
    }
}

func (p *Page) save() error {
    file, err := os.Create(p.Url)

    if (err != nil) {
        log.Panic(err);
    }

    return renderTemplate(file, p)
}

func renderTemplate(f *os.File, p *Page) error {
    t, err := template.ParseFiles(p.Template)

    if (err != nil) {
        log.Panic(err);
    }

    return t.Execute(f, p)
}
