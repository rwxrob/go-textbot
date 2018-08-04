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
