package main

import (
	"bytes"
	"fmt"
	"log"
	"reflect"

	"github.com/pkg/errors"
)

func encode(buffer *bytes.Buffer, value reflect.Value) error {
	switch value.Kind() {
	case reflect.Invalid:
		buffer.WriteString("nil")
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buffer, "%d", value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buffer, "%d", value.Uint())
	case reflect.String:
		fmt.Fprintf(buffer, "%q", value.String())
	case reflect.Ptr:
		return encode(buffer, value.Elem())
	case reflect.Array, reflect.Slice:
		buffer.WriteByte('(')

		for i := 0; i < value.Len(); i++ {
			if i > 0 {
				buffer.WriteByte(' ')
			}

			if err := encode(buffer, value.Index(i)); err != nil {
				return errors.WithStack(err)
			}
		}

		buffer.WriteByte(')')

	case reflect.Struct:
		buffer.WriteString("(\n")

		for i := 0; i < value.NumField(); i++ {
			fmt.Fprintf(buffer, "  (%s ", value.Type().Field(i).Name)

			if err := encode(buffer, value.Field(i)); err != nil {
				return errors.WithStack(err)
			}

			buffer.WriteString(")\n")
		}

		buffer.WriteByte(')')

	case reflect.Map:
		buffer.WriteString("(\n")

		for _, key := range value.MapKeys() {
			buffer.WriteString("    (")

			if err := encode(buffer, key); err != nil {
				return errors.WithStack(err)
			}

			buffer.WriteByte(' ')

			if err := encode(buffer, value.MapIndex(key)); err != nil {
				return errors.WithStack(err)
			}

			buffer.WriteString(")\n")
		}

		buffer.WriteString("  )")

	default:
		return fmt.Errorf("unsupported type: %s", value.Type())
	}

	return nil
}

func Marshal(value interface{}) ([]byte, error) {
	var buffer bytes.Buffer

	if err := encode(&buffer, reflect.ValueOf(value)); err != nil {
		return nil, errors.WithStack(err)
	}

	return buffer.Bytes(), nil
}

type User struct {
	Name         string
	Age          int
	Emails       []string
	SkillRatings map[string]int
}

func main() {
	user := User{
		Name:   "John Doe",
		Age:    21,
		Emails: []string{"johndoe1@email.com", "johndoe2@email.com"},
		SkillRatings: map[string]int{
			"Go":         9,
			"Rust":       7,
			"JavaScript": 10,
		},
	}

	sExpression, err := Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(sExpression))
	/*
			(
		  	(Name "John Doe")
		  	(Age 21)
		  	(Emails ("johndoe1@email.com" "johndoe2@email.com"))
		  	(SkillRatings (
		    	("Go" 9)
		    	("Rust" 7)
		    	("JavaScript" 10)
		  	))
		 	 )
	*/
}
