package cheky_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInt(t *testing.T) {
	tcs := []*IntTC{
		&IntTC{
			Desc:           "parse '10' into 10",
			ValidatedValue: "10",
			ExpectedAs:     10,
			AfterMethods:   nil,
			Error:          ShouldBeNil,
		},
		&IntTC{
			Desc:           "parse '-8' into -8",
			ValidatedValue: "-8",
			ExpectedAs:     -8,
			AfterMethods:   nil,
			Error:          ShouldBeNil,
		},
		&IntTC{
			Desc:           "11 Gt 10",
			ValidatedValue: "11",
			ExpectedAs:     11,
			AfterMethods: map[string][]interface{}{
				"Gt": args(10),
			},
			Error: ShouldBeNil,
		},
		&IntTC{
			Desc:           "10 Ge 10",
			ValidatedValue: "10",
			ExpectedAs:     10,
			AfterMethods: map[string][]interface{}{
				"Ge": args(10),
			},
			Error: ShouldBeNil,
		},
		&IntTC{
			Desc:           "8 Lt 10",
			ValidatedValue: "8",
			ExpectedAs:     8,
			AfterMethods: map[string][]interface{}{
				"Lt": args(10),
			},
			Error: ShouldBeNil,
		},
		&IntTC{
			Desc:           "10 Le 10",
			ValidatedValue: "10",
			ExpectedAs:     10,
			AfterMethods: map[string][]interface{}{
				"Le": args(10),
			},
			Error: ShouldBeNil,
		},
		&IntTC{
			Desc:           "6 Gt 10",
			ValidatedValue: "6",
			ExpectedAs:     0,
			AfterMethods: map[string][]interface{}{
				"Gt": args(10),
			},
			Error: ShouldBeError,
		},
		&IntTC{
			Desc:           "10 Gt 10",
			ValidatedValue: "10",
			ExpectedAs:     0,
			AfterMethods: map[string][]interface{}{
				"Gt": args(10),
			},
			Error: ShouldBeError,
		},
		&IntTC{
			Desc:           "parse 'nan'",
			ValidatedValue: "nan",
			ExpectedAs:     0,
			AfterMethods:   nil,
			Error:          ShouldBeError,
		},
		&IntTC{
			Desc:           "should be 1, 2 or 3",
			ValidatedValue: "3",
			ExpectedAs:     3,
			AfterMethods: map[string][]interface{}{
				"OneOf": args(1, 2, 3),
			},
			Error: ShouldBeNil,
		},
		&IntTC{
			Desc:           "when not one of 1, 2 or 3",
			ValidatedValue: "4",
			ExpectedAs:     0,
			AfterMethods: map[string][]interface{}{
				"OneOf": args(1, 2, 3),
			},
			Error: ShouldBeError,
		},
		&IntTC{
			Desc:           "should not be 1, 2 or 3",
			ValidatedValue: "4",
			ExpectedAs:     4,
			AfterMethods: map[string][]interface{}{
				"NoOne": args(1, 2, 3),
			},
			Error: ShouldBeNil,
		},
		&IntTC{
			Desc:           "when one of 1, 2 or 3",
			ValidatedValue: "3",
			ExpectedAs:     0,
			AfterMethods: map[string][]interface{}{
				"NoOne": args(1, 2, 3),
			},
			Error: ShouldBeError,
		},
		&IntTC{
			Desc:           "default when error",
			ValidatedValue: "3",
			ExpectedAs:     13,
			AfterMethods: map[string][]interface{}{
				"NoOne":   args(1, 2, 3),
				"Default": args(13),
			},
			Error: ShouldBeNil,
		},
		&IntTC{
			Desc:           "default when no error",
			ValidatedValue: "3",
			ExpectedAs:     3,
			AfterMethods: map[string][]interface{}{
				"OneOf":   args(1, 2, 3),
				"Default": args(13),
			},
			Error: ShouldBeNil,
		},
	}

	for _, tc := range tcs {
		tc.Do(t)
	}
}

type IntTC struct {
	Desc           string
	ValidatedValue string
	ExpectedAs     int
	AfterMethods   map[string][]interface{}
	Error          func(interface{}, ...interface{}) string
}

func (tc *IntTC) Do(t *testing.T) {
	Convey(tc.Desc, t, func() {
		var check = newCheck(tc.ValidatedValue)
		var i int
		var v = check.Query(tc.ValidatedValue).Int(&i)

		runMethods(v, tc.AfterMethods)

		So(check.Err(), tc.Error)
		So(i, ShouldEqual, tc.ExpectedAs)
	})
}
