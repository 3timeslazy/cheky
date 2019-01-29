package cheky

// Context keeps query and path parameters.
type Context interface {
	Query(name string) string
	Path(name string) string
}

// Ctx returns new Checker from a given Context.
func Ctx(ctx Context) *Checker {
	return &Checker{
		ctx:    ctx,
		params: []Param{},
	}
}
