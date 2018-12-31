package cheky

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGtGeLtLe(t *testing.T) {
	type TestCase struct {
		Less, More interface{}
		Both       interface{}
	}

	testCases := []TestCase{
		TestCase{
			Less: int(1),
			More: int(2),
			Both: []int{1, 2},
		},
		TestCase{
			Less: int8(1),
			More: int8(2),
			Both: []int8{1, 2},
		},
		TestCase{
			Less: int16(1),
			More: int16(2),
			Both: []int16{1, 2},
		},
		TestCase{
			Less: int32(1),
			More: int32(2),
			Both: []int32{1, 2},
		},
		TestCase{
			Less: int64(1),
			More: int64(2),
			Both: []int64{1, 2},
		},
		TestCase{
			Less: uint(1),
			More: uint(2),
			Both: []uint{1, 2},
		},
		TestCase{
			Less: uint8(1),
			More: uint8(2),
			Both: []uint8{1, 2},
		},
		TestCase{
			Less: uint16(1),
			More: uint16(2),
			Both: []uint16{1, 2},
		},
		TestCase{
			Less: uint32(1),
			More: uint32(2),
			Both: []uint32{1, 2},
		},
		TestCase{
			Less: uint64(1),
			More: uint64(2),
			Both: []uint64{1, 2},
		},
		TestCase{
			Less: float32(1),
			More: float32(2),
			Both: []float32{1, 2},
		},
		TestCase{
			Less: float64(1),
			More: float64(2),
			Both: []float64{1, 2},
		},
		TestCase{
			Less: time.Duration(1),
			More: time.Duration(2),
			Both: []time.Duration{
				time.Duration(1),
				time.Duration(2),
			},
		},
	}

	for _, tc := range testCases {
		Convey(
			fmt.Sprintf("gt, ge, lt, le for %T", tc.Less), t,
			func() {
				So(gt(tc.Less, tc.More), ShouldBeError)
				So(gt(tc.More, tc.Less), ShouldBeNil)

				So(lt(tc.Less, tc.More), ShouldBeNil)
				So(lt(tc.More, tc.Less), ShouldBeError)

				So(le(tc.More, tc.Less), ShouldBeError)
				So(ge(tc.Less, tc.More), ShouldBeError)

				So(le(tc.Less, tc.Less), ShouldBeNil)
				So(ge(tc.Less, tc.Less), ShouldBeNil)
			},
		)

		var err error
		var less = reflect.ValueOf(tc.Both).Slice(0, 1).Interface()
		var more = reflect.ValueOf(tc.Both).Slice(1, 2).Interface()

		Convey(
			"enum", t,
			func() {
				// tc.Less one of [tc.Less]
				err = enum(
					tc.Less,
					less,
					true,
				)
				So(err, ShouldBeNil)

				// tc.Less no one of [tc.Less]
				err = enum(
					tc.Less,
					less,
					false,
				)
				So(err, ShouldBeError)

				// tc.Less one of [tc.More]
				err = enum(
					tc.Less,
					more,
					true,
				)
				So(err, ShouldBeError)

				// tc.Less no one of [tc.More]
				err = enum(
					tc.Less,
					more,
					false,
				)
				So(err, ShouldBeNil)
			},
		)
	}
}
