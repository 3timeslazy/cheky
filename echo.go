package cheky

// EchoContext describes the methods required
// for extracting parameters from the
// github.com/labstack/echo.Context
type EchoContext interface {
	QueryParam(name string) (value string)
	Param(name string) (value string)
}

type echo struct{ EchoContext }

func (e echo) Query(name string) string {
	return e.QueryParam(name)
}

func (e echo) Path(name string) string {
	return e.Param(name)
}

// Echo creates Checker which extracts
// parameters from the context of labstack's Context.
func Echo(ctx EchoContext) Checker {
	return Ctx(echo{ctx})
}
