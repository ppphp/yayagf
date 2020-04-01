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
		PropNamingStrategy: "camelcase",
		OutputDir:          filepath.Join(root, "app", "doc"),
		ParseVendor:        false,
		ParseDependency:    false,
		MarkdownFilesDir:   "",
		GeneratedTime:      true,
		Format:             "go",
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

	// PropNamingStrategy represents property naming strategy like snakecase,camelcase,pascalcase
	PropNamingStrategy string

	// ParseVendor whether swag should be parse vendor folder
	ParseVendor bool

	// ParseDependencies whether swag should be parse outside dependency folder
	ParseDependency bool

	// MarkdownFilesDir used to find markdownfiles, which can be used for tag descriptions
	MarkdownFilesDir string

	// GeneratedTime whether swag should generate the timestamp at the top of docs.go
	GeneratedTime bool

	Format string
}

// Build builds swagger json file  for given searchDir and mainAPIFile. Returns json
func (g *Gen) Build(config *Config) error {
	if _, err := os.Stat(config.SearchDir); os.IsNotExist(err) {
		return fmt.Errorf("dir: %s is not exist", config.SearchDir)
	}

	log.Println("Generate swagger docs....")
	p := swag.New(swag.SetMarkdownFileDirectory(config.MarkdownFilesDir))
	p.PropNamingStrategy = config.PropNamingStrategy
	p.ParseVendor = config.ParseVendor
	p.ParseDependency = config.ParseDependency

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

	switch config.Format {
	case "json":
		jsonFileName := path.Join(config.OutputDir, "swagger.json")
		err = g.writeFile(b, jsonFileName)
		if err != nil {
			return err
		}
		log.Printf("create swagger.json at  %+v", jsonFileName)
	case "yaml":
		yamlFileName := path.Join(config.OutputDir, "swagger.yaml")
		y, err := g.jsonToYAML(b)
		if err != nil {
			return fmt.Errorf("cannot covert json to yaml error: %s", err)
		}
		err = g.writeFile(y, yamlFileName)
		if err != nil {
			return err
		}
		log.Printf("create swagger.yaml at  %+v", yamlFileName)
	case "go":

		docFileName := path.Join(config.OutputDir, "docs.go")
		err := ioutil.WriteFile(docFileName, []byte(fmt.Sprintf(`package doc
const Swagger = %s
`, strconv.Quote(string(b)))), 0644)
		if err != nil {
			return err
		}
		log.Printf("create docs.go at  %+v", docFileName)
	}
	return nil
}

func (g *Gen) writeFile(b []byte, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(b)
	return err
}
