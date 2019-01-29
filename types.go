package cheky

const (
	msgCantParse = "can't parse '%s' into %T: %v"
)

// TypesChecker implements type step.
type TypesChecker struct {
	name  string
	where string

	src string

	storeParam func(Param)
}

// Str returns String checker.
func (types TypesChecker) Str(dst *string) *String {
	s := &String{
		name:  types.name,
		where: types.where,

		val: types.src,
		src: dst,

		checks: []check{},
	}

	types.storeParam(s)
	return s
}
