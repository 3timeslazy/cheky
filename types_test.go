package cheky_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTypes(t *testing.T) {
	Convey("init int", t, func() {
		var i int
		var i8 int8
		var i16 int16
		var i32 int32
		var i64 int64

		check := newCheck("1")
		check.Query("1").Int(&i)
		check.Query("1").Int8(&i8)
		check.Query("1").Int16(&i16)
		check.Query("1").Int32(&i32)
		check.Query("1").Int64(&i64)
	})

	Convey("init uint", t, func() {
		var u uint
		var u8 uint8
		var u16 uint16
		var u32 uint32
		var u64 uint64

		check := newCheck("1")
		check.Query("1").Uint(&u)
		check.Query("1").Uint8(&u8)
		check.Query("1").Uint16(&u16)
		check.Query("1").Uint32(&u32)
		check.Query("1").Uint64(&u64)
	})

	Convey("init string", t, func() {
		var s string

		check := newCheck("1")
		check.Query("1").Str(&s)
	})
}
