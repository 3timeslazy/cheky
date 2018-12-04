package cheky

import (
	"fmt"
	"strings"
)

// TypesChecker implements type step.
type TypesChecker struct {
	src    string
	checks pushCheck
}

// Bool returns boolean checkers.
func (types TypesChecker) Bool(dst *bool) Bool {
	b, err := parseBool(types.src)
	if err != nil {
		// push error into root Checker
		types.checks(func() error {
			return fmt.Errorf("can't parse '%s' into bool", types.src)
		})
		return Bool{
			checks: func(c check) {},
			val:    b,
		}
	}

	*dst = b

	return Bool{
		checks: types.checks,
		val:    b,
	}
}

func parseBool(s string) (b bool, err error) {
	switch strings.ToLower(s) {
	case "true", "1":
		b = true
	case "false", "0":
		b = false
	default:
		err = fmt.Errorf("value '%s' isn't boolean", s)
	}
	return
}
