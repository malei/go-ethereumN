package tools

import (
	"fmt"
	"reflect"
)

func Print(T interface{}) {
	s := reflect.ValueOf(T).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf(" --%2d: %s (%s) = %v\n",
			i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
