package cheky_test

import (
	"reflect"

	"github.com/3timeslazy/cheky"
)

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

func runMethods(obj interface{}, methods map[string][]interface{}) {
	currVal := reflect.ValueOf(obj)

	for method, args := range methods {
		in := make([]reflect.Value, len(args))

		for i, arg := range args {
			in[i] = reflect.ValueOf(arg)
		}

		currVal = currVal.MethodByName(method).Call(in)[0]
	}
}

func args(args ...interface{}) []interface{} {
	return args
}
