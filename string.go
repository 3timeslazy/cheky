package cheky

import (
	"fmt"
	"regexp"
)

// String checks strings.
type String struct {
	name  string
	where string

	val        string
	defaultVal *string

	src *string

	checks []check
}

// Gt checks if length of string greater than i.
func (s *String) Gt(i int) *String {
	s.checks = append(s.checks, func() error {
		if len(s.val) <= i {
			return fmt.Errorf("length should be greater than %d", i)
		}
		return nil
	})
	return s
}

// Ge checks if length of string greater or equal to i.
func (s *String) Ge(i int) *String {
	s.checks = append(s.checks, func() error {
		if len(s.val) < i {
			return fmt.Errorf("length should be greater or equal to %d", i)
		}
		return nil
	})
	return s
}

// Lt checks if length of string less than i.
func (s *String) Lt(i int) *String {
	s.checks = append(s.checks, func() error {
		if len(s.val) >= i {
			return fmt.Errorf("length should be less than %d", i)
		}
		return nil
	})
	return s
}

// Le checks if length of string less or equal to i.
func (s *String) Le(i int) *String {
	s.checks = append(s.checks, func() error {
		if len(s.val) > i {
			return fmt.Errorf("length should be less or equal to %d", i)
		}
		return nil
	})
	return s
}

// OneOf checks if there are value among ones.
func (s *String) OneOf(ones ...string) *String {
	s.checks = append(s.checks, func() error {
		return oneOf(s.val, ones)
	})
	return s
}

// NoOne checks if there are no value among ones.
func (s *String) NoOne(ones ...string) *String {
	s.checks = append(s.checks, func() error {
		return noOne(s.val, ones)
	})
	return s
}

// Regexp checks if value satisfies pattern.
func (s *String) Regexp(pattern string) *String {
	s.checks = append(s.checks, func() error {
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

// Default will set val into dst, if
// error occured during s.Validate() execution.
func (s *String) Default(val string) *String {
	s.defaultVal = &val
	return s
}

// Validate returns result of
// internal checks
func (s *String) Validate() error {
	var err error

	for _, check := range s.checks {
		err = check()
		if err != nil {
			if s.defaultVal == nil {
				return fmt.Errorf("%s '%s': %v", s.where, s.name, err)
			}

			*s.src = *s.defaultVal
			return nil
		}
	}

	*s.src = s.val

	return nil
}
