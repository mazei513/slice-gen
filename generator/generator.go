package generator

import (
	"bytes"
	"fmt"
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
	data, err := getData(objType, wd)
	if err != nil {
		return err
	}

	filename := strings.ToLower(data.Name) + "_slicegen.go"

	if err := writeTemplate(filepath.Join(wd, filename), data); err != nil {
		return err
	}

	return nil
}

func getData(objType, wd string) (templateData, error) {
	genPkg := getPackage(wd)
	if genPkg == nil {
		return templateData{}, fmt.Errorf("unable to find package info for " + wd)
	}

	return templateData{
		Package: genPkg.Name,
		ObjType: objType,
		Name:    uppercaseFirst(stripPointer(objType)),
	}, nil
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

func writeTemplate(filepath string, data templateData) error {
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return errors.Wrap(err, "generating code")
	}

	src, err := imports.Process(filepath, buf.Bytes(), nil)
	if err != nil {
		return errors.Wrap(err, "unable to gofmt")
	}

	if err := ioutil.WriteFile(filepath, src, 0644); err != nil {
		return errors.Wrap(err, "writing output")
	}

	return nil
}

func stripPointer(s string) string {
	return strings.TrimLeft(s, "*")
}

func uppercaseFirst(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}
