package responder

import (
	tb "github.com/robmuh/go-textbot"
	s "strings"
)

var Responder = &tb.Responder{

	UUID: "090c0e4b-663b-408b-8ee8-dd1e549a52ca",

	RespondTo: func(t string, c *tb.State) string {
		t = s.ToLower(t)
		if s.HasPrefix(t, "show state") {
			c.PrettyPrint()
		}
		return ""
	},
}
