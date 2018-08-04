package textbot

import (
	"encoding/json"
	"fmt"
)

type Data map[string]interface{}

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
		key := p[0].(string)
		if _, ok := d[key]; !ok {
			d[key] = Data{}
		}
		d[key].(Data).Set(p[1:]...)
	}

	if len(p) < 2 {
		panic(MissingParams)
	}
}

func (d Data) Pretty() string {
	byt, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return fmt.Sprintf("{\"ERROR\":\"%v\"}", err)
	}
	return string(byt)
}

func (d Data) Print() {
	fmt.Print(d.Pretty())
}

func (d Data) String() string {
	byt, err := json.Marshal(d)
	if err != nil {
		return fmt.Sprintf("{\"ERROR\":\"%v\"}", err)
	}
	return string(byt)
}
