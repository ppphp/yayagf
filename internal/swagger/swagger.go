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

	"gitlab.papegames.com/fengche/yayagf/internal/file"
	"gitlab.papegames.com/fengche/yayagf/internal/swagger/swag"
)

func GenerateSwagger() error {
	root, err := file.GetAppRoot()
	if err != nil {
		return err
	}

	if _, err := os.Stat(root); os.IsNotExist(err) {
		return fmt.Errorf("dir: %s is not exist", root)
	}

	p := swag.New()

	if err := p.ParseAPI(root, "main.go"); err != nil {
		return err
	}
	swagger := p.GetSwagger()

	b, err := json.MarshalIndent(swagger, "", "    ")
	if err != nil {
		return err
	}

	OutputDir := filepath.Join(root, "app", "doc")
	if err := os.MkdirAll(OutputDir, os.ModePerm); err != nil {
		return err
	}

	docFileName := path.Join(OutputDir, "docs.go")
	if err := ioutil.WriteFile(docFileName, []byte(fmt.Sprintf(`package doc
const Swagger = %s
`, strconv.Quote(string(b)))), 0644); err != nil {
		return err
	}
	log.Printf("create docs.go at  %+v", docFileName)
	return nil

}
