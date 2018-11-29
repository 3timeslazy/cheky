package cheky

// Ctx returns new QPBuilder from a given Context.
func Ctx(ctx Context) QPBuilder {
	return &QueryPath{ctx: ctx}
}
