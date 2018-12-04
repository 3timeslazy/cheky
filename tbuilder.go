package cheky

import (
	"fmt"
	"strings"
)

// types implements TBuilder interface.
type types struct {
	errkp ErrorKeeper
	val   string
}

// Keep keeps an error in a
// nested ErrorKeeper.
func (t *types) Keep(err error) {
	t.errkp.Keep(err)
}

// Err returns an error stored in
// a nested ErrorKeeper
func (t *types) Err() error {
	return t.errkp.Err()
}

// Bool returns BoolChecker.
// Result value will be stored into b.
// Available bool values:
// * True or False in any case
// * 1 or 0
func (t *types) Bool(boolean *bool) BoolChecker {
	b, err := parseBool(t.val)
	if err != nil {
		t.Keep(err)
	}
	*boolean = b
	return &booleans{val: boolean, keeper: t}
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

// Str returns StringChecker.
// Result value will be stored into str.
func (t *types) Str(str *string) StringChecker {
	*str = t.val
	return &stringChecker{
		errkp: t,
		val:   str,
	}
}
