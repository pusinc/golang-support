package merge

import "reflect"

func Slice(typ reflect.Type, data ...reflect.Value) reflect.Value {
	result := reflect.MakeSlice(typ, 0, 0)
	for _, datum := range data {
		for i := 0; i < datum.Len(); i++ {
			result = reflect.Append(result, datum.Index(i))
		}
	}
	return result
}

func Interface(typ reflect.Type, data ...interface{}) reflect.Value {
	result := reflect.MakeSlice(typ, 0, 0)
	for _, datum := range data {
		v := reflect.ValueOf(datum)
		if v.Elem().Type() != typ {
			break
		}
		for i := 0; i < v.Len(); i++ {
			result = reflect.Append(result, v.Index(i))
		}
	}
	return result
}

func String(data ...[]string) []string {
	result := Interface(reflect.TypeOf(""), data)
	return result.Interface().([]string)
}

func Int(data ...[]int) []int {
	result := Interface(reflect.TypeOf(0), data)
	return result.Interface().([]int)
}

func Int8(data ...[]int8) []int8 {
	result := Interface(reflect.TypeOf(int8(0)), data)
	return result.Interface().([]int8)
}

func Int16(data ...[]int16) []int16 {
	result := Interface(reflect.TypeOf(int16(0)), data)
	return result.Interface().([]int16)
}

func Int32(data ...[]int32) []int32 {
	result := Interface(reflect.TypeOf(int32(0)), data)
	return result.Interface().([]int32)
}

func Int64(data ...[]int64) []int64 {
	result := Interface(reflect.TypeOf(int64(0)), data)
	return result.Interface().([]int64)
}

func Uint(data ...[]uint) []uint {
	result := Interface(reflect.TypeOf(uint(0)), data)
	return result.Interface().([]uint)
}

func Uint8(data ...[]uint8) []uint8 {
	result := Interface(reflect.TypeOf(uint8(0)), data)
	return result.Interface().([]uint8)
}

func Uint16(data ...[]uint16) []uint16 {
	result := Interface(reflect.TypeOf(uint16(0)), data)
	return result.Interface().([]uint16)
}

func Uint32(data ...[]uint32) []uint32 {
	result := Interface(reflect.TypeOf(uint32(0)), data)
	return result.Interface().([]uint32)
}

func Uint64(data ...[]uint64) []uint64 {
	result := Interface(reflect.TypeOf(uint64(0)), data)
	return result.Interface().([]uint64)
}

func Float32(data ...[]float32) []float32 {
	result := Interface(reflect.TypeOf(float32(0)), data)
	return result.Interface().([]float32)
}

func Float64(data ...[]float64) []float64 {
	result := Interface(reflect.TypeOf(float64(0)), data)
	return result.Interface().([]float64)
}
