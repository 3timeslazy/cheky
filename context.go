package cheky

// Context keeps query and path parameters.
type Context interface {
	Query(name string) string
	Path(name string) string
}

// Ctx returns new QPBuilder from a given Context.
func Ctx(ctx Context) QPBuilder {
	return &queryPath{ctx: ctx}
}
