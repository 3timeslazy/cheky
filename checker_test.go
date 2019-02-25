package cheky_test

import (
	"testing"

	"github.com/3timeslazy/cheky"

	. "github.com/smartystreets/goconvey/convey"
)

func TestChecker(t *testing.T) {
	Convey("get values from path and query", t, func() {
		var ctx = NewContext("query", "path")
		var check = cheky.Ctx(ctx)
		var query, path string

		check.Path("path").Str(&path).OneOf("path").Gt(4)
		check.Query("query").Str(&query).OneOf("query").Ge(4)

		So(check.Err(), ShouldBeError)
		So(path, ShouldEqual, "path")
		So(query, ShouldEqual, "query")
	})

	Convey("validate values", t, func() {
		var check = cheky.WithoutCtx()
		var i int
		var i8 int8

		check.Value("123").Int(&i).Gt(122)

		So(check.Err(), ShouldBeNil)
		So(i, ShouldEqual, 123)

		check.Value("1000").Int8(&i8)

		So(check.Err(), ShouldBeError)
		So(i8, ShouldEqual, 0)
	})
}
