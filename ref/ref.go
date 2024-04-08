package main

import (
	"POMOIKA/ref/test"
	"fmt"
	"reflect"
)

func main() {
	test1 := test.Gog{Hoh: "ch_to_to"}
	v := reflect.ValueOf(test1)
	field := v.FieldByName("Hoh")
	if field.IsValid() {
		value := field.Interface()
		fmt.Printf("%T", value)
	} else {
		fmt.Println("nifiga")
	}
}
