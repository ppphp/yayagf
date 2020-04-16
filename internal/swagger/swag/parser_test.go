package swag

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gitlab.papegames.com/fengche/yayagf/pkg/spec"
	"gitlab.papegames.com/fengche/yayagf/pkg/spec/it"
	"os"
	"path/filepath"
	"testing"
)

func TestNew(t *testing.T) {
	it.Should("new").Run(func() {
		p := New()
		if p == nil {
			t.Errorf("new nil")
		}
	})
	it.Should("option").Run(func() {
		p := New(SetLogger(os.Stdout))
		if p == nil {
			t.Errorf("new nil")
		}
	})
}

func TestParser_ParseGeneralApiInfo(t *testing.T) {
	it.Should("parse general").Run(func() {
		g := spec.LoadGolden("./testdata/ParseGeneralApiInfoBase.golden")
		p := New()
		err := p.ParseGeneralAPIInfo("./testdata/ParseGeneralApiInfoBase.go")
		assert.NoError(t, err)

		b, _ := json.MarshalIndent(p.swagger, "", "    ")
		assert.True(t, g.Compare(b))
	}).Run(func() {
		g := spec.LoadGolden("./testdata/ParseGeneralApiInfoInfo.golden")
		p := New()
		err := p.ParseGeneralAPIInfo("./testdata/ParseGeneralApiInfoInfo.go")
		assert.NoError(t, err)

		b, _ := json.MarshalIndent(p.swagger, "", "    ")
		assert.True(t, g.Compare(b))
	}) /* TODO: finish it
	.Run(func() {
		g:=spec.LoadGolden("./testdata/ParseGeneralApiInfoAll.golden")
		p := New()
		err := p.ParseGeneralAPIInfo("./testdata/ParseGeneralApiInfoAll.go")
		assert.NoError(t, err)

		b, _ := json.MarshalIndent(p.swagger, "", "    ")
		assert.True(t, g.Compare(b))
	})*/

	it.Should("cover no main").Run(func() {
		p := New()
		assert.Error(t, p.ParseGeneralAPIInfo("./testdata/nothing.go"))
	})

	it.Should("cover fail x-").Run(func() {
		p := New()
		err := p.ParseGeneralAPIInfo("./testdata/ParseGeneralApiInfoXFail.go")
		assert.Error(t, err)
	})

}

func TestParser_visit(t *testing.T) {
	it.Should("visit a file").Run(func() {
		searchDir := filepath.Join("testdata", "pet", "main.go")
		p := New()
		f, err := os.Open(searchDir)
		assert.NoError(t, err)
		i, err := f.Stat()
		assert.NoError(t, err)
		err = p.visit(searchDir, i, nil)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(p.files))
	})
	it.Should("skip").Run(func() {
		searchDir := "testdata/pet"
		p := New()
		err := p.visit(searchDir, nil, nil)
		assert.Equal(t, err, filepath.SkipDir)
		assert.Equal(t, 0, len(p.files))
	})
	it.Should("visit error").Run(func() {
		searchDir := "testdata/malform.go"
		p := New()
		err := p.visit(searchDir, nil, nil)
		assert.Error(t, err)
	})
}

// from old file
func TestParser_GetAllGoFileInfo(t *testing.T) {
	it.Should("parse 2 files").Run(func() {
		searchDir := "testdata/pet"

		p := New()
		err := p.getAllGoFileInfo(searchDir)

		assert.NoError(t, err)
		assert.NotEmpty(t, p.files[filepath.Join("testdata", "pet", "main.go")])
		assert.NotEmpty(t, p.files[filepath.Join("testdata", "pet", "web", "handler.go")])
		assert.Equal(t, 2, len(p.files))
	})
}
