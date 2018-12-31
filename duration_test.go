package cheky_test

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDuration(t *testing.T) {
	Convey("param=1s", t, func() {
		check := newCheck("1s")
		d := time.Duration(0)

		check.Query("1s").Duration(&d)
		So(check.Err(), ShouldBeNil)
		So(d, ShouldEqual, time.Second)
	})

	Convey("param=-1s", t, func() {
		check := newCheck("-1s")
		d := time.Duration(0)

		check.Query("-1s").Duration(&d)
		So(check.Err(), ShouldBeNil)
		So(d, ShouldEqual, -1*time.Second)
	})

	Convey("should be greater than 10", t, func() {
		check := newCheck("11s")
		d := time.Duration(0)

		check.Query("11s").Duration(&d).Gt(10 * time.Second)
		So(check.Err(), ShouldBeNil)
		So(d, ShouldEqual, 11*time.Second)
	})

	Convey("should be greater or equal to 10", t, func() {
		check := newCheck("10ns")
		d := time.Duration(0)

		check.Query("10ns").Duration(&d).Ge(10)
		So(check.Err(), ShouldBeNil)
		So(d, ShouldEqual, 10*time.Nanosecond)
	})

	Convey("should be less than 10", t, func() {
		check := newCheck("8m")
		d := time.Duration(0)

		check.Query("8m").Duration(&d).Lt(10 * time.Minute)
		So(check.Err(), ShouldBeNil)
		So(d, ShouldEqual, 8*time.Minute)
	})

	Convey("should be less or equal to 10", t, func() {
		check := newCheck("10s")
		d := time.Duration(0)

		check.Query("10s").Duration(&d).Le(10 * time.Second)
		So(check.Err(), ShouldBeNil)
		So(d, ShouldEqual, 10*time.Second)
	})

	Convey("should be greater than 10, but got less", t, func() {
		check := newCheck("6s")
		d := time.Duration(0)

		check.Query("6s").Duration(&d).Gt(10 * time.Second)
		So(check.Err(), ShouldBeError)
		So(d, ShouldEqual, 6*time.Second)
	})

	Convey("should be greater than 10, but got 10", t, func() {
		check := newCheck("10s")
		d := time.Duration(0)

		check.Query("10s").Duration(&d).Gt(10 * time.Second)
		So(check.Err(), ShouldBeError)
		So(d, ShouldEqual, 10*time.Second)
	})

	Convey("param=nad", t, func() {
		check := newCheck("nad")
		d := time.Duration(0)

		check.Query("nad").Duration(&d)
		So(check.Err(), ShouldBeError)
		So(d, ShouldEqual, time.Duration(0))
	})

	Convey("should be 1, 2 or 3", t, func() {
		check := newCheck("3s")
		d := time.Duration(0)

		check.Query("3s").Duration(&d).OneOf(
			1*time.Second,
			2*time.Second,
			3*time.Second,
		)
		So(check.Err(), ShouldBeNil)
		So(d, ShouldEqual, 3*time.Second)
	})

	Convey("when not one of 1, 2 or 3", t, func() {
		check := newCheck("4s")
		d := time.Duration(0)

		check.Query("4s").Duration(&d).OneOf(
			1*time.Second,
			2*time.Second,
			3*time.Second,
		)
		So(check.Err(), ShouldBeError)
		So(d, ShouldEqual, 4*time.Second)
	})

	Convey("should not be 1, 2 or 3", t, func() {
		check := newCheck("4s")
		d := time.Duration(0)

		check.Query("4s").Duration(&d).NoOne(
			1*time.Second,
			2*time.Second,
			3*time.Second,
		)
		So(check.Err(), ShouldBeNil)
		So(d, ShouldEqual, 4*time.Second)
	})

	Convey("when one of 1, 2 or 3", t, func() {
		check := newCheck("3s")
		d := time.Duration(0)

		check.Query("3s").Duration(&d).NoOne(
			1*time.Second,
			2*time.Second,
			3*time.Second,
		)
		So(check.Err(), ShouldBeError)
		So(d, ShouldEqual, 3*time.Second)
	})

	Convey("all parsed", t, func() {
		check := newCheck("1ns")
		d := time.Duration(0)

		check.Query("1ns").Duration(&d)
		So(check.Err(), ShouldBeNil)
		So(d, ShouldEqual, time.Nanosecond)

		check = newCheck("1us")

		check.Query("1us").Duration(&d)
		So(check.Err(), ShouldBeNil)
		So(d, ShouldEqual, time.Microsecond)

		check = newCheck("1µs")

		check.Query("1µs").Duration(&d)
		So(check.Err(), ShouldBeNil)
		So(d, ShouldEqual, time.Microsecond)

		check = newCheck("1ms")

		check.Query("1ms").Duration(&d)
		So(check.Err(), ShouldBeNil)
		So(d, ShouldEqual, time.Millisecond)

		check = newCheck("1s")

		check.Query("1s").Duration(&d)
		So(check.Err(), ShouldBeNil)
		So(d, ShouldEqual, time.Second)

		check = newCheck("1m")

		check.Query("1m").Duration(&d)
		So(check.Err(), ShouldBeNil)
		So(d, ShouldEqual, time.Minute)

		check = newCheck("1h")

		check.Query("1h").Duration(&d)
		So(check.Err(), ShouldBeNil)
		So(d, ShouldEqual, time.Hour)

		check = newCheck("2h45m")

		check.Query("2h45m").Duration(&d)
		So(check.Err(), ShouldBeNil)
		So(d, ShouldEqual, 2*time.Hour+45*time.Minute)
	})
}
