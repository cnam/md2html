# Markdown to html page generator

For create beautiful documentation from your wiki or your documentation in repository.
Your may be use multiply repositories wiki for create one portal with documentation.

We may be use beautiful templates for generate documentation or added your custom template.

You do not need anything except to git, because md2html created in golang and compile to binary

### Installation

We may be use installation from binary or with docker

#### Installation from binary

[Download release](https://github.com/cnam/md2html/releases)

##### Mac os

```bash
wget https://github.com/cnam/md2html/releases/download/0.2.0/darwin_md2html

mv darwin_md2html /usr/local/bin/md2html

chmod +x /usr/local/bin/md2html
```

##### Linux

```bash
wget https://github.com/cnam/md2html/releases/download/0.2.0/linux_md2html

mv linux_md2html /usr/local/bin/md2html

chmod +x /usr/local/bin/md2html
```

#### Usage with docker

We assume that you have already installed [Docker](https://www.docker.com/)

```bash
docker pull cnam/md2html:latest

docker run -v wiki.dir:wiki.dir \
           -v documentation:documentation \
           -v template_dir:template_dir \
           cnam/md2html:latest -i wiki.dir -o documentation -t template_dir/template.tpl -p /docs
```

### Generate documentation

md2html -i wiki.dir -o documentation -t template.tpl -p /docs

**WHERE:**

- **-i or --input** Directory with markdown files
- **-o or --output** Directory for output generated html files
- **-t or --template** Template for generated documentation
- **-p or --path** Subdirectory in site eg /docs where documentation location
- **-s or --sidebar** Path to your custom sidebar

### Templates

Your must be create html template with variables [example](https://github.com/cnam/md2html/blob/master/templates/documentation.tpl)

**Default Variables**

- **{{ .Title }}** Page title similarly filename
- **{{ .Sidebar }}** Custom sidebar for view menu, html represents in file _Sidebar.md
- **{{ .Items }}** List pages for generate sidebar automatically. it contains
    - **{{ .Title }}** Name page from generated
    - **{{ .Url}}**    url to the pages
- **{{ .Body}}** Generated html body from markdown

### Sidebar as list pages

```html
<ul>
    <li><a href="/docs">Home</a></li>
    {{range .Items}}
        <li><a href="/{{ .Url}}">{{ .Title}}</a></li>
    {{end}}
</ul>
```

### Sidebar as static sidebar

```html
<div id="sidebar">
    {{ .Sidebar }}
</div>
```

### Mixed sidebar as html page and static

```html
{{if.Sidebar}}
    {{ .Sidebar}}
{{else}}
    <ul>
        <li><a href="/docs">Home</a></li>
        {{range .Items}}
            <li><a href="/{{ .Url}}">{{ .Title}}</a></li>
            {{end}}
        {{end}}
    </ul>
{{end}}
```

### Create sidebar as static file

In order to make your sidebar in each sub-directory, he can be your own or a common,
if it is set within a directory in the root of the documentation.

Sidebar represents as **_Sidebar.md** file and allow markdown syntax
