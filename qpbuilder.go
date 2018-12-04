package cheky

import (
	"fmt"
)

const (
	path  = "path"
	query = "query"
)

// QueryPath implements QPBuilder interface.
type queryPath struct {
	err error
	val string
	ctx Context
}

// Query implements QPBuilder's Query.
func (qpb *queryPath) Query(param string) TBuilder {
	qpb.val = qpb.ctx.Query(param)
	return &types{val: qpb.val, errkp: &errorFormatter{
		where: query,
		name:  param,
		qpb:   qpb,
	}}
}

// Path implements QPBuilder's Path.
func (qpb *queryPath) Path(param string) TBuilder {
	qpb.val = qpb.ctx.Path(param)
	return &types{val: qpb.val, errkp: &errorFormatter{
		where: path,
		name:  param,
		qpb:   qpb,
	}}
}

// Err returns keeped error.
func (qpb *queryPath) Err() error {
	return qpb.err
}

// Keep keeps error inside.
func (qpb *queryPath) Keep(err error) {
	if qpb.err == nil && err != nil {
		qpb.err = err
	}
}

// errorFormatter formats the error message,
// substituting additional information there.
// Also implements ErrorKeeper interface.
type errorFormatter struct {
	where string
	name  string
	qpb   *queryPath
}

func (errfmt *errorFormatter) Keep(err error) {
	if errfmt.qpb.err == nil && err != nil {
		errfmt.qpb.err = fmt.Errorf("%s '%s': %v", errfmt.where, errfmt.name, err)
	}
}

func (errfmt *errorFormatter) Err() error {
	return errfmt.qpb.err
}
