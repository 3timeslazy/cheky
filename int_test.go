package cheky_test

import (
	"testing"

	"github.com/3timeslazy/cheky"

	. "github.com/smartystreets/goconvey/convey"
)

func newCheck(param string) cheky.Checker {
	return cheky.Ctx(NewContext(param, "any"))
}

func TestInt(t *testing.T) {
	Convey("param=10", t, func() {
		check := newCheck("10")
		i := 0

		check.Query("10").Int(&i)
		So(check.Err(), ShouldBeNil)
		So(i, ShouldEqual, 10)
	})

	Convey("param=-8", t, func() {
		check := newCheck("-8")
		i := 0

		check.Query("-8").Int(&i)
		So(check.Err(), ShouldBeNil)
		So(i, ShouldEqual, -8)
	})

	Convey("should be greater than 10", t, func() {
		check := newCheck("11")
		i := 0

		check.Query("11").Int(&i).Gt(10)
		So(check.Err(), ShouldBeNil)
		So(i, ShouldEqual, 11)
	})

	Convey("should be greater or equal to 10", t, func() {
		check := newCheck("10")
		i := 0

		check.Query("10").Int(&i).Ge(10)
		So(check.Err(), ShouldBeNil)
		So(i, ShouldEqual, 10)
	})

	Convey("should be less than 10", t, func() {
		check := newCheck("8")
		i := 0

		check.Query("8").Int(&i).Lt(10)
		So(check.Err(), ShouldBeNil)
		So(i, ShouldEqual, 8)
	})

	Convey("should be less or equal to 10", t, func() {
		check := newCheck("10")
		i := 0

		check.Query("10").Int(&i).Le(10)
		So(check.Err(), ShouldBeNil)
		So(i, ShouldEqual, 10)
	})

	Convey("should be greater than 10, but got less", t, func() {
		check := newCheck("6")
		i := 0

		check.Query("6").Int(&i).Gt(10)
		So(check.Err(), ShouldBeError)
		So(i, ShouldEqual, 6)
	})

	Convey("should be greater than 10, but got 10", t, func() {
		check := newCheck("10")
		i := 0

		check.Query("10").Int(&i).Gt(10)
		So(check.Err(), ShouldBeError)
		So(i, ShouldEqual, 10)
	})

	Convey("param=nan", t, func() {
		check := newCheck("nan")
		i := 0

		check.Query("nan").Int(&i)
		So(check.Err(), ShouldBeError)
		So(i, ShouldEqual, 0)
	})

	Convey("overflow ints", t, func() {
		check := newCheck("128")
		i8 := int8(0)

		check.Query("128").Int8(&i8)
		So(check.Err(), ShouldBeError)
		So(i8, ShouldEqual, 0)

		check = newCheck("32768")
		i16 := int16(0)

		check.Query("32768").Int16(&i16)
		So(check.Err(), ShouldBeError)
		So(i16, ShouldEqual, 0)

		check = newCheck("2147483648")
		i32 := int32(0)

		check.Query("2147483648").Int32(&i32)
		So(check.Err(), ShouldBeError)
		So(i32, ShouldEqual, 0)

		check = newCheck("9223372036854775808")
		i64 := int64(0)

		check.Query("9223372036854775808").Int64(&i64)
		So(check.Err(), ShouldBeError)
		So(i64, ShouldEqual, 0)
	})
}
