package model

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// 有什么意义吗，有的，import少写两行。。。
func TestOpen(t *testing.T) {
	_, err := Open("postgres", "localhost")
	require.NoError(t, err)
}
