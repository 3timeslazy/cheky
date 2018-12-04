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
	init := func(param string) (cheky.QPBuilder, bool) {
		return cheky.Ctx(NewContext(param, "any")), false
	}

	Convey("param=true", t, func() {
		check, b := init("true")

		check.Query("true").Bool(&b).True()

		So(check.Err(), ShouldBeNil)
		So(b, ShouldBeTrue)
	})

	Convey("param=false", t, func() {
		check, b := init("false")

		check.Query("false").Bool(&b).False()

		So(check.Err(), ShouldBeNil)
		So(b, ShouldBeFalse)
	})

	Convey("param=qwerty", t, func() {
		check, b := init("qwerty")

		check.Query("qwerty").Bool(&b)

		So(check.Err(), ShouldBeError)
		So(b, ShouldBeFalse)
	})

	Convey("param=true, but need false", t, func() {
		check, b := init("true")

		check.Query("true").Bool(&b).False()

		So(check.Err(), ShouldBeError)
		So(b, ShouldBeTrue)
	})

	Convey("param=1", t, func() {
		check, b := init("1")

		check.Query("1").Bool(&b)

		So(check.Err(), ShouldBeNil)
		So(b, ShouldBeTrue)
	})

	Convey("param=true, but need false", t, func() {
		check, b := init("true")

		check.Query("true").Bool(&b).Not()

		So(check.Err(), ShouldBeNil)
		So(b, ShouldBeFalse)
	})

	Convey("param=tRuE", t, func() {
		check, b := init("tRue")

		check.Query("tRuE").Bool(&b)

		So(check.Err(), ShouldBeNil)
		So(b, ShouldBeTrue)
	})
}
