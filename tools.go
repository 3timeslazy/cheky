package cheky

import (
	"fmt"
	"time"
)

func oneOf(val, ones interface{}) error {
	return enum(val, ones, true)
}

func noOne(val, ones interface{}) error {
	return enum(val, ones, false)
}

func enum(val, ones interface{}, oneOf bool) (err error) {
	errUpd := func(isEqual bool) {
		if isEqual {
			err = fmt.Errorf(
				"value should not be one of %v, got '%v'",
				ones, val,
			)
		}
	}

	if oneOf {
		err = fmt.Errorf(
			"value should be one of %v, got '%v'",
			ones, val,
		)
		errUpd = func(isEqual bool) {
			if isEqual {
				err = nil
			}
		}
	}

	switch ones := ones.(type) {
	case []string:
		val := val.(string)
		for _, one := range ones {
			errUpd(val == one)
		}
	case []int:
		val := val.(int)
		for _, one := range ones {
			errUpd(val == one)
		}
	case []int8:
		val := val.(int8)
		for _, one := range ones {
			errUpd(val == one)
		}
	case []int16:
		val := val.(int16)
		for _, one := range ones {
			errUpd(val == one)
		}
	case []int32:
		val := val.(int32)
		for _, one := range ones {
			errUpd(val == one)
		}
	case []int64:
		val := val.(int64)
		for _, one := range ones {
			errUpd(val == one)
		}
	case []uint:
		val := val.(uint)
		for _, one := range ones {
			errUpd(val == one)
		}
	case []uint8:
		val := val.(uint8)
		for _, one := range ones {
			errUpd(val == one)
		}
	case []uint16:
		val := val.(uint16)
		for _, one := range ones {
			errUpd(val == one)
		}
	case []uint32:
		val := val.(uint32)
		for _, one := range ones {
			errUpd(val == one)
		}
	case []uint64:
		val := val.(uint64)
		for _, one := range ones {
			errUpd(val == one)
		}
	case []float32:
		val := val.(float32)
		for _, one := range ones {
			errUpd(val == one)
		}
	case []float64:
		val := val.(float64)
		for _, one := range ones {
			errUpd(val == one)
		}
	case []time.Duration:
		val := val.(time.Duration)
		for _, one := range ones {
			errUpd(val == one)
		}
	}

	return err
}
