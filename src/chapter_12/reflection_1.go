package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// Go provides a mechanism to update variables and inspect
// their value at run time, to call their methods,
// and to apply the operations intrinsic to their representation,
// al without knowing their types at compile time.
// This mechanism is called reflection. Reflection also lets
// us treat types themselves as first-class values.
//
// Examples of packages that use reflection:
// fmt
// encoding/json
// encoding/xml
// text/template
// html/template
//
// Why reflection?
// Sometimes we need to write a function capable of dealing
// uniformly with values of types that don't satisfy a common interface,
// don't have a known representation, or don't exist at the time
// we design the function.

// Example:
// fmt.Sprint like function
func SprintWithoutReflection(value interface{}) string {
	// same as fmt.Stringer
	type stringer interface {
		String() string
	}

	switch value := value.(type) {
	case stringer:
		return value.String()
	case string:
		return value
	case int:
		return strconv.Itoa(value)
	case bool:
		if value {
			return "true"
		}
		return "false"

	default:
		// array, chan, func, map, pointer, slice, struct
		panic("???")
		// There is a problem: we can't add all the possible types to the type switch.
	}
}

// Reflection is provided by the reflect package.
// It defines two important types, Type and Value.
// A Type represents a Go type. It is an interface with many methods
// for discriminating among types and inspecting their components, like the fields
// of a struct or the parameters of a function. The sole implementation
// of reflect.Type is the type descriptor, the same entity that
// identifies the dynamic type of an interface value.//

func main() {
	// The reflect.TypeOf function accepts any interface{} and returns
	// its dynamic type as a reflect.Type
	t := reflect.TypeOf(3)  // t is a reflect.Type
	fmt.Println(t.String()) // int
	// this is the same as
	fmt.Printf("%T\n", 3)
	// which uses reflect.TypeOf internally.

	// The other important type in the reflect package is Value.
	// A reflect.Value can hold a value of any type. The reflect.ValueOf
	// function accepts any interface{} and returns a
	// reflect.Value containing the interface's dynamic value.
	// As with reflect.TypeOf, the result of reflect.ValueOf are always
	// concrete, but a reflect.Value can hold interface values too.
	v := reflect.ValueOf(3) // v is a reflect.Value
	fmt.Printf("%v\n", v)   //3
	fmt.Println(v.String()) // <int Value>
}

func SprintWithReflection(value interface{}) string {
	v := reflect.ValueOf(value)

	// there are missing cases
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct,reflect.Interface
		return v.Type().String() + " value"
	}
}
