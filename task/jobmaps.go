package task

import (
	"errors"
	"reflect"
)

type stubMapping map[string]interface{}

var StubStorage = stubMapping{}

func Call(funcName string, params ...interface{}) (result interface{}, err error) {
	function, ok := StubStorage[funcName]
	if !ok {
		err = errors.New("Function Not Found.")
		return
	}
	t := reflect.TypeOf(function)
	if t.Kind() != reflect.Func {
		err = errors.New("Not Function.")
		return
	}
	f := reflect.ValueOf(function)
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is out of index.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	var res []reflect.Value
	res = f.Call(in)
	if len(res) > 0 {
		result = res[0].Interface()
	}
	return
}
