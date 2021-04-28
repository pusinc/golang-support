package contain

import "reflect"

func Interface(s interface{}, v interface{}) bool {
	sValue, ok := s.(reflect.Value)
	if !ok {
		sValue = reflect.ValueOf(s)
	}
	if sValue.Type().Kind() == reflect.Ptr {
		sValue = sValue.Elem()
	}
	vValue, ok := v.(reflect.Value)
	if !ok {
		vValue = reflect.ValueOf(v)
	}
	if vValue.Type().Kind() == reflect.Ptr {
		vValue = vValue.Elem()
	}
	for i := 0; i < sValue.Len(); i++ {
		vv := sValue.Index(i)
		if reflect.DeepEqual(vv.Interface(), vValue.Interface()) {
			return true
		}
	}
	return false
}
