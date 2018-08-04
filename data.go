package textbot

import (
	"encoding/json"
	"fmt"
)

type Data map[string]interface{}

func (d Data) Get(keys ...string) interface{} {

	if len(keys) == 1 {
		return d[keys[0]]
	}

	if len(keys) > 1 {
		i := d
		for n := 0; n < len(keys)-1; n++ {
			if newi, ok := i[keys[n]]; ok {
				i = newi.(Data)
			} else {
				return nil
			}
		}
		return i[keys[len(keys)-1]]
	}

	if len(keys) < 1 {
		panic(MissingParams)
	}

	return nil
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
