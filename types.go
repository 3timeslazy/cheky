package cheky

import (
	"fmt"
	"strconv"
)

const (
	msgCantParse = "can't parse '%s' into %T: %v"
)

// TypesChecker implements type step.
type TypesChecker struct {
	name  string
	where string

	src string

	storeParam func(Param)
}

// Str returns String checker.
func (types TypesChecker) Str(dst *string) *String {
	s := &String{
		name:  types.name,
		where: types.where,

		val: types.src,
		src: dst,

		checks: []check{},
	}

	types.storeParam(s)
	return s
}

// Int returns Int checker.
func (types TypesChecker) Int(dst *int) *Int {
	i := &Int{
		name:  types.name,
		where: types.where,

		src: dst,

		checks: []check{},
	}

	val, err := strconv.ParseInt(types.src, 10, 0)
	if err != nil {
		i.checks = append(i.checks, func() error {
			return fmt.Errorf(msgCantParse, types.src, *dst, err)
		})
	}

	i.val = int(val)

	types.storeParam(i)
	return i
}
