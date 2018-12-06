package cheky_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUint(t *testing.T) {
	Convey("param=10", t, func() {
		check := newCheck("10")
		u := uint(0)

		check.Query("10").Uint(&u)
		So(check.Err(), ShouldBeNil)
		So(u, ShouldEqual, 10)
	})

	Convey("cant parse negative into uint", t, func() {
		check := newCheck("-8")
		u := uint(0)

		check.Query("-8").Uint(&u)
		So(check.Err(), ShouldBeError)
		So(u, ShouldEqual, 0)
	})

	Convey("should be greater than 10", t, func() {
		check := newCheck("11")
		u := uint(0)

		check.Query("11").Uint(&u).Gt(10)
		So(check.Err(), ShouldBeNil)
		So(u, ShouldEqual, 11)
	})

	Convey("should be greater or equal to 10", t, func() {
		check := newCheck("10")
		u := uint(0)

		check.Query("10").Uint(&u).Ge(10)
		So(check.Err(), ShouldBeNil)
		So(u, ShouldEqual, 10)
	})

	Convey("should be less than 10", t, func() {
		check := newCheck("8")
		u := uint(0)

		check.Query("8").Uint(&u).Lt(10)
		So(check.Err(), ShouldBeNil)
		So(u, ShouldEqual, 8)
	})

	Convey("should be less or equal to 10", t, func() {
		check := newCheck("10")
		u := uint(0)

		check.Query("10").Uint(&u).Le(10)
		So(check.Err(), ShouldBeNil)
		So(u, ShouldEqual, 10)
	})

	Convey("should be greater than 10, but got less", t, func() {
		check := newCheck("6")
		u := uint(0)

		check.Query("6").Uint(&u).Gt(10)
		So(check.Err(), ShouldBeError)
		So(u, ShouldEqual, 6)
	})

	Convey("should be greater than 10, but got 10", t, func() {
		check := newCheck("10")
		u := uint(0)

		check.Query("10").Uint(&u).Gt(10)
		So(check.Err(), ShouldBeError)
		So(u, ShouldEqual, 10)
	})

	Convey("param=nan", t, func() {
		check := newCheck("nan")
		u := uint(0)

		check.Query("nan").Uint(&u)
		So(check.Err(), ShouldBeError)
		So(u, ShouldEqual, 0)
	})

	Convey("overflow uints", t, func() {
		check := newCheck("256")
		u8 := uint8(0)

		check.Query("256").Uint8(&u8)
		So(check.Err(), ShouldBeError)
		So(u8, ShouldEqual, 0)

		check = newCheck("65536")
		u16 := uint16(0)

		check.Query("65536").Uint16(&u16)
		So(check.Err(), ShouldBeError)
		So(u16, ShouldEqual, 0)

		check = newCheck("4294967296")
		u32 := uint32(0)

		check.Query("4294967296").Uint32(&u32)
		So(check.Err(), ShouldBeError)
		So(u32, ShouldEqual, 0)

		check = newCheck("18446744073709551616")
		u64 := uint64(0)

		check.Query("18446744073709551616").Uint64(&u64)
		So(check.Err(), ShouldBeError)
		So(u64, ShouldEqual, 0)
	})

	Convey("should be 1, 2 or 3", t, func() {
		check := newCheck("3")
		u := uint(0)

		check.Query("3").Uint(&u).OneOf(1, 2, 3)
		So(check.Err(), ShouldBeNil)
		So(u, ShouldEqual, 3)
	})

	Convey("when not one of 1, 2 or 3", t, func() {
		check := newCheck("4")
		u := uint(0)

		check.Query("4").Uint(&u).OneOf(1, 2, 3)
		So(check.Err(), ShouldBeError)
		So(u, ShouldEqual, 4)
	})

	Convey("should not be 1, 2 or 3", t, func() {
		check := newCheck("4")
		u := uint(0)

		check.Query("4").Uint(&u).NoOne(1, 2, 3)
		So(check.Err(), ShouldBeNil)
		So(u, ShouldEqual, 4)
	})

	Convey("when one of 1, 2 or 3", t, func() {
		check := newCheck("3")
		u := uint(0)

		check.Query("3").Uint(&u).NoOne(1, 2, 3)
		So(check.Err(), ShouldBeError)
		So(u, ShouldEqual, 3)
	})
}
