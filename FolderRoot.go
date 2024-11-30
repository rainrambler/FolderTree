package main

import (
	"fmt"
	"path/filepath"
	"strings"

	multimap "github.com/jwangsadinata/go-multimap"
	"github.com/jwangsadinata/go-multimap/slicemultimap"
)

const DesiredLayers = 2
const TmplFile = `FolderTree_templ.html`

func IterRead(dirname string) {
	var fr FolderRoot
	fr.dir2infos = slicemultimap.New()
	fr.Read(dirname, 0)

	//fr.PrintResult()
	js := fr.StartToJson()
	fmt.Println(js)
}

func ConvertChart(dirname, outfile string) {
	var fr FolderRoot
	fr.dir2infos = slicemultimap.New()
	fr.Read(dirname, 0)

	js := fr.StartToJson()

	content, err := ReadTextFile(TmplFile)
	if err != nil {
		fmt.Printf("WARN: Cannot read template file: %s!\n", TmplFile)
		return
	}

	content = strings.Replace(content, `$REALDATA$`, js, 1)
	err = WriteTextFile(outfile, content)
	if err != nil {
		fmt.Printf("WARN: Cannot write to file: %s!\n", outfile)
	} else {
		fmt.Printf("INFO: Saved to file: %s.\n", outfile)
	}
}

type FolderRoot struct {
	dir2infos multimap.MultiMap
	dirName   string
	partName  string
	filecount int
}

func (p *FolderRoot) Read(dirname string, curlayer int) bool {
	if curlayer > DesiredLayers {
		return false
	}
	p.dirName = dirname
	children := ListSubDirs(dirname)

	for _, child := range children {
		fr := new(FolderRoot)
		fr.dir2infos = slicemultimap.New()

		fullpath := filepath.Join(dirname, child)
		fr.filecount = FindFileCountInDir(fullpath)
		fr.dirName = fullpath
		fr.partName = child

		res := fr.Read(fullpath, curlayer+1)
		if res {
			p.dir2infos.Put(fullpath, fr)
		}
	}

	return true
}

func (p *FolderRoot) PrintResult() {
	fmt.Printf("%v | File count: %d\n", p.dirName, p.filecount)

	for _, e := range p.dir2infos.Entries() {
		e.Value.(*FolderRoot).PrintResult()
	}
}

func (p *FolderRoot) ToJson() string {
	if p.dir2infos.Size() == 0 {
		format := "{ %s }"
		return fmt.Sprintf(format, p.props2json())
	}

	s := "["
	for _, e := range p.dir2infos.Entries() {
		s += e.Value.(*FolderRoot).ToJson() + ","
	}

	s = RemoveLastChar(s)
	s += "]"

	format := `{%s, "children": %s }`
	return fmt.Sprintf(format, p.props2json(), s)
}

func (p *FolderRoot) StartToJson() string {
	if p.dir2infos.Size() == 0 {
		return ""
	}

	s := "["
	for _, e := range p.dir2infos.Entries() {
		s += e.Value.(*FolderRoot).ToJson() + ","
	}

	s = RemoveLastChar(s)
	s += "]"
	return s
}

func (p *FolderRoot) props2json() string {
	s := `"value": %d, "name": "%s", "path": "%s"`
	return fmt.Sprintf(s, p.filecount, p.partName, p.dirName)
}

func RemoveLastChar(s string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return ""
	default:
		return s[:len(s)-1]
	}
}
