package cli

import (
	"reflect"
	"testing"
)

func TestCommand_exec(t *testing.T) {
	type fields struct {
		Commands map[string]CommandFactory
		Args     []string
		Flags    map[string]string
		Run      func(args []string, flags map[string]string) (int, error)
	}
	type args struct {
		cargs []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Command{
				Commands: tt.fields.Commands,
				Args:     tt.fields.Args,
				Flags:    tt.fields.Flags,
				Run:      tt.fields.Run,
			}
			got, err := c.exec(tt.args.cargs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Command.exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Command.exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommand_parseArgs(t *testing.T) {
	type fields struct {
		Commands map[string]CommandFactory
		Args     []string
		Flags    map[string]string
		Run      func(args []string, flags map[string]string) (int, error)
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Command{
				Commands: tt.fields.Commands,
				Args:     tt.fields.Args,
				Flags:    tt.fields.Flags,
				Run:      tt.fields.Run,
			}
			c.parseArgs(tt.args.args)
		})
	}
}

func TestNewApp(t *testing.T) {
	type args struct {
		name    string
		version string
	}
	tests := []struct {
		name string
		args args
		want *App
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewApp(tt.args.name, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewApp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_Run(t *testing.T) {
	type fields struct {
		Name    string
		Version string
		Command *Command
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				Name:    tt.fields.Name,
				Version: tt.fields.Version,
				Command: tt.fields.Command,
			}
			got, err := a.Run()
			if (err != nil) != tt.wantErr {
				t.Errorf("App.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("App.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_RunArgs(t *testing.T) {
	type fields struct {
		Name    string
		Version string
		Command *Command
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				Name:    tt.fields.Name,
				Version: tt.fields.Version,
				Command: tt.fields.Command,
			}
			got, err := a.RunArgs(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("App.RunArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("App.RunArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
