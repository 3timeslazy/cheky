package cheky

// QPBuilder creates new TBuilders that
// take values from a path or query.
// QPBuilder also keeps TBuilders errors inside.
type QPBuilder interface {
	Query(param string) TBuilder
	Path(param string) TBuilder
	Err() error
}

// TBuilder returns a typed builder:
// Int32Builder, BoolBuilder etc.
type TBuilder interface {
	// keep keeps error into parent QPBuilder
	keep(err error)
}

// Context keeps query and path parameters.
type Context interface {
	Query(name string) string
	Path(name string) string
}
