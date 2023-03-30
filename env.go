package env

import (
	"os"
	"strconv"
	"strings"
)

// String returns the value of the environment variable named by the key.
func String(key string, defaultValue ...string) string {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return ""
	}
	return value
}

// Strings returns the list of value of the environment variable named by the key.
func Strings(key string, defaultValue ...[]string) []string {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return []string{}
	}
	return strings.Split(value, ",")
}

// Int returns the value of the environment variable named by the key.
func Int(key string, defaultValue ...int) int {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	output, err := strconv.Atoi(value)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return output
}

// Ints returns the list of value of the environment variable named by the key.
func Ints(key string, defaultValue ...[]int) []int {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return []int{}
	}
	li := strings.Split(value, ",")
	output := make([]int, 0)

	for _, v := range li {
		o, err := strconv.Atoi(v)
		if err != nil {
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}
			return []int{}
		}
		output = append(output, o)
	}

	return output
}

// Int64 returns the value of the environment variable named by the key.
func Int64(key string, defaultValue ...int64) int64 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	output, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return output
}

// Int64s returns the list of value of the environment variable named by the key.
func Int64s(key string, defaultValue ...[]int64) []int64 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return []int64{}
	}
	li := strings.Split(value, ",")
	output := make([]int64, 0)

	for _, v := range li {
		o, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}
			return []int64{}
		}
		output = append(output, o)
	}

	return output
}

// Int32 returns the value of the environment variable named by the key.
func Int32(key string, defaultValue ...int32) int32 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	output, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return int32(output)
}

// Int32s returns the list of value of the environment variable named by the key.
func Int32s(key string, defaultValue ...[]int32) []int32 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return []int32{}
	}
	li := strings.Split(value, ",")
	output := make([]int32, 0)

	for _, v := range li {
		o, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}
			return []int32{}
		}
		output = append(output, int32(o))
	}

	return output
}

// Int16 returns the value of the environment variable named by the key.
func Int16(key string, defaultValue ...int16) int16 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	output, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return int16(output)
}

// Int16s returns the list of value of the environment variable named by the key.
func Int16s(key string, defaultValue ...[]int16) []int16 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return []int16{}
	}
	li := strings.Split(value, ",")
	output := make([]int16, 0)

	for _, v := range li {
		o, err := strconv.ParseInt(v, 10, 16)
		if err != nil {
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}
			return []int16{}
		}
		output = append(output, int16(o))
	}

	return output
}

// Int8 returns the value of the environment variable named by the key.
func Int8(key string, defaultValue ...int8) int8 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	output, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return int8(output)
}

// Int8s returns the list of value of the environment variable named by the key.
func Int8s(key string, defaultValue ...[]int8) []int8 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return []int8{}
	}
	li := strings.Split(value, ",")
	output := make([]int8, 0)

	for _, v := range li {
		o, err := strconv.ParseInt(v, 10, 8)
		if err != nil {
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}
			return []int8{}
		}
		output = append(output, int8(o))
	}

	return output
}

// Float64 returns the value of the environment variable named by the key.
func Float64(key string, defaultValue ...float64) float64 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	output, err := strconv.ParseFloat(value, 64)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return output
}

// Float64s returns the list of value of the environment variable named by the key.
func Float64s(key string, defaultValue ...[]float64) []float64 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return []float64{}
	}
	li := strings.Split(value, ",")
	output := make([]float64, 0)

	for _, v := range li {
		o, err := strconv.ParseFloat(v, 64)
		if err != nil {
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}
			return []float64{}
		}
		output = append(output, o)
	}

	return output
}

// Float32 returns the value of the environment variable named by the key.
func Float32(key string, defaultValue ...float32) float32 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	output, err := strconv.ParseFloat(value, 32)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return float32(output)
}

// Float32s returns the list of value of the environment variable named by the key.
func Float32s(key string, defaultValue ...[]float32) []float32 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return []float32{}
	}
	li := strings.Split(value, ",")
	output := make([]float32, 0)

	for _, v := range li {
		o, err := strconv.ParseFloat(v, 32)
		if err != nil {
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}
			return []float32{}
		}
		output = append(output, float32(o))
	}

	return output
}

// Uint returns the value of the environment variable named by the key.
func Uint(key string, defaultValue ...uint) uint {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	output, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return uint(output)
}

// Uints returns the list of value of the environment variable named by the key.
func Uints(key string, defaultValue ...[]uint) []uint {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return []uint{}
	}
	li := strings.Split(value, ",")
	output := make([]uint, 0)

	for _, v := range li {
		o, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}
			return []uint{}
		}
		output = append(output, uint(o))
	}

	return output
}

// Uint64 returns the value of the environment variable named by the key.
func Uint64(key string, defaultValue ...uint64) uint64 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	output, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return output
}

// Uint64s returns the list of value of the environment variable named by the key.
func Uint64s(key string, defaultValue ...[]uint64) []uint64 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return []uint64{}
	}
	li := strings.Split(value, ",")
	output := make([]uint64, 0)

	for _, v := range li {
		o, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}
			return []uint64{}
		}
		output = append(output, o)
	}

	return output
}

// Uint32 returns the value of the environment variable named by the key.
func Uint32(key string, defaultValue ...uint32) uint32 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	output, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return uint32(output)
}

// Uint32s returns the list of value of the environment variable named by the key.
func Uint32s(key string, defaultValue ...[]uint32) []uint32 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return []uint32{}
	}
	li := strings.Split(value, ",")
	output := make([]uint32, 0)

	for _, v := range li {
		o, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}
			return []uint32{}
		}
		output = append(output, uint32(o))
	}

	return output
}

// Uint16 returns the value of the environment variable named by the key.
func Uint16(key string, defaultValue ...uint16) uint16 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	output, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	return uint16(output)
}

// Uint16s returns the list of value of the environment variable named by the key.
func Uint16s(key string, defaultValue ...[]uint16) []uint16 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return []uint16{}
	}
	li := strings.Split(value, ",")
	output := make([]uint16, 0)

	for _, v := range li {
		o, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}
			return []uint16{}
		}
		output = append(output, uint16(o))
	}

	return output
}

// Uint8 returns the value of the environment variable named by the key.
func Uint8(key string, defaultValue ...uint8) uint8 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	output, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	return uint8(output)
}

// Uint8s returns the list of value of the environment variable named by the key.
func Uint8s(key string, defaultValue ...[]uint8) []uint8 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return []uint8{}
	}
	li := strings.Split(value, ",")
	output := make([]uint8, 0)

	for _, v := range li {
		o, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}
			return []uint8{}
		}
		output = append(output, uint8(o))
	}

	return output
}

// Bool returns the value of the environment variable named by the key.
func Bool(key string, defaultValue ...bool) bool {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return false
	}

	output, err := strconv.ParseBool(value)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return false
	}

	return output
}

// Bools returns the list of value of the environment variable named by the key.
func Bools(key string, defaultValue ...[]bool) []bool {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return []bool{}
	}
	li := strings.Split(value, ",")
	output := make([]bool, 0)

	for _, v := range li {
		o, err := strconv.ParseBool(v)
		if err != nil {
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}
			return []bool{}
		}
		output = append(output, o)
	}

	return output
}

// Complex128 returns the value of the environment variable named by the key.
func Complex128(key string, defaultValue ...complex128) complex128 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	output, err := strconv.ParseComplex(value, 128)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	return output
}

// Complex128s returns the list of value of the environment variable named by the key.
func Complex128s(key string, defaultValue ...[]complex128) []complex128 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return []complex128{}
	}
	li := strings.Split(value, ",")
	output := make([]complex128, 0)

	for _, v := range li {
		o, err := strconv.ParseComplex(v, 128)
		if err != nil {
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}
			return []complex128{}
		}
		output = append(output, o)
	}

	return output
}

// Complex64 returns the value of the environment variable named by the key.
func Complex64(key string, defaultValue ...complex64) complex64 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	output, err := strconv.ParseComplex(value, 64)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	return complex64(output)
}

// Complex64s returns the list of value of the environment variable named by the key.
func Complex64s(key string, defaultValue ...[]complex64) []complex64 {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return []complex64{}
	}

	li := strings.Split(value, ",")
	output := make([]complex64, 0)

	for _, v := range li {
		o, err := strconv.ParseComplex(v, 64)
		if err != nil {
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}
			return []complex64{}
		}
		output = append(output, complex64(o))
	}

	return output
}

// Byte returns the value of the environment variable named by the key.
func Byte(key string, defaultValue ...byte) byte {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	output, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	return byte(output)
}

// Bytes returns the list of value of the environment variable named by the key.
func Bytes(key string, defaultValue ...[]byte) []byte {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return []byte{}
	}

	li := strings.Split(value, ",")
	output := make([]byte, 0)

	for _, v := range li {
		o, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}
			return []byte{}
		}
		output = append(output, byte(o))
	}

	return output
}

// Rune returns the value of the environment variable named by the key.
func Rune(key string, defaultValue ...rune) rune {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	output, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	return rune(output)
}

// Runes returns the list of value of the environment variable named by the key.
func Runes(key string, defaultValue ...[]rune) []rune {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return []rune{}
	}

	li := strings.Split(value, ",")
	output := make([]rune, 0)

	for _, v := range li {
		o, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}
			return []rune{}
		}
		output = append(output, rune(o))
	}

	return output
}
