package cheky

import (
	"fmt"
)

// String checks strings.
type String struct {
	checks pushCheck
	val    string
}

// Gt checks if length of string greater than i.
func (s String) Gt(i int) String {
	s.checks(func() error {
		if len(s.val) <= i {
			return fmt.Errorf("length should be greater than %d", i)
		}
		return nil
	})
	return s
}

// Ge checks if length of string greater or equal to i.
func (s String) Ge(i int) String {
	s.checks(func() error {
		if len(s.val) < i {
			return fmt.Errorf("length should be greater or equal to %d", i)
		}
		return nil
	})
	return s
}

// Lt checks if length of string less than i.
func (s String) Lt(i int) String {
	s.checks(func() error {
		if len(s.val) >= i {
			return fmt.Errorf("length should be less than %d", i)
		}
		return nil
	})
	return s
}

// Le checks if length of string less or equal to i.
func (s String) Le(i int) String {
	s.checks(func() error {
		if len(s.val) > i {
			return fmt.Errorf("length should be less or equal to %d", i)
		}
		return nil
	})
	return s
}
