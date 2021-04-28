package uniq

import (
	"fmt"
	"reflect"
)

func Interface(dst interface{}) reflect.Value {
	dstV := reflect.Value{}
	if dstVNode, ok := dst.(reflect.Value); ok {
		dstV = dstVNode
	} else {
		dstV = reflect.ValueOf(dst)
	}
	if dstV.Type().Kind() == reflect.Ptr {
		dstV = dstV.Elem()
	}
	if dstV.Type().Kind() != reflect.Slice {
		panic(fmt.Errorf("dst must be kind of %s", reflect.Slice))
	}
	length := dstV.Len()
	seen := make(map[interface{}]bool)

	j := 0

	for i := 0; i < length; i++ {
		v := dstV.Index(i)
		if _, ok := seen[v.Interface()]; ok {
			continue
		}
		seen[v.Interface()] = true
		dstV.Index(j).Set(v)
		j++
	}

	return dstV.Slice(0, j)
}

func String(a []string) []string {
	result := Interface(a)
	return result.Interface().([]string)
}

func Int(a []int) []int {
	result := Interface(a)
	return result.Interface().([]int)
}

func Int8(a []int8) []int8 {
	result := Interface(a)
	return result.Interface().([]int8)
}

func Int16(a []int16) []int16 {
	result := Interface(a)
	return result.Interface().([]int16)
}

func Int32(a []int32) []int32 {
	result := Interface(a)
	return result.Interface().([]int32)
}

func Int64(a []int64) []int64 {
	result := Interface(a)
	return result.Interface().([]int64)
}

func Uint(a []uint) []uint {
	result := Interface(a)
	return result.Interface().([]uint)
}

func Uint8(a []uint8) []uint8 {
	result := Interface(a)
	return result.Interface().([]uint8)
}

func Uint16(a []uint16) []uint16 {
	result := Interface(a)
	return result.Interface().([]uint16)
}

func Uint32(a []uint32) []uint32 {
	result := Interface(a)
	return result.Interface().([]uint32)
}

func Uint64(a []uint64) []uint64 {
	result := Interface(a)
	return result.Interface().([]uint64)
}

func Float32(a []float32) []float32 {
	result := Interface(a)
	return result.Interface().([]float32)
}

func Float64(a []float64) []float64 {
	result := Interface(a)
	return result.Interface().([]float64)
}
