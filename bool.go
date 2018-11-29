package cheky

import (
	"fmt"
)

// BoolChecker performs boolean values checks.
type BoolChecker interface {
	True() BoolChecker
	False() BoolChecker
	BoolMutator
}

// BoolMutator mutate boolean values.
type BoolMutator interface {
	Not() BoolMutator
}

type booleans struct {
	keeper ErrorKeeper
	val    *bool
}

// True checks if value is true.
func (b *booleans) True() BoolChecker {
	if b.keeper.Err() != nil {
		return b
	}

	if *b.val != true {
		b.keeper.Keep(fmt.Errorf("expected: true, got: false"))
	}
	return b
}

// False checks if value is false.
func (b *booleans) False() BoolChecker {
	if b.keeper.Err() != nil {
		return b
	}

	if *b.val != false {
		b.keeper.Keep(fmt.Errorf("expected: false, got: true"))
	}
	return b
}

// Returns true if value is false.
// Return false, otherwise.
func (b *booleans) Not() BoolMutator {
	if b.keeper.Err() != nil {
		return b
	}

	*b.val = !*b.val
	return b
}
