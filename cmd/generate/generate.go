package generate

import (
	"github.com/ppphp/yayagf/cmd/generate/ci"
	"github.com/ppphp/yayagf/cmd/generate/controller"
	"github.com/ppphp/yayagf/cmd/generate/curd"
	"github.com/ppphp/yayagf/cmd/generate/doc"
	"github.com/ppphp/yayagf/cmd/generate/schema"
	"github.com/ppphp/yayagf/pkg/cli"
)

func CommandFactory() (*cli.Command, error) {
	c := &cli.Command{
		Commands: map[string]cli.CommandFactory{
			"schema": schema.CommandFactory, "s": schema.CommandFactory,
			"model": schema.CommandFactory,
			"crud":  curd.CommandFactory, "curd": curd.CommandFactory,
			"doc": doc.CommandFactory, "docs": doc.CommandFactory,
			"ci":         ci.CommandFactory,
			"controller": controller.CommandFactory, "c": controller.CommandFactory,
		},
	}
	return c, nil
}
