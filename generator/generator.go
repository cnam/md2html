package generator

import (
    "github.com/shurcooL/github_flavored_markdown"
    "os"
    "io/ioutil"
    "log"
    "strings"
    "html/template"
    "github.com/codegangsta/cli"
)

type Page struct {
    Title string
    Url string
    Body template.HTML
    Items []Page
    Sidebar template.HTML
    Template string
}

func GenerateDoc(c *cli.Context) {
    var pages []Page
    var page Page
    var sidebar template.HTML
    mdDir := c.String("input")
    htmlDir := c.String("output")
    t := c.String("template")

    if mdDir == "" {
        cli.ShowAppHelp(c)
        return
    }

    MdSidebar := c.String("input") + "/_Sidebar.md"

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
            page.Body = template.HTML(github_flavored_markdown.Markdown(markdown))

            pages = append(pages, page)
        }
    }

    file, err := ioutil.ReadFile(MdSidebar);

    if err == nil {
        sidebar = template.HTML(github_flavored_markdown.Markdown(file))
    }

    os.MkdirAll(htmlDir, 0775)

    for _, p := range pages {
        p.Template = t
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

    return p.renderTemplate(file)
}

func (p *Page) renderTemplate(f *os.File) error {
    t, err := template.ParseFiles(p.Template)

    if (err != nil) {
        log.Panic(err);
    }

    return t.Execute(f, p)
}
