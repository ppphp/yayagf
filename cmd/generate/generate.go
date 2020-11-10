package generate

import (
	"gitlab.papegames.com/fengche/yayagf/cmd/generate/ci"
	"gitlab.papegames.com/fengche/yayagf/cmd/generate/curd"
	"gitlab.papegames.com/fengche/yayagf/cmd/generate/doc"
	"gitlab.papegames.com/fengche/yayagf/cmd/generate/job"
	"gitlab.papegames.com/fengche/yayagf/cmd/generate/schema"
	"gitlab.papegames.com/fengche/yayagf/pkg/cli"
)

func CommandFactory() (*cli.Command, error) {
	c := &cli.Command{
		Commands: map[string]cli.CommandFactory{
			"schema": schema.CommandFactory, "s": schema.CommandFactory,
			"model": schema.CommandFactory,
			"crud":  curd.CommandFactory, "curd": curd.CommandFactory,
			"doc": doc.CommandFactory, "docs": doc.CommandFactory,
			"ci":  ci.CommandFactory,
			"job": job.CommandFactory,
		},
	}
	return c, nil
}
