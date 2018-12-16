package cheky

import (
	"fmt"
	"regexp"
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

// OneOf checks if there are value among ones.
func (s String) OneOf(ones ...string) String {
	s.checks(func() error {
		return oneOf(s.val, ones)
	})
	return s
}

// NoOne checks if there are no value among ones.
func (s String) NoOne(ones ...string) String {
	s.checks(func() error {
		return noOne(s.val, ones)
	})
	return s
}

// Regexp checks if value satisfies pattern.
func (s String) Regexp(pattern string) String {
	s.checks(func() error {
		matched, err := regexp.Match(pattern, []byte(s.val))
		if err != nil {
			return fmt.Errorf("regexp: %v", err)
		}
		if !matched {
			return fmt.Errorf("regexp: pattern not match")
		}
		return nil
	})
	return s
}
