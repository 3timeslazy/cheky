package cheky

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	msgCantParse = "can't parse '%s' into %T: %v"
)

// TypesChecker implements type step.
type TypesChecker struct {
	src    string
	checks pushCheck
}

// Bool returns Bool checker.
func (types TypesChecker) Bool(dst *bool) Bool {
	b, err := parseBool(types.src)
	if err != nil {
		// push error into root Checker
		types.checks(func() error {
			return fmt.Errorf(msgCantParse, types.src, *dst, err)
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

// Str returns String checker.
func (types TypesChecker) Str(dst *string) String {
	*dst = types.src

	return String{
		checks: types.checks,
		val:    types.src,
	}
}

// Int returns Int checker.
func (types TypesChecker) Int(dst *int) Int {
	i, err := strconv.ParseInt(types.src, 10, 0)
	if err != nil {
		types.checks(func() error {
			return fmt.Errorf(msgCantParse, types.src, *dst, err)
		})
		return Int{
			checks: func(c check) {},
			val:    i,
		}
	}

	*dst = int(i)

	return Int{
		checks: types.checks,
		val:    i,
	}
}

// Int8 returns Int checker.
func (types TypesChecker) Int8(dst *int8) Int {
	i, err := strconv.ParseInt(types.src, 10, 8)
	if err != nil {
		types.checks(func() error {
			return fmt.Errorf(msgCantParse, types.src, *dst, err)
		})
		return Int{
			checks: func(c check) {},
			val:    i,
		}
	}

	*dst = int8(i)

	return Int{
		checks: types.checks,
		val:    i,
	}
}

// Int16 returns Int checker.
func (types TypesChecker) Int16(dst *int16) Int {
	i, err := strconv.ParseInt(types.src, 10, 16)
	if err != nil {
		types.checks(func() error {
			return fmt.Errorf(msgCantParse, types.src, *dst, err)
		})
		return Int{
			checks: func(c check) {},
			val:    i,
		}
	}

	*dst = int16(i)

	return Int{
		checks: types.checks,
		val:    i,
	}
}

// Int32 returns Int checker.
func (types TypesChecker) Int32(dst *int32) Int {
	i, err := strconv.ParseInt(types.src, 10, 32)
	if err != nil {
		types.checks(func() error {
			return fmt.Errorf(msgCantParse, types.src, *dst, err)
		})
		return Int{
			checks: func(c check) {},
			val:    i,
		}
	}

	*dst = int32(i)

	return Int{
		checks: types.checks,
		val:    i,
	}
}

// Int64 returns Int checker.
func (types TypesChecker) Int64(dst *int64) Int {
	i, err := strconv.ParseInt(types.src, 10, 64)
	if err != nil {
		types.checks(func() error {
			return fmt.Errorf(msgCantParse, types.src, *dst, err)
		})
		return Int{
			checks: func(c check) {},
			val:    i,
		}
	}

	*dst = int64(i)

	return Int{
		checks: types.checks,
		val:    i,
	}
}

// Uint returns Uint checker.
func (types TypesChecker) Uint(dst *uint) Uint {
	u, err := strconv.ParseUint(types.src, 10, 0)
	if err != nil {
		types.checks(func() error {
			return fmt.Errorf(msgCantParse, types.src, *dst, err)
		})
		return Uint{
			checks: func(c check) {},
			val:    u,
		}
	}

	*dst = uint(u)

	return Uint{
		checks: types.checks,
		val:    u,
	}
}

// Uint8 returns Int checker.
func (types TypesChecker) Uint8(dst *uint8) Uint {
	u, err := strconv.ParseUint(types.src, 10, 8)
	if err != nil {
		types.checks(func() error {
			return fmt.Errorf(msgCantParse, types.src, *dst, err)
		})
		return Uint{
			checks: func(c check) {},
			val:    u,
		}
	}

	*dst = uint8(u)

	return Uint{
		checks: types.checks,
		val:    u,
	}
}

// Uint16 returns Uint checker.
func (types TypesChecker) Uint16(dst *uint16) Uint {
	u, err := strconv.ParseUint(types.src, 10, 16)
	if err != nil {
		types.checks(func() error {
			return fmt.Errorf(msgCantParse, types.src, *dst, err)
		})
		return Uint{
			checks: func(c check) {},
			val:    u,
		}
	}

	*dst = uint16(u)

	return Uint{
		checks: types.checks,
		val:    u,
	}
}

// Uint32 returns Uint checker.
func (types TypesChecker) Uint32(dst *uint32) Uint {
	u, err := strconv.ParseUint(types.src, 10, 32)
	if err != nil {
		types.checks(func() error {
			return fmt.Errorf(msgCantParse, types.src, *dst, err)
		})
		return Uint{
			checks: func(c check) {},
			val:    u,
		}
	}

	*dst = uint32(u)

	return Uint{
		checks: types.checks,
		val:    u,
	}
}

// Uint64 returns Uint checker.
func (types TypesChecker) Uint64(dst *uint64) Uint {
	u, err := strconv.ParseUint(types.src, 10, 64)
	if err != nil {
		types.checks(func() error {
			return fmt.Errorf(msgCantParse, types.src, *dst, err)
		})
		return Uint{
			checks: func(c check) {},
			val:    u,
		}
	}

	*dst = uint64(u)

	return Uint{
		checks: types.checks,
		val:    u,
	}
}

// Duration returns Duration checker.
func (types TypesChecker) Duration(dst *time.Duration) Duration {
	d, err := time.ParseDuration(types.src)
	if err != nil {
		types.checks(func() error {
			return fmt.Errorf(msgCantParse, types.src, *dst, err)
		})
		return Duration{
			checks: func(c check) {},
			val:    d,
		}
	}

	*dst = d

	return Duration{
		checks: types.checks,
		val:    d,
	}
}

func parseBool(s string) (b bool, err error) {
	switch strings.ToLower(s) {
	case "true", "1":
		b = true
	case "false", "0":
		b = false
	default:
		err = fmt.Errorf("bool can be presented only as 'true','false','1','0', got '%s'", s)
	}

	return
}
