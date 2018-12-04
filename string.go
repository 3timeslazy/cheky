package cheky

import (
	"fmt"
)

// StringChecker performs string values checks.
type StringChecker interface {
	Gt(i int) StringChecker
	Ge(i int) StringChecker
	Lt(i int) StringChecker
	Le(i int) StringChecker
	// OneOf(strs ...string) StringChecker
	// NoOne(strs ...string) StringChecker
	// Regexp(str string) StringChecker
}

type stringChecker struct {
	errkp ErrorKeeper
	val   *string
}

func (ch *stringChecker) Gt(i int) StringChecker {
	if ch.errkp.Err() != nil {
		return ch
	}

	l := len(*ch.val)
	if l <= i {
		ch.errkp.Keep(fmt.Errorf("length. expected > '%d', got '%d'", i, l))
	}
	return ch
}

func (ch *stringChecker) Ge(i int) StringChecker {
	if ch.errkp.Err() != nil {
		return ch
	}

	l := len(*ch.val)
	if l < i {
		ch.errkp.Keep(fmt.Errorf("length. expected >= '%d', got '%d'", i, l))
	}
	return ch
}

func (ch *stringChecker) Lt(i int) StringChecker {
	if ch.errkp.Err() != nil {
		return ch
	}

	l := len(*ch.val)
	if l >= i {
		ch.errkp.Keep(fmt.Errorf("length. expected < '%d', got '%d'", i, l))
	}
	return ch
}

func (ch *stringChecker) Le(i int) StringChecker {
	if ch.errkp.Err() != nil {
		return ch
	}

	l := len(*ch.val)
	if l > i {
		ch.errkp.Keep(fmt.Errorf("length. expected <= '%d', got '%d'", i, l))
	}
	return ch
}
