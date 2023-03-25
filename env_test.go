package env_test

import (
	"github.com/ginger-go/env"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Env", func() {

	It("should can get value from environment", func() {
		v1 := env.TestGetEnv("test", "default")
		Expect(v1).To(Equal("test"))

		v2 := env.TestGetEnv("12", 12)
		Expect(v2).To(Equal(12))

		v3 := env.TestGetEnv("12", int64(12))
		Expect(v3).To(Equal(int64(12)))

		v4 := env.TestGetEnv("12", int32(12))
		Expect(v4).To(Equal(int32(12)))

		v5 := env.TestGetEnv("12", int16(12))
		Expect(v5).To(Equal(int16(12)))

		v6 := env.TestGetEnv("12", int8(12))
		Expect(v6).To(Equal(int8(12)))

		v7 := env.TestGetEnv("12", uint(12))
		Expect(v7).To(Equal(uint(12)))

		v8 := env.TestGetEnv("12", uint64(12))
		Expect(v8).To(Equal(uint64(12)))

		v9 := env.TestGetEnv("12", uint32(12))
		Expect(v9).To(Equal(uint32(12)))

		v10 := env.TestGetEnv("12", uint16(12))
		Expect(v10).To(Equal(uint16(12)))

		v11 := env.TestGetEnv("12", uint8(12))
		Expect(v11).To(Equal(uint8(12)))

		v12 := env.TestGetEnv("12", float64(12))
		Expect(v12).To(Equal(float64(12)))

		v13 := env.TestGetEnv("12", float32(12))
		Expect(v13).To(Equal(float32(12)))

		v14 := env.TestGetEnv("true", true)
		Expect(v14).To(Equal(true))

		v15 := env.TestGetEnv("false", false)
		Expect(v15).To(Equal(false))
	})

	It("should can get value from environment with default value", func() {
		v1 := env.GetEnv("", "default")
		Expect(v1).To(Equal("default"))

		v2 := env.GetEnv("", 12)
		Expect(v2).To(Equal(12))

		v3 := env.GetEnv("", int64(12))
		Expect(v3).To(Equal(int64(12)))

		v4 := env.GetEnv("", int32(12))
		Expect(v4).To(Equal(int32(12)))

		v5 := env.GetEnv("", int16(12))
		Expect(v5).To(Equal(int16(12)))

		v6 := env.GetEnv("", int8(12))
		Expect(v6).To(Equal(int8(12)))

		v7 := env.GetEnv("", uint(12))
		Expect(v7).To(Equal(uint(12)))

		v8 := env.GetEnv("", uint64(12))
		Expect(v8).To(Equal(uint64(12)))

		v9 := env.GetEnv("", uint32(12))
		Expect(v9).To(Equal(uint32(12)))

		v10 := env.GetEnv("", uint16(12))
		Expect(v10).To(Equal(uint16(12)))

		v11 := env.GetEnv("", uint8(12))
		Expect(v11).To(Equal(uint8(12)))

		v12 := env.GetEnv("", float64(12))
		Expect(v12).To(Equal(float64(12)))

		v13 := env.GetEnv("", float32(12))
		Expect(v13).To(Equal(float32(12)))

		v14 := env.GetEnv("", true)
		Expect(v14).To(Equal(true))

		v15 := env.GetEnv("", false)
		Expect(v15).To(Equal(false))
	})

	It("should can get list of value from environment", func() {
		v1 := env.TestGetEnvList("v1,v2,v3", []string{"default1", "default2", "default3"})
		Expect(v1).To(Equal([]string{"v1", "v2", "v3"}))

		v2 := env.TestGetEnvList("1,2,3", []int{1, 2, 3})
		Expect(v2).To(Equal([]int{1, 2, 3}))

		v3 := env.TestGetEnvList("1,2,3", []int64{1, 2, 3})
		Expect(v3).To(Equal([]int64{1, 2, 3}))

		v4 := env.TestGetEnvList("1,2,3", []int32{1, 2, 3})
		Expect(v4).To(Equal([]int32{1, 2, 3}))

		v5 := env.TestGetEnvList("1,2,3", []int16{1, 2, 3})
		Expect(v5).To(Equal([]int16{1, 2, 3}))

		v6 := env.TestGetEnvList("1,2,3", []int8{1, 2, 3})
		Expect(v6).To(Equal([]int8{1, 2, 3}))

		v7 := env.TestGetEnvList("1,2,3", []uint{1, 2, 3})
		Expect(v7).To(Equal([]uint{1, 2, 3}))

		v8 := env.TestGetEnvList("1,2,3", []uint64{1, 2, 3})
		Expect(v8).To(Equal([]uint64{1, 2, 3}))

		v9 := env.TestGetEnvList("1,2,3", []uint32{1, 2, 3})
		Expect(v9).To(Equal([]uint32{1, 2, 3}))

		v10 := env.TestGetEnvList("1,2,3", []uint16{1, 2, 3})
		Expect(v10).To(Equal([]uint16{1, 2, 3}))

		v11 := env.TestGetEnvList("1,2,3", []uint8{1, 2, 3})
		Expect(v11).To(Equal([]uint8{1, 2, 3}))

		v12 := env.TestGetEnvList("1,2,3", []float64{1, 2, 3})
		Expect(v12).To(Equal([]float64{1, 2, 3}))

		v13 := env.TestGetEnvList("1,2,3", []float32{1, 2, 3})
		Expect(v13).To(Equal([]float32{1, 2, 3}))

		v14 := env.TestGetEnvList("true,false,true", []bool{true, false, true})
		Expect(v14).To(Equal([]bool{true, false, true}))

		v15 := env.TestGetEnvList("true,false,true", []bool{true, false, true})
		Expect(v15).To(Equal([]bool{true, false, true}))
	})

	It("should can get list of value from environment with default value", func() {
		v1 := env.GetEnvList("", []string{"default1", "default2", "default3"})
		Expect(v1).To(Equal([]string{"default1", "default2", "default3"}))

		v2 := env.GetEnvList("", []int{1, 2, 3})
		Expect(v2).To(Equal([]int{1, 2, 3}))

		v3 := env.GetEnvList("", []int64{1, 2, 3})
		Expect(v3).To(Equal([]int64{1, 2, 3}))

		v4 := env.GetEnvList("", []int32{1, 2, 3})
		Expect(v4).To(Equal([]int32{1, 2, 3}))

		v5 := env.GetEnvList("", []int16{1, 2, 3})
		Expect(v5).To(Equal([]int16{1, 2, 3}))

		v6 := env.GetEnvList("", []int8{1, 2, 3})
		Expect(v6).To(Equal([]int8{1, 2, 3}))

		v7 := env.GetEnvList("", []uint{1, 2, 3})
		Expect(v7).To(Equal([]uint{1, 2, 3}))

		v8 := env.GetEnvList("", []uint64{1, 2, 3})
		Expect(v8).To(Equal([]uint64{1, 2, 3}))

		v9 := env.GetEnvList("", []uint32{1, 2, 3})
		Expect(v9).To(Equal([]uint32{1, 2, 3}))

		v10 := env.GetEnvList("", []uint16{1, 2, 3})
		Expect(v10).To(Equal([]uint16{1, 2, 3}))

		v11 := env.GetEnvList("", []uint8{1, 2, 3})
		Expect(v11).To(Equal([]uint8{1, 2, 3}))

		v12 := env.GetEnvList("", []float64{1, 2, 3})
		Expect(v12).To(Equal([]float64{1, 2, 3}))

		v13 := env.GetEnvList("", []float32{1, 2, 3})
		Expect(v13).To(Equal([]float32{1, 2, 3}))

		v14 := env.GetEnvList("", []bool{true, false, true})
		Expect(v14).To(Equal([]bool{true, false, true}))

		v15 := env.GetEnvList("", []bool{true, false, true})
		Expect(v15).To(Equal([]bool{true, false, true}))
	})

	It("should return default value if unable to parse value", func() {
		v1 := env.TestGetEnv("invalid", 1)
		Expect(v1).To(Equal(1))

		v2 := env.TestGetEnv("invalid", int64(1))
		Expect(v2).To(Equal(int64(1)))

		v3 := env.TestGetEnv("invalid", int32(1))
		Expect(v3).To(Equal(int32(1)))

		v4 := env.TestGetEnv("invalid", int16(1))
		Expect(v4).To(Equal(int16(1)))

		v5 := env.TestGetEnv("invalid", int8(1))
		Expect(v5).To(Equal(int8(1)))

		v6 := env.TestGetEnv("invalid", uint(1))
		Expect(v6).To(Equal(uint(1)))

		v7 := env.TestGetEnv("invalid", uint64(1))
		Expect(v7).To(Equal(uint64(1)))

		v8 := env.TestGetEnv("invalid", uint32(1))
		Expect(v8).To(Equal(uint32(1)))

		v9 := env.TestGetEnv("invalid", uint16(1))
		Expect(v9).To(Equal(uint16(1)))

		v10 := env.TestGetEnv("invalid", uint8(1))
		Expect(v10).To(Equal(uint8(1)))

		v11 := env.TestGetEnv("invalid", float64(1))
		Expect(v11).To(Equal(float64(1)))

		v12 := env.TestGetEnv("invalid", float32(1))
		Expect(v12).To(Equal(float32(1)))

		v13 := env.TestGetEnv("invalid", true)
		Expect(v13).To(Equal(true))

		v14 := env.TestGetEnv("invalid", false)
		Expect(v14).To(Equal(false))
	})

})
