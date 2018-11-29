package cheky

const (
	path  = "path"
	query = "query"
)

// QueryPath implements QPBuilder interface.
type QueryPath struct {
	err error
	val string
	ctx Context
}

// Query implements QPBuilder's Query.
func (qpb *QueryPath) Query(param string) TBuilder {
	qpb.val = qpb.ctx.Query(param)
	return &Type{keeper: qpb, val: qpb.val}
}

// Path implements QPBuilder's Path.
func (qpb *QueryPath) Path(param string) TBuilder {
	qpb.val = qpb.ctx.Path(param)
	return &Type{keeper: qpb, val: qpb.val}
}

// Err returns keeped error.
func (qpb *QueryPath) Err() error {
	return qpb.err
}

// Keep keeps error inside.
func (qpb *QueryPath) Keep(err error) {
	if qpb.err == nil && err != nil {
		qpb.err = err
	}
}
