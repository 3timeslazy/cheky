package cheky

// Uint checks follow types:
// - uint
// - uint8
// - uint16
// - uint32
// - uint64
type Uint struct {
	checks pushCheck
	val    uint64
}

// Gt checks if value greater than other.
func (u Uint) Gt(other uint64) Uint {
	u.checks(func() error {
		return gt(u.val, other)
	})
	return u
}

// Ge checks if value greater or equal
// to other.
func (u Uint) Ge(other uint64) Uint {
	u.checks(func() error {
		return ge(u.val, other)
	})
	return u
}

// Lt checks if value less than other.
func (u Uint) Lt(other uint64) Uint {
	u.checks(func() error {
		return lt(u.val, other)
	})
	return u
}

// Le checks if value less or equal
// to other.
func (u Uint) Le(other uint64) Uint {
	u.checks(func() error {
		return le(u.val, other)
	})
	return u
}

// OneOf checks if there are value among ones.
func (u Uint) OneOf(ones ...uint64) Uint {
	u.checks(func() error {
		return oneOf(u.val, ones)
	})
	return u
}

// NoOne checks if there are no value among ones.
func (u Uint) NoOne(ones ...uint64) Uint {
	u.checks(func() error {
		return noOne(u.val, ones)
	})
	return u
}
