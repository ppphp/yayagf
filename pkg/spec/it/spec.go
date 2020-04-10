package it

import "testing"

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
		s.T.Fatalf("%v failed when %v because %v", s.Title, state, err)
	}
	s.Condition = append(s.Condition, )
	return s
}

func (s *Spec) Then(fs func() error ) *Spec {
	err := fs()
	if err != nil {
		s.T.Errorf("%v failed when testing because %v", s.Title, err)
	}
	return s
}
