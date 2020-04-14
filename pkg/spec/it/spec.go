package it

import (
	"fmt"
	"testing"
)

type Spec struct {
	Title     string
	Condition []string
	T         *testing.T
}

func Should(title string) *Spec {
	s := &Spec{Title: title}
	return s
}

func (s *Spec) With(t *testing.T) *Spec {
	s.T = t
	return s
}

func (s *Spec) When(initF func() (string, error)) *Spec {
	state, err := initF()
	if err != nil {
		panic(fmt.Sprintf("%v failed when %v because %v", s.Title, state, err))
	}
	s.Condition = append(s.Condition, state)
	return s
}

func (s *Spec) Run(fs func()) *Spec {
	fs()
	return s
}

