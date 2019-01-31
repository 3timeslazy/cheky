package cheky_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestString(t *testing.T) {
	tcs := []*StringTC{
		&StringTC{
			Desc:           "len('qwerty') Ge 6",
			ValidatedValue: "qwerty",
			ExpectedAs:     "qwerty",
			AfterMethods: map[string][]interface{}{
				"Ge": args(6),
			},
			Error: ShouldBeNil,
		},
		&StringTC{
			Desc:           "len('qwerty') Ge 7",
			ValidatedValue: "qwerty",
			ExpectedAs:     "",
			AfterMethods: map[string][]interface{}{
				"Ge": args(7),
			},
			Error: ShouldBeError,
		},
		&StringTC{
			Desc:           "len('qwerty') Gt 5",
			ValidatedValue: "qwerty",
			ExpectedAs:     "qwerty",
			AfterMethods: map[string][]interface{}{
				"Gt": args(5),
			},
			Error: ShouldBeNil,
		},
		&StringTC{
			Desc:           "len('qwerty') Gt 6",
			ValidatedValue: "qwerty",
			ExpectedAs:     "",
			AfterMethods: map[string][]interface{}{
				"Gt": args(6),
			},
			Error: ShouldBeError,
		},
		&StringTC{
			Desc:           "len('qwerty') Le 6",
			ValidatedValue: "qwerty",
			ExpectedAs:     "qwerty",
			AfterMethods: map[string][]interface{}{
				"Le": args(6),
			},
			Error: ShouldBeNil,
		},
		&StringTC{
			Desc:           "len('qwerty') Le 5",
			ValidatedValue: "qwerty",
			ExpectedAs:     "",
			AfterMethods: map[string][]interface{}{
				"Le": args(5),
			},
			Error: ShouldBeError,
		},
		&StringTC{
			Desc:           "len('qwerty') Lt 8",
			ValidatedValue: "qwerty",
			ExpectedAs:     "qwerty",
			AfterMethods: map[string][]interface{}{
				"Lt": args(8),
			},
			Error: ShouldBeNil,
		},
		&StringTC{
			Desc:           "len('qwerty') Lt 6",
			ValidatedValue: "qwerty",
			ExpectedAs:     "",
			AfterMethods: map[string][]interface{}{
				"Lt": args(6),
			},
			Error: ShouldBeError,
		},
		&StringTC{
			Desc:           "should be one of 'a', 'b'",
			ValidatedValue: "a",
			ExpectedAs:     "a",
			AfterMethods: map[string][]interface{}{
				"OneOf": args("a", "b"),
			},
			Error: ShouldBeNil,
		},
		&StringTC{
			Desc:           "should be one of 'a', 'b', but got 'c'",
			ValidatedValue: "c",
			ExpectedAs:     "",
			AfterMethods: map[string][]interface{}{
				"OneOf": args("a", "b"),
			},
			Error: ShouldBeError,
		},
		&StringTC{
			Desc:           "should not be 'a' or 'b'",
			ValidatedValue: "c",
			ExpectedAs:     "c",
			AfterMethods: map[string][]interface{}{
				"NoOne": args("a", "b"),
			},
			Error: ShouldBeNil,
		},
		&StringTC{
			Desc:           "should not be 'a' or 'b', but got 'a'",
			ValidatedValue: "a",
			ExpectedAs:     "",
			AfterMethods: map[string][]interface{}{
				"NoOne": args("a", "b"),
			},
			Error: ShouldBeError,
		},
		&StringTC{
			Desc:           "'Hello, World!' not satisfies '^[0-9]$'",
			ValidatedValue: "Hello, World!",
			ExpectedAs:     "",
			AfterMethods: map[string][]interface{}{
				"Regexp": args("^[0-9]$"),
			},
			Error: ShouldBeError,
		},
		&StringTC{
			Desc:           "invalid regexp pattern",
			ValidatedValue: "Hello, World!",
			ExpectedAs:     "",
			AfterMethods: map[string][]interface{}{
				"Regexp": args("^[0-9"),
			},
			Error: ShouldBeError,
		},
		&StringTC{
			Desc:           "'Hello, World!' satisfies '^Hello, .+!$'",
			ValidatedValue: "Hello, World!",
			ExpectedAs:     "Hello, World!",
			AfterMethods: map[string][]interface{}{
				"Regexp": args("^Hello, .+!$"),
			},
			Error: ShouldBeNil,
		},
		&StringTC{
			Desc:           "default when there is no error",
			ValidatedValue: "abc",
			ExpectedAs:     "abc",
			AfterMethods: map[string][]interface{}{
				"OneOf":   args("abc"),
				"Default": args("cba"),
			},
			Error: ShouldBeNil,
		},
		&StringTC{
			Desc:           "default when there is error",
			ValidatedValue: "a",
			ExpectedAs:     "cba",
			AfterMethods: map[string][]interface{}{
				"OneOf":   args("abc"),
				"Default": args("cba"),
			},
			Error: ShouldBeNil,
		},
	}

	for _, tc := range tcs {
		tc.Do(t)
	}
}

type StringTC struct {
	Desc           string
	ValidatedValue string
	ExpectedAs     string
	AfterMethods   map[string][]interface{}
	Error          func(interface{}, ...interface{}) string
}

func (tc *StringTC) Do(t *testing.T) {
	Convey(tc.Desc, t, func() {
		var check = newCheck(tc.ValidatedValue)
		var s string
		var v = check.Query(tc.ValidatedValue).Str(&s)

		runMethods(v, tc.AfterMethods)

		So(check.Err(), tc.Error)
		So(s, ShouldEqual, tc.ExpectedAs)
	})
}
