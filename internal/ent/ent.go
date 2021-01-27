package ent

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/facebook/ent/entc"
	"github.com/facebook/ent/entc/gen"
	"github.com/facebook/ent/schema/field"
	"golang.org/x/tools/go/packages"
)

// generate path and model names
func GenerateSchema(path string, names []string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
	}
	if err != nil {
		return err
	}

	for _, name := range names {
		b := bytes.NewBuffer(nil)
		if err := tmpl.Execute(b, name); err != nil {
			return err
		}
		target := filepath.Join(path, strings.ToLower(name+".go"))
		if err := ioutil.WriteFile(target, b.Bytes(), 0644); err != nil {
			return err
		}
	}
	return nil
}

var tmpl = template.Must(template.New("schema").
	Parse(`package schema

import "github.com/facebook/ent"

// {{ . }} holds the schema definition for the {{ . }} entity.
type {{ . }} struct {
	ent.Schema
}

// Fields of the {{ . }}.
func ({{ . }}) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the {{ . }}.
func ({{ . }}) Edges() []ent.Edge {
	return nil
}

// Indexes of the {{ . }}.
func ({{ . }}) Indexes() []ent.Index {
    return []ent.Index{}
}

`))

var DefaultConfig = &packages.Config{Mode: packages.NeedName}

// cmd/internal/base/base.go
func GenerateCRUDFiles(mod, path, target string, template []string) error {
	cfg := gen.Config{Target: target, IDType: &field.TypeInfo{Type: field.TypeInt}, Package: filepath.Join(mod, "app", "crud")}
	opts := []entc.Option{entc.Storage("sql")}
	for _, tmpl := range template {
		opts = append(opts, entc.TemplateDir(tmpl))
	}
	// If the target directory is not inferred from
	// the schema path, resolve its package path.
	pkgPath, err := PkgPath(DefaultConfig, cfg.Target)
	if err != nil {
		return err
	}
	cfg.Package = pkgPath
	if err := entc.Generate(path, &cfg, opts...); err != nil {
		return err
	}

	// generate a client
	err = ioutil.WriteFile(filepath.Join(target, "c.go"), []byte("package crud\nvar C *Client\n"), 0644)
	if err != nil {
		return err
	}
	return nil
}

func PkgPath(config *packages.Config, target string) (string, error) {
	if config == nil {
		config = DefaultConfig
	}
	pathCheck, err := filepath.Abs(target)
	if err != nil {
		return "", err
	}
	var parts []string
	if _, err := os.Stat(pathCheck); os.IsNotExist(err) {
		parts = append(parts, filepath.Base(pathCheck))
		pathCheck = filepath.Dir(pathCheck)
	}
	// Try maximum 2 directories above the given
	// target to find the root package or module.
	for i := 0; i < 2; i++ {
		pkgs, err := packages.Load(config, pathCheck)
		if err != nil {
			return "", fmt.Errorf("load package info: %v", err)
		}
		if len(pkgs) == 0 || len(pkgs[0].Errors) != 0 {
			parts = append(parts, filepath.Base(pathCheck))
			pathCheck = filepath.Dir(pathCheck)
			continue
		}
		pkgPath := pkgs[0].PkgPath
		for j := len(parts) - 1; j >= 0; j-- {
			pkgPath = path.Join(pkgPath, parts[j])
		}
		return pkgPath, nil
	}
	return "", fmt.Errorf("root package or module was not found for: %s", target)
}
