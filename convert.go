package env

import (
	"reflect"
	"strconv"
)

func convertString[T any](value string) (T, error) {
	var output interface{}
	var err error

	v := reflect.TypeOf(new(T))
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.String:
		output = value
	case reflect.Int:
		var i int
		i, err = strconv.Atoi(value)
		if err != nil {
			output = int(0)
		} else {
			output = int(i)
		}
	case reflect.Int64:
		var i int64
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			output = int64(0)
		} else {
			output = int64(i)
		}
	case reflect.Int32:
		var i int64
		i, err = strconv.ParseInt(value, 10, 32)
		if err != nil {
			output = int32(0)
		} else {
			output = int32(i)
		}
	case reflect.Int16:
		var i int64
		i, err = strconv.ParseInt(value, 10, 16)
		if err != nil {
			output = int16(0)
		} else {
			output = int16(i)
		}
	case reflect.Int8:
		var i int64
		i, err = strconv.ParseInt(value, 10, 8)
		if err != nil {
			output = int8(0)
		} else {
			output = int8(i)
		}
	case reflect.Uint:
		var i uint64
		i, err = strconv.ParseUint(value, 10, 64)
		if err != nil {
			output = uint(0)
		} else {
			output = uint(i)
		}
	case reflect.Uint64:
		var i uint64
		i, err = strconv.ParseUint(value, 10, 64)
		if err != nil {
			output = uint64(0)
		} else {
			output = uint64(i)
		}
	case reflect.Uint32:
		var i uint64
		i, err = strconv.ParseUint(value, 10, 32)
		if err != nil {
			output = uint32(0)
		} else {
			output = uint32(i)
		}
	case reflect.Uint16:
		var i uint64
		i, err = strconv.ParseUint(value, 10, 16)
		if err != nil {
			output = uint16(0)
		} else {
			output = uint16(i)
		}
	case reflect.Uint8:
		var i uint64
		i, err = strconv.ParseUint(value, 10, 8)
		if err != nil {
			output = uint8(0)
		} else {
			output = uint8(i)
		}
	case reflect.Float64:
		var i float64
		i, err = strconv.ParseFloat(value, 64)
		if err != nil {
			output = float64(0)
		} else {
			output = float64(i)
		}
	case reflect.Float32:
		var i float64
		i, err = strconv.ParseFloat(value, 32)
		if err != nil {
			output = float32(0)
		} else {
			output = float32(i)
		}
	case reflect.Bool:
		var i bool
		i, err = strconv.ParseBool(value)
		if err != nil {
			output = bool(false)
		} else {
			output = bool(i)
		}
	default:
		panic("only support string, int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, bool")
	}
	return output.(T), err
}
