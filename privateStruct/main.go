package main

import (
	"reflect"
)

type privateOne struct {
	M1 string
	M2 string
}

func sub(from interface{}) interface{} {
	t := reflect.ValueOf(from).Elem().Type()
	obj := reflect.New(t)
	p := obj.Elem()
	p.Field(0).SetString("foo")
	p.Field(1).SetString("bar")
	return obj.Interface()
}

func main() {
	value := sub(&privateOne{"a", "b"})

	if val, ok := value.(*privateOne); ok {
		println(val.M1, val.M2)
	} else {
		println(reflect.TypeOf(value).String())
	}
}
