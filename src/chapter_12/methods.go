package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func print(value interface{}) {
	v := reflect.ValueOf(value)

	valueType := v.Type()

	fmt.Printf("type %s\n", valueType)

	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()

		fmt.Printf("func (%s) %s%s\n", valueType, valueType.Method(i).Name,
			strings.TrimPrefix(methodType.String(), "func"))
	}
}

func main() {
	print(time.Hour)
	/*
		type time.Duration
		func (time.Duration) Hours() float64
		func (time.Duration) Microseconds() int64
		func (time.Duration) Milliseconds() int64
		func (time.Duration) Minutes() float64
		func (time.Duration) Nanoseconds() int64
		func (time.Duration) Round(time.Duration) time.Duration
		func (time.Duration) Seconds() float64
		func (time.Duration) String() string
		func (time.Duration) Truncate(time.Duration) time.Duration
	*/
}
