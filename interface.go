package cheky

// ErrorKeeper keeps errors.
type ErrorKeeper interface {
	// Keep keeps errors.
	Keep(err error)
}

// QPBuilder creates new TBuilders that
// take values from a path or query.
// QPBuilder also keeps TBuilders errors inside.
type QPBuilder interface {
	Query(param string) TBuilder
	Path(param string) TBuilder
}

// TBuilder returns a typed builder:
// IntBuilder, BoolBuilder etc.
type TBuilder interface {
}

// Context keeps query and path parameters.
type Context interface {
	Query(name string) string
	Path(name string) string
}
