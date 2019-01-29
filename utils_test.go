package cheky_test

import "github.com/3timeslazy/cheky"

func newCheck(param string) *cheky.Checker {
	return cheky.Ctx(NewContext(param, "any"))
}

type CustomContext struct{ P, Q string }

func (cctx *CustomContext) Query(name string) string {
	return cctx.Q
}

func (cctx *CustomContext) Path(name string) string {
	return cctx.P
}

func NewContext(q, p string) cheky.Context {
	return &CustomContext{P: p, Q: q}
}
