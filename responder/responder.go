package responder

import (
	tb "github.com/robmuh/go-textbot"
)

var show = tb.X(`(?i)show state|cache`)
var dump = tb.X(`(?i)braindump`)

type Responder struct{}

func (r *Responder) UUID() string {
	return "090c0e4b-663b-408b-8ee8-dd1e549a52ca"
}

func (r *Responder) Keys() []string {
	return []string{}
}

func (r *Responder) RespondTo(t string, c *tb.State) string {
	if show.Is(t) || dump.Is(t) {
		return c.Pretty()
	}
	return ""
}

func (r *Responder) String() string { return tb.JSONString(r) }
