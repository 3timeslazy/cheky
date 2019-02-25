package cheky

// Context keeps query and path parameters.
type Context interface {
	Query(name string) string
	Path(name string) string
}

// Ctx returns new Checker from a given Context.
func Ctx(ctx Context) Checker {
	return Checker{
		ctx:    ctx,
		path:   map[string][]check{},
		query:  map[string][]check{},
		values: map[string][]check{},
	}
}

// WithoutCtx returns new Checker without Context.
//
// Use it when you need to validate with
// Value method.
func WithoutCtx() Checker {
	return Checker{
		ctx:    emptyContext{},
		path:   map[string][]check{},
		query:  map[string][]check{},
		values: map[string][]check{},
	}
}

type emptyContext struct{}

func (ctx emptyContext) Query(name string) string { return "" }

func (ctx emptyContext) Path(name string) string { return "" }
