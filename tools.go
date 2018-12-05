package cheky

import (
	"fmt"
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
	}

	return err
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
	}

	return err
}
