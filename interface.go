package cheky

// QPBuilder creates new TBuilders that
// take values from a path or query.
// QPBuilder also keeps TBuilders errors inside.
type QPBuilder interface {
	ErrorKeeper
	Query(param string) TBuilder
	Path(param string) TBuilder
}

// TBuilder returns a typed builder:
// IntChecker, BoolChecker etc.
type TBuilder interface {
	ErrorKeeper
	Bool(*bool) BoolChecker
}

// ErrorKeeper keeps and returns errors.
type ErrorKeeper interface {
	Keep(err error)
	Err() error
}
