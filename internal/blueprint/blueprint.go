package blueprint

import "html/template"

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
