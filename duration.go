package cheky

import "time"

// Duration check time.Duration.
type Duration struct {
	checks pushCheck
	val    time.Duration
}

// Gt checks if value greater than other.
func (d Duration) Gt(other time.Duration) Duration {
	d.checks(func() error {
		return gt(d.val, other)
	})
	return d
}

// Ge checks if value greater or equal
// to other.
func (d Duration) Ge(other time.Duration) Duration {
	d.checks(func() error {
		return ge(d.val, other)
	})
	return d
}

// Lt checks if value less than other.
func (d Duration) Lt(other time.Duration) Duration {
	d.checks(func() error {
		return lt(d.val, other)
	})
	return d
}

// Le checks if value less or equal
// to other.
func (d Duration) Le(other time.Duration) Duration {
	d.checks(func() error {
		return le(d.val, other)
	})
	return d
}

// OneOf checks if there are value among ones.
func (d Duration) OneOf(ones ...time.Duration) Duration {
	d.checks(func() error {
		return oneOf(d.val, ones)
	})
	return d
}

// NoOne checks if there are no value among ones.
func (d Duration) NoOne(ones ...time.Duration) Duration {
	d.checks(func() error {
		return noOne(d.val, ones)
	})
	return d
}
