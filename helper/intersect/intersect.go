package intersect

import "reflect"

func Interface(x interface{}, y interface{}) reflect.Value {
	xValue, yValue := reflect.Value{}, reflect.Value{}
	if xValueNode, ok := x.(reflect.Value); ok {
		xValue = xValueNode
	} else {
		xValue = reflect.ValueOf(x)
	}
	if yValueNode, ok := y.(reflect.Value); ok {
		yValue = yValueNode
	} else {
		yValue = reflect.ValueOf(y)
	}
	if xValue.Type().Kind() == reflect.Ptr {
		xValue = xValue.Elem()
	}
	if yValue.Type().Kind() == reflect.Ptr {
		yValue = yValue.Elem()
	}

	set := reflect.MakeSlice(xValue.Type(), 0, 0)
	hash := make(map[interface{}]bool)

	for i := 0; i < xValue.Len(); i++ {
		v := xValue.Index(i)
		hash[v.Interface()] = true
	}
	for i := 0; i < yValue.Len(); i++ {
		v := yValue.Index(i)
		if _, ok := hash[v.Interface()]; ok {
			set = reflect.Append(set, v)
		}
	}
	return set
}

func String(x []string, y []string) []string {
	value := Interface(x, y)
	return value.Interface().([]string)
}

func Int(x []int, y []int) []int {
	value := Interface(x, y)
	return value.Interface().([]int)
}

func Int8(x []int8, y []int8) []int8 {
	value := Interface(x, y)
	return value.Interface().([]int8)
}

func Int16(x []int16, y []int16) []int16 {
	value := Interface(x, y)
	return value.Interface().([]int16)
}

func Int32(x []int32, y []int32) []int32 {
	value := Interface(x, y)
	return value.Interface().([]int32)
}

func Int64(x []int64, y []int64) []int64 {
	value := Interface(x, y)
	return value.Interface().([]int64)
}

func Uint(x []uint, y []uint) []uint {
	value := Interface(x, y)
	return value.Interface().([]uint)
}

func Uint8(x []uint8, y []uint8) []uint8 {
	value := Interface(x, y)
	return value.Interface().([]uint8)
}

func Uint16(x []uint16, y []uint16) []uint16 {
	value := Interface(x, y)
	return value.Interface().([]uint16)
}

func Uint32(x []uint32, y []uint32) []uint32 {
	value := Interface(x, y)
	return value.Interface().([]uint32)
}

func Uint64(x []uint64, y []uint64) []uint64 {
	value := Interface(x, y)
	return value.Interface().([]uint64)
}

func Float32(x []float32, y []float32) []float32 {
	value := Interface(x, y)
	return value.Interface().([]float32)
}

func Float64(x []float64, y []float64) []float64 {
	value := Interface(x, y)
	return value.Interface().([]float64)
}
