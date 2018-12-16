package cheky_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestString(t *testing.T) {
	Convey("param=qwerty", t, func() {
		check := newCheck("qwerty")
		s := ""

		check.Query("qwerty").Str(&s).Ge(6)
		So(check.Err(), ShouldBeNil)
		So(s, ShouldEqual, "qwerty")
	})

	Convey("param=qwerty", t, func() {
		check := newCheck("qwerty")
		s := ""

		check.Query("qwerty").Str(&s).Ge(7)
		So(check.Err(), ShouldBeError)
		So(s, ShouldEqual, "qwerty")
	})

	Convey("param=qwerty", t, func() {
		check := newCheck("qwerty")
		s := ""

		check.Query("qwerty").Str(&s).Gt(5)
		So(check.Err(), ShouldBeNil)
		So(s, ShouldEqual, "qwerty")
	})

	Convey("param=qwerty", t, func() {
		check := newCheck("qwerty")
		s := ""

		check.Query("qwerty").Str(&s).Gt(6)
		So(check.Err(), ShouldBeError)
		So(s, ShouldEqual, "qwerty")
	})

	Convey("param=qwerty", t, func() {
		check := newCheck("qwerty")
		s := ""

		check.Query("qwerty").Str(&s).Le(6)
		So(check.Err(), ShouldBeNil)
		So(s, ShouldEqual, "qwerty")
	})

	Convey("param=qwerty", t, func() {
		check := newCheck("qwerty")
		s := ""

		check.Query("qwerty").Str(&s).Le(5)
		So(check.Err(), ShouldBeError)
		So(s, ShouldEqual, "qwerty")
	})

	Convey("param=qwerty", t, func() {
		check := newCheck("qwerty")
		s := ""

		check.Query("qwerty").Str(&s).Lt(8)
		So(check.Err(), ShouldBeNil)
		So(s, ShouldEqual, "qwerty")
	})

	Convey("param=qwerty", t, func() {
		check := newCheck("qwerty")
		s := ""

		check.Query("qwerty").Str(&s).Lt(6)
		So(check.Err(), ShouldBeError)
		So(s, ShouldEqual, "qwerty")
	})

	Convey("should be one of 'a', 'b'", t, func() {
		check := newCheck("a")
		s := ""

		check.Query("a").Str(&s).OneOf("a", "b")
		So(check.Err(), ShouldBeNil)
		So(s, ShouldEqual, "a")
	})

	Convey("should be one of 'a', 'b', but got 'c'", t, func() {
		check := newCheck("c")
		s := ""

		check.Query("c").Str(&s).OneOf("a", "b")
		So(check.Err(), ShouldBeError)
		So(s, ShouldEqual, "c")
	})

	Convey("should not be 'a' or 'b'", t, func() {
		check := newCheck("c")
		s := ""

		check.Query("c").Str(&s).NoOne("a", "b")
		So(check.Err(), ShouldBeNil)
		So(s, ShouldEqual, "c")
	})

	Convey("should not be 'a' or 'b', but got 'a'", t, func() {
		check := newCheck("a")
		s := ""

		check.Query("a").Str(&s).NoOne("a", "b")
		So(check.Err(), ShouldBeError)
		So(s, ShouldEqual, "a")
	})

	Convey("should satisfies pattern", t, func() {
		check := newCheck("Hello, World!")
		s := ""

		check.Query("hello").Str(&s).Regexp("^Hello, .+!$")
		So(check.Err(), ShouldBeNil)
		So(s, ShouldEqual, "Hello, World!")
	})

	Convey("should not satisfies pattern", t, func() {
		check := newCheck("Hello, World!")
		s := ""

		check.Query("hello").Str(&s).Regexp("^[0-9]$")
		So(check.Err(), ShouldBeError)
		So(s, ShouldEqual, "Hello, World!")
	})

	Convey("invalid pattern", t, func() {
		check := newCheck("Hello, World!")
		s := ""

		check.Query("hello").Str(&s).Regexp("[0-9")
		So(check.Err(), ShouldBeError)
		So(s, ShouldEqual, "Hello, World!")
	})
}
