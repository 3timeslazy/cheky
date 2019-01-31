package cheky

import (
	"fmt"
)

// Int checks follow types:
// - int
// - int8
// - int16
// - int32
// - int64
type Int struct {
	name  string
	where string

	val        int
	defaultVal *int

	src *int

	checks []check
}

// Gt checks if value greater than other.
func (i *Int) Gt(other int) *Int {
	i.checks = append(i.checks, func() error {
		return gt(i.val, other)
	})
	return i
}

// Ge checks if value greater or equal
// to other.
func (i *Int) Ge(other int) *Int {
	i.checks = append(i.checks, func() error {
		return ge(i.val, other)
	})
	return i
}

// Lt checks if value less than other.
func (i *Int) Lt(other int) *Int {
	i.checks = append(i.checks, func() error {
		return lt(i.val, other)
	})
	return i
}

// Le checks if value less or equal
// to other.
func (i *Int) Le(other int) *Int {
	i.checks = append(i.checks, func() error {
		return le(i.val, other)
	})
	return i
}

// OneOf checks if there are value among ones.
func (i *Int) OneOf(ones ...int) *Int {
	i.checks = append(i.checks, func() error {
		return oneOf(i.val, ones)
	})
	return i
}

// NoOne checks if there are no value among ones.
func (i *Int) NoOne(ones ...int) *Int {
	i.checks = append(i.checks, func() error {
		return noOne(i.val, ones)
	})
	return i
}

// Default will set val into dst, if
// error occured during s.Validate() execution.
func (i *Int) Default(val int) *Int {
	i.defaultVal = &val
	return i
}

// Validate returns result of
// internal checks
func (i *Int) Validate() error {
	var err error

	for _, check := range i.checks {
		err = check()
		if err != nil {
			if i.defaultVal == nil {
				return fmt.Errorf("%s '%s': %v", i.where, i.name, err)
			}

			*i.src = *i.defaultVal
			return nil
		}
	}

	*i.src = i.val

	return nil
}
