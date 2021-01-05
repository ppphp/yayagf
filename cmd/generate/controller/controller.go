package controller

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"path/filepath"

	"gitlab.papegames.com/fengche/yayagf/pkg/cli"
)

func CommandFactory() (*cli.Command, error) {
	c := &cli.Command{
		Run: func(args []string, flags map[string]string) (int, error) {
			root, err := file.GetAppRoot()
			if err != nil {
				log.Printf("get project name failed: %v", err.Error())
				return 1, err
			}
			tmpl, err := template.New("controller").Parse(jenkinsTemplate)
			if err != nil {
				log.Printf("jenkinsTemplate parse failed: %v", err.Error())
				return 1, err
			}
			b := bytes.Buffer{}
			if err := tmpl.Execute(&b, file.GetProjectInfo()); err != nil {
				log.Printf("jenkinsTemplate parse failed: %v", err.Error())
				return 1, err
			}
			if err := ioutil.WriteFile(filepath.Join(root, "Jenkinsfile"), b.Bytes(), 0644); err != nil {
				log.Printf("write file failed: %v", err.Error())
				return 1, err
			}
			return 0, nil
		},
	}
	return c, nil
}

const controllerTemplate = `

`
