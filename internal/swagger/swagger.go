// just rewrite, a lot of copy paste
package swagger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/ghodss/yaml"
	"github.com/swaggo/swag"
	"gitlab.papegames.com/fengche/yayagf/internal/file"
)

func GenerateSwagger() error {
	root, err := file.GetAppRoot()
	if err != nil {
		return err
	}
	if err := New().Build(&Config{
		SearchDir:          root,
		MainAPIFile:        "main.go",
		OutputDir:          filepath.Join(root, "app", "doc"),
	}); err != nil {
		return err
	}
	return nil
}

type Gen struct {
	jsonIndent func(data interface{}) ([]byte, error)
	jsonToYAML func(data []byte) ([]byte, error)
}

// New creates a new Gen.
func New() *Gen {
	return &Gen{
		jsonIndent: func(data interface{}) ([]byte, error) {
			return json.MarshalIndent(data, "", "    ")
		},
		jsonToYAML: yaml.JSONToYAML,
	}
}

// Config presents Gen configurations.
type Config struct {
	// SearchDir the swag would be parse
	SearchDir string

	// OutputDir represents the output directory for all the generated files
	OutputDir string

	// MainAPIFile the Go file path in which 'swagger general API Info' is written
	MainAPIFile string
}

// Build builds swagger json file  for given searchDir and mainAPIFile. Returns json
func (g *Gen) Build(config *Config) error {
	if _, err := os.Stat(config.SearchDir); os.IsNotExist(err) {
		return fmt.Errorf("dir: %s is not exist", config.SearchDir)
	}

	p := swag.New()

	if err := p.ParseAPI(config.SearchDir, config.MainAPIFile); err != nil {
		return err
	}
	swagger := p.GetSwagger()

	b, err := g.jsonIndent(swagger)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(config.OutputDir, os.ModePerm); err != nil {
		return err
	}

	docFileName := path.Join(config.OutputDir, "docs.go")
	if err := ioutil.WriteFile(docFileName, []byte(fmt.Sprintf(`package doc
const Swagger = %s
`, strconv.Quote(string(b)))), 0644); err != nil {
		return err
	}
	log.Printf("create docs.go at  %+v", docFileName)
	return nil
}
