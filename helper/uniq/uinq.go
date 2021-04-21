package uniq

import "reflect"

func Interface(typ reflect.Type, a interface{}) reflect.Value {
	mapValue := reflect.ValueOf(true)
	aValue := reflect.ValueOf(a)
	length := aValue.Len()

	seen := reflect.MakeMapWithSize(typ, length)

	j := 0

	for i := 0; i < length; i++ {
		v := aValue.Index(i)
		testV := seen.MapIndex(v)
		if !testV.IsNil() {
			continue
		}
		seen.SetMapIndex(v, mapValue)
		aValue.Index(j).Set(v)
		j++
	}

	return aValue.Slice(0, j)
}

func String(a []string) []string {
	result := Interface(reflect.TypeOf(""), a)
	return result.Interface().([]string)
}

func Int(a []int) []int {
	result := Interface(reflect.TypeOf(int(0)), a)
	return result.Interface().([]int)
}

func Int8(a []int8) []int8 {
	result := Interface(reflect.TypeOf(int8(0)), a)
	return result.Interface().([]int8)
}

func Int16(a []int16) []int16 {
	result := Interface(reflect.TypeOf(int16(0)), a)
	return result.Interface().([]int16)
}

func Int32(a []int32) []int32 {
	result := Interface(reflect.TypeOf(int32(0)), a)
	return result.Interface().([]int32)
}

func Int64(a []int64) []int64 {
	result := Interface(reflect.TypeOf(int64(0)), a)
	return result.Interface().([]int64)
}

func Uint(a []uint) []uint {
	result := Interface(reflect.TypeOf(uint(0)), a)
	return result.Interface().([]uint)
}

func Uint8(a []uint8) []uint8 {
	result := Interface(reflect.TypeOf(uint8(0)), a)
	return result.Interface().([]uint8)
}

func Uint16(a []uint16) []uint16 {
	result := Interface(reflect.TypeOf(uint16(0)), a)
	return result.Interface().([]uint16)
}

func Uint32(a []uint32) []uint32 {
	result := Interface(reflect.TypeOf(uint32(0)), a)
	return result.Interface().([]uint32)
}

func Uint64(a []uint64) []uint64 {
	result := Interface(reflect.TypeOf(uint64(0)), a)
	return result.Interface().([]uint64)
}

func Float32(a []float32) []float32 {
	result := Interface(reflect.TypeOf(float32(0)), a)
	return result.Interface().([]float32)
}

func Float64(a []float64) []float64 {
	result := Interface(reflect.TypeOf(float64(0)), a)
	return result.Interface().([]float64)
}
