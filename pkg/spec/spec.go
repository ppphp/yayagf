package spec

import (
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"golang.org/x/xerrors"
	"io/ioutil"
)

func RecoverOrError(fs func()) (err error) {
	defer func() {
		x := recover()
		if x == nil {
			err = xerrors.Errorf("not panic")
		}
	}()
	fs()
	return
}

type Golden struct {
	filename string
	expect   []byte
	result   []byte
}

func (g *Golden) Compare(data []byte) bool {
	g.result = data
	if string(data) == string(g.expect) {
		return true
	}
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(string(data), string(g.expect), false)

	fmt.Println(string(data))
	fmt.Println(string(g.expect))
	fmt.Println(dmp.DiffToDelta(diffs))

	return false
}

func (g *Golden) Update() error {
	return ioutil.WriteFile(g.filename, g.result, 0644)
}

func LoadGolden(path string) *Golden {
	g := &Golden{filename: path}
	b, _ := ioutil.ReadFile(path)
	if b == nil {
		b = []byte{}
	}
	g.expect = b
	return g
}
