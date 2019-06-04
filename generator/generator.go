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
	Package  string
	TypeName string
	Name     string
}

func Generate(typeName, dir string) error {
	if typeName == "" || dir == "" {
		return errors.New("invalid arguments")
	}

	pkg, err := getPackage(dir)
	if err != nil {
		return err
	}

	d := templateData{
		Package:  pkg.Name,
		TypeName: typeName,
		Name:     getNameFromObjType(typeName),
	}

	path := getFilepath(dir, d.Name)

	b, err := generateFileBytes(path, d)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, b, 0644); err != nil {
		return err
	}

	return nil
}

func getPackage(dir string) (*packages.Package, error) {
	p, err := packages.Load(&packages.Config{
		Dir: dir,
	}, ".")
	if err != nil {
		return nil, errors.Wrap(err, "failed to load package info")
	}

	if len(p) != 1 {
		return nil, errors.Errorf("invalid directory given (%s)", dir)
	}

	return p[0], nil
}

func getNameFromObjType(s string) string {
	s = strings.TrimLeft(s, "*")
	return strings.ToUpper(s[:1]) + s[1:]
}

func getFilepath(dir, name string) string {
	fn := strings.ToLower(name) + "_slicegen.go"
	return filepath.Join(dir, fn)
}

func generateFileBytes(filename string, data templateData) ([]byte, error) {
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return nil, errors.Wrap(err, "failed to execute template")
	}

	src, err := imports.Process(filename, buf.Bytes(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to format file")
	}
	return src, nil
}
