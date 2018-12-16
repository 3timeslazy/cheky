package cheky_test

import (
	"testing"

	"github.com/3timeslazy/cheky"

	. "github.com/smartystreets/goconvey/convey"
)

type CustomContext struct{ P, Q string }

func (cctx *CustomContext) Query(name string) string {
	return cctx.Q
}

func (cctx *CustomContext) Path(name string) string {
	return cctx.P
}

func NewContext(q, p string) cheky.Context {
	return &CustomContext{P: p, Q: q}
}

func TestBool(t *testing.T) {
	Convey("param=true", t, func() {
		check := newCheck("true")
		b := false

		check.Query("true").Bool(&b).True()

		So(check.Err(), ShouldBeNil)
		So(b, ShouldBeTrue)
	})

	Convey("param=false", t, func() {
		check := newCheck("false")
		b := false

		check.Query("false").Bool(&b).False()

		So(check.Err(), ShouldBeNil)
		So(b, ShouldBeFalse)
	})

	Convey("param=qwerty", t, func() {
		check := newCheck("qwerty")
		b := false

		check.Query("qwerty").Bool(&b)

		So(check.Err(), ShouldBeError)
		So(b, ShouldBeFalse)
	})

	Convey("param=true, but need false", t, func() {
		check := newCheck("true")
		b := false

		check.Query("true").Bool(&b).False()

		So(check.Err(), ShouldBeError)
		So(b, ShouldBeTrue)
	})

	Convey("param=false, but need true", t, func() {
		check := newCheck("false")
		b := false

		check.Query("true").Bool(&b).True()

		So(check.Err(), ShouldBeError)
		So(b, ShouldBeFalse)
	})

	Convey("param=1", t, func() {
		check := newCheck("1")
		b := false

		check.Query("1").Bool(&b)

		So(check.Err(), ShouldBeNil)
		So(b, ShouldBeTrue)
	})

	Convey("param=tRuE", t, func() {
		check := newCheck("tRuE")
		b := false

		check.Query("tRuE").Bool(&b)

		So(check.Err(), ShouldBeNil)
		So(b, ShouldBeTrue)
	})
}
