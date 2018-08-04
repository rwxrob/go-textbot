package textbot

import (
	"encoding/json"
	"fmt"
)

type Data map[string]interface{}

func (d Data) String() string {
	byt, err := json.Marshal(d)
	if err != nil {
		return fmt.Sprintf("{\"ERROR\":\"%v\"}", err)
	}
	return string(byt)
}

func (d Data) Get(keys ...string) interface{} {
	v := d.Get(keys[0])
	return v
}

func (d Data) Set(p ...interface{}) {

	if len(p) == 2 {
		d[p[0].(string)] = p[1]
		return
	}

	if len(p) > 2 {
		m := d[p[0]]
		if m == nil { // make a new one
		}
		if m.(type) != Data {
			panic(MustBeDataType)
		}
	}

	if len(p) < 2 {
		panic(MissingParams)
	}
}
