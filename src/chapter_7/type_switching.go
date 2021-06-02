package main

import "fmt"

func sqlQuote(value interface{}) string {
	switch value := value.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", value)
	case bool:
		if value {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return value // not implemented
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", value, value))
	}
}

func main() {
	fmt.Println(sqlQuote(nil))
	fmt.Println(sqlQuote(int(1)))
	fmt.Println(sqlQuote(uint(2)))
	fmt.Println(sqlQuote(true))
	fmt.Println(sqlQuote(false))
	fmt.Println(sqlQuote("hello world"))
	fmt.Println(sqlQuote([]int{1, 2, 3}))
}
