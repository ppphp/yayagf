package controller

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ppphp/yayagf/internal/file"

	"github.com/ppphp/yayagf/pkg/cli"
)

func CommandFactory() (*cli.Command, error) {
	c := &cli.Command{
		Run: func(args []string, flags map[string]string) (int, error) {
			root, err := file.GetAppRoot()
			if err != nil {
				log.Printf("get project name failed: %v", err.Error())
				return 1, err
			}
			tmpl, err := template.New("controller").Parse(controllerTemplate)
			if err != nil {
				log.Printf("controller parse failed: %v", err.Error())
				return 1, err
			}
			_ = os.MkdirAll(filepath.Join(root, "app", "controller"), 0755)
			for _, a := range args {
				b := bytes.Buffer{}
				if err := tmpl.Execute(&b, struct {
					Lower, Capital string
				}{strings.ToLower(a), strings.Title(strings.ToLower(a))}); err != nil {
					log.Printf("jenkinsTemplate parse failed: %v", err.Error())
					return 1, err
				}
				if err := ioutil.WriteFile(filepath.Join(root, "app", "controller", strings.ToLower(a)+".go"), b.Bytes(), 0644); err != nil {
					log.Printf("write file failed: %v", err.Error())
					return 1, err
				}
			}
			return 0, nil
		},
	}
	return c, nil
}

const controllerTemplate = `package controller 

// Index{{.Capital}} godoc
// @Summary {{.Capital}}
// @Tags {{.Lower}}
// @Accept json
// @Produce json
// @Success 200 {int} int 0
// @Failure 200 {int} int 0
// @Failure 200 {int} int 0
// @Router /{{.Lower}} [get]
func Index{{.Capital}}(c *maotai.Context) (int, string, gin.H) {
	return 0, "", nil
}

// Create{{.Capital}} godoc
// @Summary Create{{.Capital}}
// @Tags {{.Lower}}
// @Accept json
// @Produce json
// @Success 200 {int} int 0
// @Failure 200 {int} int 0
// @Failure 200 {int} int 0
// @Router /{{.Lower}} [post]
func Create{{.Capital}}(c *maotai.Context) (int, string, gin.H) {
	return 0, "", nil
}
`
