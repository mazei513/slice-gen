package generator

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/imports"
)

type templateData struct {
	Package string
	ObjType string
	Name    string
}

func Generate(objType, wd string) error {
	if objType == "" || wd == "" {
		return errors.New("invalid arguments")
	}

	pkg := getPackage(wd)
	if pkg == nil {
		return errors.Errorf("unable to find package info for " + wd)
	}

	d := templateData{
		Package: pkg.Name,
		ObjType: objType,
		Name:    getNameFromObjType(objType),
	}

	path := getFilepath(wd, d.Name)

	b, err := generateFileBytes(path, d)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, b, 0644); err != nil {
		return err
	}

	return nil
}

func getPackage(dir string) *packages.Package {
	p, _ := packages.Load(&packages.Config{
		Dir: dir,
	}, ".")

	if len(p) != 1 {
		return nil
	}

	return p[0]
}

func getNameFromObjType(s string) string {
	s = strings.TrimLeft(s, "*")
	return strings.ToUpper(s[:1]) + s[1:]
}

func getFilepath(wd, name string) string {
	fn := strings.ToLower(name) + "_slicegen.go"
	return filepath.Join(wd, fn)
}

func generateFileBytes(filepath string, data templateData) ([]byte, error) {
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return nil, errors.Wrap(err, "generating code")
	}

	src, err := imports.Process(filepath, buf.Bytes(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to gofmt")
	}
	return src, nil
}
