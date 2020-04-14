package spec

import "golang.org/x/xerrors"

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
