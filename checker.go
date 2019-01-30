package cheky

// Checker implements Query/Path step of
// the interface of library.
type Checker struct {
	ctx    Context
	params []Param
}

// Param represents any object
// that can be validated.
type Param interface {
	Validate() error
}

// check wraps functions, that
// check values.
type check func() error

// Query returns TypesChecker that checks query params.
func (chr *Checker) Query(name string) TypesChecker {
	return TypesChecker{
		name:  name,
		where: "query",
		src:   chr.ctx.Query(name),
		storeParam: func(p Param) {
			chr.params = append(chr.params, p)
		},
	}
}

// Path returns TypesChecker that checks path params.
func (chr *Checker) Path(name string) TypesChecker {
	return TypesChecker{
		name:  name,
		where: "path",
		src:   chr.ctx.Path(name),
		storeParam: func(p Param) {
			chr.params = append(chr.params, p)
		},
	}
}

// Err run all checks and return error
// if exists.
func (chr Checker) Err() error {
	var err error

	if chr.params != nil {
		for _, param := range chr.params {
			err = param.Validate()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
