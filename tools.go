package cheky

import (
	"fmt"
	"time"
)

func gt(val, other interface{}) (err error) {
	err = fmt.Errorf(
		"value should be greater than '%d', got '%d'",
		other, val,
	)

	switch v := val.(type) {
	case int:
		if v > other.(int) {
			err = nil
		}
	case int8:
		if v > other.(int8) {
			err = nil
		}
	case int16:
		if v > other.(int16) {
			err = nil
		}
	case int32:
		if v > other.(int32) {
			err = nil
		}
	case int64:
		if v > other.(int64) {
			err = nil
		}
	case uint:
		if v > other.(uint) {
			err = nil
		}
	case uint8:
		if v > other.(uint8) {
			err = nil
		}
	case uint16:
		if v > other.(uint16) {
			err = nil
		}
	case uint32:
		if v > other.(uint32) {
			err = nil
		}
	case uint64:
		if v > other.(uint64) {
			err = nil
		}
	case float32:
		if v > other.(float32) {
			err = nil
		}
	case float64:
		if v > other.(float64) {
			err = nil
		}
	case time.Duration:
		if v > other.(time.Duration) {
			err = nil
		}
	}

	return err
}

func ge(val, other interface{}) (err error) {
	err = fmt.Errorf(
		"value should be greater or equal to '%d', got '%d'",
		other, val,
	)

	switch v := val.(type) {
	case int:
		if v >= other.(int) {
			err = nil
		}
	case int8:
		if v >= other.(int8) {
			err = nil
		}
	case int16:
		if v >= other.(int16) {
			err = nil
		}
	case int32:
		if v >= other.(int32) {
			err = nil
		}
	case int64:
		if v >= other.(int64) {
			err = nil
		}
	case uint:
		if v >= other.(uint) {
			err = nil
		}
	case uint8:
		if v >= other.(uint8) {
			err = nil
		}
	case uint16:
		if v >= other.(uint16) {
			err = nil
		}
	case uint32:
		if v >= other.(uint32) {
			err = nil
		}
	case uint64:
		if v >= other.(uint64) {
			err = nil
		}
	case float32:
		if v >= other.(float32) {
			err = nil
		}
	case float64:
		if v >= other.(float64) {
			err = nil
		}
	case time.Duration:
		if v >= other.(time.Duration) {
			err = nil
		}
	}

	return err
}

func lt(val, other interface{}) (err error) {
	err = fmt.Errorf(
		"value should be less than '%d', got '%d'",
		other, val,
	)

	switch v := val.(type) {
	case int:
		if v < other.(int) {
			err = nil
		}
	case int8:
		if v < other.(int8) {
			err = nil
		}
	case int16:
		if v < other.(int16) {
			err = nil
		}
	case int32:
		if v < other.(int32) {
			err = nil
		}
	case int64:
		if v < other.(int64) {
			err = nil
		}
	case uint:
		if v < other.(uint) {
			err = nil
		}
	case uint8:
		if v < other.(uint8) {
			err = nil
		}
	case uint16:
		if v < other.(uint16) {
			err = nil
		}
	case uint32:
		if v < other.(uint32) {
			err = nil
		}
	case uint64:
		if v < other.(uint64) {
			err = nil
		}
	case float32:
		if v < other.(float32) {
			err = nil
		}
	case float64:
		if v < other.(float64) {
			err = nil
		}
	case time.Duration:
		if v < other.(time.Duration) {
			err = nil
		}
	}

	return
}

func le(val, other interface{}) (err error) {
	err = fmt.Errorf(
		"value should be less or equal to '%d', got '%d'",
		other, val,
	)

	switch v := val.(type) {
	case int:
		if v <= other.(int) {
			err = nil
		}
	case int8:
		if v <= other.(int8) {
			err = nil
		}
	case int16:
		if v <= other.(int16) {
			err = nil
		}
	case int32:
		if v <= other.(int32) {
			err = nil
		}
	case int64:
		if v <= other.(int64) {
			err = nil
		}
	case uint:
		if v <= other.(uint) {
			err = nil
		}
	case uint8:
		if v <= other.(uint8) {
			err = nil
		}
	case uint16:
		if v <= other.(uint16) {
			err = nil
		}
	case uint32:
		if v <= other.(uint32) {
			err = nil
		}
	case uint64:
		if v <= other.(uint64) {
			err = nil
		}
	case float32:
		if v <= other.(float32) {
			err = nil
		}
	case float64:
		if v <= other.(float64) {
			err = nil
		}
	case time.Duration:
		if v <= other.(time.Duration) {
			err = nil
		}
	}

	return
}

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
