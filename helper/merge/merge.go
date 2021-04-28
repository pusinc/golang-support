package merge

import "reflect"

func Interface(data ...interface{}) reflect.Value {
	if len(data) == 0 {
		return reflect.Value{}
	}
	firstRecordV, ok := data[0].(reflect.Value)
	if !ok {
		firstRecordV = reflect.ValueOf(data[0])
	}
	if firstRecordV.Type().Kind() == reflect.Ptr {
		firstRecordV = firstRecordV.Elem()
	}
	result := reflect.MakeSlice(firstRecordV.Type(), 0, 0)
	for _, datum := range data {
		v, ok := datum.(reflect.Value)
		if !ok {
			v = reflect.ValueOf(datum)
		}
		if v.Type().Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Type().Kind() != reflect.Slice {
			continue
		}
		for i := 0; i < v.Len(); i++ {
			result = reflect.Append(result, v.Index(i))
		}
	}
	return result
}

func String(data ...[]string) []string {
	result := Interface(data)
	return result.Interface().([]string)
}

func Int(data ...[]int) []int {
	result := Interface(data)
	return result.Interface().([]int)
}

func Int8(data ...[]int8) []int8 {
	result := Interface(data)
	return result.Interface().([]int8)
}

func Int16(data ...[]int16) []int16 {
	result := Interface(data)
	return result.Interface().([]int16)
}

func Int32(data ...[]int32) []int32 {
	result := Interface(data)
	return result.Interface().([]int32)
}

func Int64(data ...[]int64) []int64 {
	result := Interface(data)
	return result.Interface().([]int64)
}

func Uint(data ...[]uint) []uint {
	result := Interface(data)
	return result.Interface().([]uint)
}

func Uint8(data ...[]uint8) []uint8 {
	result := Interface(data)
	return result.Interface().([]uint8)
}

func Uint16(data ...[]uint16) []uint16 {
	result := Interface(data)
	return result.Interface().([]uint16)
}

func Uint32(data ...[]uint32) []uint32 {
	result := Interface(data)
	return result.Interface().([]uint32)
}

func Uint64(data ...[]uint64) []uint64 {
	result := Interface(data)
	return result.Interface().([]uint64)
}

func Float32(data ...[]float32) []float32 {
	result := Interface(data)
	return result.Interface().([]float32)
}

func Float64(data ...[]float64) []float64 {
	result := Interface(data)
	return result.Interface().([]float64)
}
