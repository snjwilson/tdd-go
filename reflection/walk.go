package reflection

import "reflect"

func walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		switch f.Kind() {
		case reflect.String:
			fn(f.String())
		case reflect.Struct:
			walk(f.Interface(), fn)
		}
	}
}
