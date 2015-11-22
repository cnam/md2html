package generator

import (
	"fmt"
	"os"
	"html/template"
)

// Dir represents file directory
type Dir struct {
	dir      []*Dir
	mdDir    string
	htmlDir  string
	pages    []*Page
	sidebar  template.HTML
	template string
	path     string
	static   []*StaticFile
}

// NewDir returns new dir
func NewDir(md, html, t, path string) *Dir {
	return &Dir{
		mdDir: md,
		htmlDir: html,
		template: t,
		path: path,
	}
}
// read reading all child directory and pages from dir
func (d *Dir) read() error {
	fmt.Printf("Read dir: %s\n", d.mdDir)
	osd, err := os.Open(d.mdDir);
	defer osd.Close()

	fi, err := osd.Readdir(-1)

	if err != nil {
		return err
	}

	for _, f := range fi {
		if f.Mode().IsDir() {
			dir := NewDir(getPath(d.mdDir, f.Name()),
				getPath(d.htmlDir, f.Name()),
				d.template,
				getPath(d.path, f.Name()),
			)
			dir.read()
			d.addDir(dir)
		}
		if f.Mode().IsRegular() {
			page, err := d.NewPage(f)
			if (err == nil) {
				d.addPage(page)
			} else {
				st, err := d.NewStatic(f)
				if (err == nil) {
					d.addStatic(st);
				}
			}
		}
	}

	return nil;
}
// write writes content to html directory
func (d *Dir) write(parent *Dir) error {
	err := os.MkdirAll(d.htmlDir, 0775)
	if err != nil {
		return err
	}
	sd, err := NewSidebar(d.mdDir)

	if err == nil {
		fmt.Printf("Create new sidebar \n\t%s\n", d.mdDir)
		d.sidebar = sd
	} else {
		fmt.Printf("Sidebar not found \n\t%s\n", d.mdDir)
		d.sidebar = parent.sidebar
	}

	for _, p := range d.pages {
		err := p.save(d)
		if err != nil {
			return err
		}
	}

	for _, dir := range d.dir {
		err := dir.write(d)
		if err != nil {
			return err
		}
	}

	for _, st := range d.static  {
		err := st.write(d)
		if err != nil {
			return err
		}
	}

	return nil
}

// addPage adding new page to current dir
func (d *Dir) addPage(p *Page) {
	d.pages = append(d.pages, p)
}

// addDir adding new child directory
func (d *Dir) addDir(dir *Dir) {
	d.dir = append(d.dir, dir)
}

// addStating adding new static file to direcotory
func (d *Dir) addStatic(s *StaticFile) {
	d.static = append(d.static, s)
}

//  getPath returns concat string for current dir path
func getPath(c, f string) string {
	return fmt.Sprintf("%s%s%s", c, string(os.PathSeparator), f)
}