package test

import (
	"fmt"
	"reflect"
)

func IsEqual(t1, t2 interface{}) bool {
	t1r := reflect.ValueOf(&t1).Elem()

	t2r := reflect.ValueOf(&t2).Elem()

	for x := 0; x < t1r.NumField(); x++ {
		field1 := t1r.Field(x)
		field2 := t2r.Field(x)

		switch field1.Kind() {
		case reflect.Pointer:
			if reflect.Indirect(field1.Elem()).Interface() != reflect.Indirect(field2.Elem()).Interface() {
				return false
			}
		case reflect.Slice:
			if field1.Len() != field2.Len() {
				return false
			}
			for i := 0; i < field1.Len(); i++ {
				if field1.Index(i).Interface() != field2.Index(i).Interface() {
					return false
				}
			}
		case reflect.Array:
			for i := 0; i < field1.Len(); i++ {
				fmt.Println(field1.Index(i))
			}

		default:
			if field1.Interface() != field2.Interface() {
				return false
			}
		}
	}

	return true
}

func Hueta(any interface{}, any2 interface{}) bool {
	fmt.Println(any2, any)
	reflect.TypeOf(any2).Kind()
	return false
}
