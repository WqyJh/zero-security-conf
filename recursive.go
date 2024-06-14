package confz

import "reflect"

func recursive(v reflect.Value) (_ reflect.Value, success bool) {
	switch v.Kind() {
	case reflect.Ptr:
		v = v.Elem()
		if !v.IsValid() {
			return
		}
		return recursive(v)
	case reflect.Interface:
		v = v.Elem()
		return recursive(v)
	case reflect.Struct:
		if v.Type() == reflect.TypeOf(SecurityConf{}) {
			return v, true
		}
		for i := 0; i < v.NumField(); i += 1 {
			ret, ok := recursive(v.Field(i))
			if ok {
				return ret, true
			}
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i += 1 {
			ret, ok := recursive(v.Index(i))
			if ok {
				return ret, true
			}
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			ret, ok := recursive(v.MapIndex(key))
			if ok {
				return ret, true
			}
		}
	default:
	}
	return
}
