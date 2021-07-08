package main

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"text/scanner"
)

type lexer struct {
	scan         scanner.Scanner
	currentToken rune
}

func (lex *lexer) next() {
	lex.currentToken = lex.scan.Scan()
}

func (lex *lexer) text() string {
	return lex.scan.TokenText()
}

func (lex *lexer) expect(expected rune) {
	if lex.currentToken != expected {
		log.Fatal(fmt.Sprintf("got %q, expected %q", lex.text(), expected))
	}
}

func (lex *lexer) consume(expected rune) {
	lex.expect(expected)

	lex.next()
}

func read(lex *lexer, value reflect.Value) {
	switch lex.currentToken {
	case scanner.Ident:
		if lex.text() == "nil" {
			value.Set(reflect.Zero(value.Type()))
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text())
		value.SetString(s)
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text())
		value.SetInt(int64(i))
	case '(':
		lex.next()
		readList(lex, value)
	default:
		log.Fatal(fmt.Sprintf("unexpected token %q", lex.text()))
	}

	lex.next()
}

func readList(lex *lexer, value reflect.Value) {
	switch value.Kind() {
	case reflect.Array:
		for i := 0; !isListEnd(lex); i++ {
			read(lex, value.Index(i))
		}
	case reflect.Slice:
		for isListEnd(lex) {
			item := reflect.New(value.Type().Elem()).Elem()
			read(lex, item)
			value.Set(reflect.Append(value, item))
		}
	case reflect.Struct:
		for !isListEnd(lex) {
			lex.consume('(')

			lex.expect(scanner.Ident)

			fieldName := lex.text()

			lex.next()

			read(lex, value.FieldByName(fieldName))

			lex.consume(')')
		}
	case reflect.Map:
		value.Set(reflect.MakeMap(value.Type()))

		for !isListEnd(lex) {
			lex.consume('(')

			key := reflect.New(value.Type().Key()).Elem()

			read(lex, key)

			value := reflect.New(value.Type().Elem()).Elem()

			read(lex, value)

			value.SetMapIndex(key, value)

			lex.consume(')')
		}
	default:
		log.Fatal(fmt.Sprintf("cannot decod list into %v", value.Type()))
	}
}

func isListEnd(lex *lexer) bool {
	switch lex.currentToken {
	case scanner.EOF:
		log.Fatal("end of file")
	case ')':
		return true
	}

	return false
}

func Unmarshal(data []byte, out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}

	lex.scan.Init(bytes.NewReader(data))

	lex.next()

	defer func() {
		if reason := recover(); reason != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, reason)
		}
	}()

	read(lex, reflect.ValueOf(out).Elem())

	return nil
}

func main() {

}
