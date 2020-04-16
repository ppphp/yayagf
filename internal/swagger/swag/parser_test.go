package swag

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gitlab.papegames.com/fengche/yayagf/pkg/spec"
	"gitlab.papegames.com/fengche/yayagf/pkg/spec/it"
	"os"
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
		g:=spec.LoadGolden("./testdata/ParseGeneralApiInfoBase.golden")
		p := New()
		err := p.ParseGeneralAPIInfo("./testdata/ParseGeneralApiInfoBase.go")
		assert.NoError(t, err)

		b, _ := json.MarshalIndent(p.swagger, "", "    ")
		assert.True(t, g.Compare(b))
	}).Run(func() {
		g:=spec.LoadGolden("./testdata/ParseGeneralApiInfoInfo.golden")
		p := New()
		err := p.ParseGeneralAPIInfo("./testdata/ParseGeneralApiInfoInfo.go")
		assert.NoError(t, err)

		b, _ := json.MarshalIndent(p.swagger, "", "    ")
		assert.True(t, g.Compare(b))
	})/* TODO: finish it
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
