package it

import (
	"gitlab.papegames.com/fengche/yayagf/pkg/spec"
	"golang.org/x/xerrors"
	"testing"
)

func TestShould(t *testing.T) {
	Should("not panic")
}

func TestSpec_With(t *testing.T) {
	s := Should("not be nil after with")
	s.With(t).
		Run(func() {
			if s.T == nil {
				t.Errorf("t is nil")
			}
		})
}

func TestSpec_When(t *testing.T) {
	i := 0
	Should("run when").
		With(t).
		When(func() (string, error) {
			i++
			return "", nil
		}).
		Run(func() {
			if i != 1 {
				t.Errorf("i not changed %v", i)
			}
		})
	Should("panic when").
		With(t).Run(
		func() {
			if err := spec.RecoverOrError(
				func() {
					Should("panic ").When(func() (string, error) {
						return "", xerrors.Errorf("panic")
					})
				}); err != nil {
				t.Errorf("not panic, wrong")
			}
		})
}

func TestSpec_Run(t *testing.T) {
	Should("run when").
		Run(func() {
			// running!
		})
}
