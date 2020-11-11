package command

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDoCommand(t *testing.T) {
	err, e, o := DoCommand("ls")
	require.NoErrorf(t, err, "ls should be done, %v,%v, %v", e, o, err)

	err, e, o = DoCommand("cnm")
	require.Errorf(t, err, "cnm should not be done, %v,%v, %v", e, o, err)
}

func TestGoCommand(t *testing.T) {

	out, errs := &bytes.Buffer{}, &bytes.Buffer{}
	cmd := GoCommand("cnm", nil, out, errs)
	require.NotNil(t, cmd, "cnm should not be done, %v,%v", out.String(), errs.String())
}
