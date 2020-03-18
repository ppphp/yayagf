package ent

import (
	"bytes"
	"fmt"
	"github.com/facebookincubator/ent/entc"
	"github.com/facebookincubator/ent/entc/gen"
	"github.com/facebookincubator/ent/schema/field"
	"golang.org/x/tools/go/packages"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
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

import "github.com/facebookincubator/ent"

// {{ . }} holds the schema definition for the {{ . }} entity.
type {{ . }} struct {
	ent.Schema
}

// Fields of the {{ . }}.
func ({{ . }}) Fields() []ent.Field {
	return nil
}

// Edges of the {{ . }}.
func ({{ . }}) Edges() []ent.Edge {
	return nil
}
`))

var DefaultConfig = &packages.Config{Mode: packages.NeedName}

func GenerateCRUDFiles(path string) error {
	type idType field.Type
	var (
		storage  string
		cfg      gen.Config
		template []string
		idtype   = idType(field.TypeInt)
	)

	opts := []entc.Option{entc.Storage(storage)}
	for _, tmpl := range template {
		opts = append(opts, entc.TemplateDir(tmpl))
	}
	// If the target directory is not inferred from
	// the schema path, resolve its package path.
	if cfg.Target != "" {
		pkgPath, err := PkgPath(DefaultConfig, cfg.Target)
		if err != nil {
			return err
		}
		cfg.Package = pkgPath
	}
	cfg.IDType = &field.TypeInfo{Type: field.Type(idtype)}
	err := entc.Generate(path, &cfg, opts...)
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