// Exports utility to recursively print whole value structure.

package display

import (
	"fmt"
	"reflect"
)

func display(path string, value reflect.Value) {
	switch value.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), value.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, value.Type().Field(i).Name)
			display(fieldPath, value.Field(i))
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, fmt.Sprint(key)), value.MapIndex(key))
		}
	case reflect.Ptr:
		if value.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), value.Elem())
		}
	case reflect.Interface:
		if value.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, value.Elem().Type())
			display(path+".value", value.Elem())
		}
	default:
		fmt.Printf("%s = %s\n", path, fmt.Sprint(value))
	}
}

func Display(name string, value interface{}) {
	display(name, reflect.ValueOf(value))
}
