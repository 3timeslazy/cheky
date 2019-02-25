package cheky

import (
	"fmt"
)

// Checker implements Query/Path step of
// the interface of library.
type Checker struct {
	ctx Context

	path   map[string][]check
	query  map[string][]check
	values map[string][]check
}

// check wraps functions, that
// check values.
type check func() error

// pushCheck pushes c into root Checker.
type pushCheck func(c check)

// Query returns TypesChecker that checks query params.
func (chr Checker) Query(name string) TypesChecker {
	return TypesChecker{
		src: chr.ctx.Query(name),
		checks: func(c check) {
			chr.query[name] = append(chr.query[name], c)
		},
	}
}

// Path returns TypesChecker that checks path params.
func (chr Checker) Path(name string) TypesChecker {
	return TypesChecker{
		src: chr.ctx.Path(name),
		checks: func(c check) {
			chr.path[name] = append(chr.path[name], c)
		},
	}
}

// Value returns TypesChecker that checks v.
func (chr Checker) Value(v string) TypesChecker {
	return TypesChecker{
		src: v,
		checks: func(c check) {
			chr.values[v] = append(chr.values[v], c)
		},
	}
}

// Err run all checks and return error
// if exists.
func (chr Checker) Err() error {
	var err error

	for param, chks := range chr.path {
		for _, chk := range chks {
			if err = chk(); err != nil {
				return fmt.Errorf("path '%s': %v", param, err)
			}
		}
	}

	for param, chks := range chr.query {
		for _, chk := range chks {
			if err = chk(); err != nil {
				return fmt.Errorf("query '%s': %v", param, err)
			}
		}
	}

	for param, chks := range chr.values {
		for _, chk := range chks {
			if err = chk(); err != nil {
				return fmt.Errorf("value '%s': %v", param, err)
			}
		}
	}

	return nil
}
