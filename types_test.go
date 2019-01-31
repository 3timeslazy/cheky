package cheky_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTypes(t *testing.T) {
	Convey("init string", t, func() {
		var s string

		check := newCheck("1")
		check.Query("1").Str(&s)
	})
}
