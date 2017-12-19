package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	funcs := map[string]interface{}{
		"foo":foo,
		"bar":bar,
	}
	fmt.Println(Call(funcs, "bar", 1,2,3))
}

func Call(m map[string]interface{}, name string, params ...interface{})(result []reflect.Value, err error){
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("the number of params is not adapted")
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	result = f.Call(in)
	return
}

func foo(){
	fmt.Println("in foo")
}

func bar(a,b,c int)int{
	fmt.Printf("in bar:%d, %d, %d\n", a,b,c)
	return a+b+c
}