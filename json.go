package textbot

import (
	"encoding/json"
	"fmt"
)

// JSONString can be used to easily fulfill the String() method of the
// Responder interface:
//
//       func (r *Responder) String() string { return tb.JSONString(r) }
//
// Returns the marshalled JSON of the structure (st) or a JSON map with
// only ERROR set to the error encountered during marshalling.
func JSONString(st interface{}) string {
	byt, err := json.Marshal(st)
	if err != nil {
		return fmt.Sprintf("{\"ERROR\":\"%v\"}", err)
	}
	return string(byt)
}
