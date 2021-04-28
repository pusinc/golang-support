package diff

import (
	"github.com/pusinc/golang-support/helper/contain"
	"reflect"
)

func Interface(x interface{}, y interface{}) (reflect.Value, reflect.Value) {
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
	leftSlice := reflect.MakeSlice(xValue.Type(), 0, 0)
	rightSlice := reflect.MakeSlice(xValue.Type(), 0, 0)

	for i := 0; i < xValue.Len(); i++ {
		v := xValue.Index(i)
		if contain.Interface(y, v) == false {
			leftSlice = reflect.Append(leftSlice, v)
		}
	}

	for i := 0; i < yValue.Len(); i++ {
		v := yValue.Index(i)
		if contain.Interface(x, v) == false {
			rightSlice = reflect.Append(rightSlice, v)
		}
	}

	return leftSlice, rightSlice
}

func String(x []string, y []string) ([]string, []string) {
	xResult, yResult := Interface(x, y)
	return xResult.Interface().([]string), yResult.Interface().([]string)
}

func Int(x []int, y []int) ([]int, []int) {
	xResult, yResult := Interface(x, y)
	return xResult.Interface().([]int), yResult.Interface().([]int)
}

func Int8(x []int8, y []int8) ([]int8, []int8) {
	xResult, yResult := Interface(x, y)
	return xResult.Interface().([]int8), yResult.Interface().([]int8)
}

func Int16(x []int16, y []int16) ([]int16, []int16) {
	xResult, yResult := Interface(x, y)
	return xResult.Interface().([]int16), yResult.Interface().([]int16)
}

func Int32(x []int32, y []int32) ([]int32, []int32) {
	xResult, yResult := Interface(x, y)
	return xResult.Interface().([]int32), yResult.Interface().([]int32)
}

func Int64(x []int64, y []int64) ([]int64, []int64) {
	xResult, yResult := Interface(x, y)
	return xResult.Interface().([]int64), yResult.Interface().([]int64)
}

func Uint(x []uint, y []uint) ([]uint, []uint) {
	xResult, yResult := Interface(x, y)
	return xResult.Interface().([]uint), yResult.Interface().([]uint)
}

func Uint8(x []uint8, y []uint8) ([]uint8, []uint8) {
	xResult, yResult := Interface(x, y)
	return xResult.Interface().([]uint8), yResult.Interface().([]uint8)
}

func Uint16(x []uint16, y []uint16) ([]uint16, []uint16) {
	xResult, yResult := Interface(x, y)
	return xResult.Interface().([]uint16), yResult.Interface().([]uint16)
}

func Uint32(x []uint32, y []uint32) ([]uint32, []uint32) {
	xResult, yResult := Interface(x, y)
	return xResult.Interface().([]uint32), yResult.Interface().([]uint32)
}

func Uint64(x []uint64, y []uint64) ([]uint64, []uint64) {
	xResult, yResult := Interface(x, y)
	return xResult.Interface().([]uint64), yResult.Interface().([]uint64)
}

func Float32(x []float32, y []float32) ([]float32, []float32) {
	xResult, yResult := Interface(x, y)
	return xResult.Interface().([]float32), yResult.Interface().([]float32)
}

func Float64(x []float64, y []float64) ([]float64, []float64) {
	xResult, yResult := Interface(x, y)
	return xResult.Interface().([]float64), yResult.Interface().([]float64)
}
