package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var ErrUnsupportedKind = errors.New("unsupported kind")

func populate(target reflect.Value, value string) error {
	switch target.Kind() {
	case reflect.String:
		target.SetString(value)
	case reflect.Int:
		integer, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return errors.WithStack(err)
		}
		target.SetInt(integer)
	case reflect.Bool:
		boolean, err := strconv.ParseBool(value)
		if err != nil {
			return errors.WithStack(err)
		}
		target.SetBool(boolean)
	default:
		return errors.WithMessage(ErrUnsupportedKind, fmt.Sprintf("%s is not supported", target.Type()))
	}

	return nil
}

func unpack(request *http.Request, out interface{}) error {
	if err := request.ParseForm(); err != nil {
		return errors.WithStack(err)
	}

	fields := make(map[string]reflect.Value)

	targetStruct := reflect.ValueOf(out).Elem()

	for i := 0; i < targetStruct.NumField(); i++ {
		structField := targetStruct.Type().Field(i)

		name := structField.Tag.Get("http")

		if name == "" {
			name = strings.ToLower(structField.Name)
		}

		fields[name] = targetStruct.Field(i)
	}

	for name, values := range request.Form {
		field := fields[name]

		if !field.IsValid() {
			continue
		}

		for _, value := range values {
			if field.Kind() == reflect.Slice {
				elem := reflect.New(field.Type().Elem()).Elem()

				if err := populate(elem, value); err != nil {
					return errors.WithStack(err)
				}

				field.Set(reflect.Append(field, elem))
			} else {
				if err := populate(field, value); err != nil {
					return errors.WithStack(err)
				}
			}
		}
	}

	return nil
}

func search(response http.ResponseWriter, request *http.Request) {
	var data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}

	data.MaxResults = 10

	if err := unpack(request, &data); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	// process request
}

func main() {

}
