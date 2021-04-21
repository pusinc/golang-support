package contain

import "reflect"

func Interface(s interface{}, v interface{}) bool {
	sValue, ok := s.(reflect.Value)
	if !ok {
		sValue = reflect.ValueOf(s)
	}
	vValue, ok := v.(reflect.Value)
	if !ok {
		vValue = reflect.ValueOf(v)
	}
	for i := 0; i < sValue.Len(); i++ {
		vv := sValue.Index(i)
		if reflect.DeepEqual(vv, vValue) {
			return true
		}
	}
	return false
}
