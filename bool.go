package cheky

import (
	"fmt"
)

// Bool checks booleans.
type Bool struct {
	checks pushCheck
	val    bool
}

// True checks if value is true.
func (b Bool) True() Bool {
	b.checks(func() error {
		if b.val != true {
			return fmt.Errorf("expected 'true', got 'false'")
		}
		return nil
	})
	return b
}

// False checks if value is false.
func (b Bool) False() Bool {
	b.checks(func() error {
		if b.val != false {
			return fmt.Errorf("expected 'false', got 'true'")
		}
		return nil
	})
	return b
}
