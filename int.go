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
	checks pushCheck
	val    int64
}

// Gt checks if value greater than other.
func (i Int) Gt(other int64) Int {
	i.checks(func() error {
		if i.val <= other {
			return fmt.Errorf(
				"value should be greater than '%d', got '%d'",
				other, i.val,
			)
		}
		return nil
	})
	return i
}

// Ge checks if value greater or equal
// to other.
func (i Int) Ge(other int64) Int {
	i.checks(func() error {
		if i.val < other {
			return fmt.Errorf(
				"value should be greater or equal to '%d', got '%d'",
				other, i.val,
			)
		}
		return nil
	})
	return i
}

// Lt checks if value less than other.
func (i Int) Lt(other int64) Int {
	i.checks(func() error {
		if i.val >= other {
			return fmt.Errorf(
				"value should be less than '%d', got '%d'",
				other, i.val,
			)
		}
		return nil
	})
	return i
}

// Le checks if value less or equal
// to other.
func (i Int) Le(other int64) Int {
	i.checks(func() error {
		if i.val > other {
			return fmt.Errorf(
				"value should be less or equal to '%d', got '%d'",
				other, i.val,
			)
		}
		return nil
	})
	return i
}

// OneOf checks if there are value among ones.
func (i Int) OneOf(ones ...int64) Int {
	i.checks(func() error {
		for _, one := range ones {
			if i.val == one {
				return nil
			}
		}
		return fmt.Errorf("value should be one of %v, got '%d'", ones, i.val)
	})
	return i
}

// NoOne checks if there are no value among ones.
func (i Int) NoOne(ones ...int64) Int {
	i.checks(func() error {
		for _, one := range ones {
			if i.val == one {
				return fmt.Errorf("value should not be one of %v, got '%d'", ones, i.val)
			}
		}
		return nil
	})
	return i
}