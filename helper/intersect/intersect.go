package intersect

import "reflect"

func Interface(typ reflect.Type, x interface{}, y interface{}) reflect.Value {
	set := reflect.MakeSlice(typ, 0, 0)
	hash := reflect.MakeMap(typ)
	mapValue := reflect.ValueOf(true)

	xValue := reflect.ValueOf(x)
	yValue := reflect.ValueOf(y)
	for i := 0; i < xValue.Len(); i++ {
		hash.SetMapIndex(xValue.Index(i), mapValue)
	}
	for i := 0; i < yValue.Len(); i++ {
		v := yValue.Index(i)
		mapNode := hash.MapIndex(v)
		if !mapNode.IsNil() {
			set = reflect.Append(set, v)
		}
	}
	return set
}

func String(x []string, y []string) []string {
	if len(x) == 0 || len(y) == 0 {
		return make([]string, 0)
	}
	value := Interface(reflect.TypeOf(x[0]), x, y)
	return value.Interface().([]string)
}

func Int(x []int, y []int) []int {
	if len(x) == 0 || len(y) == 0 {
		return make([]int, 0)
	}
	value := Interface(reflect.TypeOf(x[0]), x, y)
	return value.Interface().([]int)
}

func Int8(x []int8, y []int8) []int8 {
	if len(x) == 0 || len(y) == 0 {
		return make([]int8, 0)
	}
	value := Interface(reflect.TypeOf(x[0]), x, y)
	return value.Interface().([]int8)
}

func Int16(x []int16, y []int16) []int16 {
	if len(x) == 0 || len(y) == 0 {
		return make([]int16, 0)
	}
	value := Interface(reflect.TypeOf(x[0]), x, y)
	return value.Interface().([]int16)
}

func Int32(x []int32, y []int32) []int32 {
	if len(x) == 0 || len(y) == 0 {
		return make([]int32, 0)
	}
	value := Interface(reflect.TypeOf(x[0]), x, y)
	return value.Interface().([]int32)
}

func Int64(x []int64, y []int64) []int64 {
	if len(x) == 0 || len(y) == 0 {
		return make([]int64, 0)
	}
	value := Interface(reflect.TypeOf(x[0]), x, y)
	return value.Interface().([]int64)
}

func Uint(x []uint, y []uint) []uint {
	if len(x) == 0 || len(y) == 0 {
		return make([]uint, 0)
	}
	value := Interface(reflect.TypeOf(x[0]), x, y)
	return value.Interface().([]uint)
}

func Uint8(x []uint8, y []uint8) []uint8 {
	if len(x) == 0 || len(y) == 0 {
		return make([]uint8, 0)
	}
	value := Interface(reflect.TypeOf(x[0]), x, y)
	return value.Interface().([]uint8)
}

func Uint16(x []uint16, y []uint16) []uint16 {
	if len(x) == 0 || len(y) == 0 {
		return make([]uint16, 0)
	}
	value := Interface(reflect.TypeOf(x[0]), x, y)
	return value.Interface().([]uint16)
}

func Uint32(x []uint32, y []uint32) []uint32 {
	if len(x) == 0 || len(y) == 0 {
		return make([]uint32, 0)
	}
	value := Interface(reflect.TypeOf(x[0]), x, y)
	return value.Interface().([]uint32)
}

func Uint64(x []uint64, y []uint64) []uint64 {
	if len(x) == 0 || len(y) == 0 {
		return make([]uint64, 0)
	}
	value := Interface(reflect.TypeOf(x[0]), x, y)
	return value.Interface().([]uint64)
}

func Float32(x []float32, y []float32) []float32 {
	if len(x) == 0 || len(y) == 0 {
		return make([]float32, 0)
	}
	value := Interface(reflect.TypeOf(x[0]), x, y)
	return value.Interface().([]float32)
}

func Float64(x []float64, y []float64) []float64 {
	if len(x) == 0 || len(y) == 0 {
		return make([]float64, 0)
	}
	value := Interface(reflect.TypeOf(x[0]), x, y)
	return value.Interface().([]float64)
}
