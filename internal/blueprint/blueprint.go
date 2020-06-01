package blueprint

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

var blueprints = []Blueprint{}

func AddBlueprint(name, template string) {
	b := Blueprint{
		Name: name,
		//Tpl:  template.New(),
	}
	blueprints = append(blueprints, b)
}

type Blueprint struct {
	Name string
	Tpl  *template.Template
}

func GenerateBlueprint() {
}

func WriteFileWithTmpl(path string, tmpl string, params interface{}) error {
	tmp, err := template.New("").Parse(tmpl)
	if err != nil {
		return err
	}
	bts := &bytes.Buffer{}
	if err := tmp.Execute(bts, params); err != nil {
		return err
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	dir := filepath.Dir(abs)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	if err := ioutil.WriteFile(abs, bts.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}
