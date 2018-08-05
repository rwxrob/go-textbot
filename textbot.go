package textbot

import (
	"fmt"
	"os"
	"strings"
)

type TextBot struct {
	responders []Responder
	index      map[string]Responder
	state      *State
	keys       []string
}

func New(responders ...Responder) *TextBot {
	tb := new(TextBot)
	tb.index = map[string]Responder{}
	tb.state = NewState()
	tb.Add(responders...)
	tb.state.Every = "10s"
	tb.SetDef("_", "prompt", "> ")
	tb.Save()
	return tb
}

func (tb *TextBot) Keys() []string { return tb.keys }

func (tb *TextBot) Set(p ...interface{}) {
	tb.state.Set(p...)
}

func (tb *TextBot) SetDef(p ...interface{}) {
	tb.state.SetDef(p...)
}

func (tb *TextBot) Get(keys ...string) interface{} {
	return tb.state.Get(keys...)
}

func (tb *TextBot) Add(responders ...Responder) *TextBot {
	for _, r := range responders {
		tb.index[r.UUID()] = r
		for _, k := range r.Keys() {
			tb.keys = append(tb.keys, k)
		}
		tb.responders = append(tb.responders, r)
	}
	return tb
}

func (tb *TextBot) LockFor(r Responder) {
	tb.Set("_", "lock", r.UUID())
}

//TODO Remove

func (tb *TextBot) Respond() {
	if len(os.Args) > 1 {
		tb.PrintResponseTo(strings.Join(os.Args[1:], " "))
	} else {
		tb.RespondToREPL()
	}
}

func (tb *TextBot) Save() {
	tb.state.Save()
}

func (tb *TextBot) PrintResponseTo(text string) {
	fmt.Println(tb.RespondTo(text))
}

func (tb *TextBot) RespondToREPL() {
	//TODO
}

//TODO Listen()

// Eventually modify RespondTo to responders into blocks (arrays) of
// responders and run several goroutines to asynchronously find the
// first responder to answer. This should support thousands of
// responders.

func (tb *TextBot) RespondTo(text string) string {
	text = strings.Trim(text, " ")

	// fetch the function refs that match uuids

	last := tb.Get("_", "last")
	lastr := []Responder{}
	for _, s := range last.([]interface{}) {
		lastr = append(lastr, tb.index[s.(string)])
	}

	// call the last 10 most recent

	if last != nil {
		for _, r := range lastr {
			response := r.RespondTo(text, tb.state)
			if response != "" {
				tb.push(r)
				tb.Save()
				return response
			}
		}
	}

	// then try them all (except the last)

	for _, r := range tb.responders {
		if isrecent(r, lastr) {
			continue
		}
		response := r.RespondTo(text, tb.state)
		if response != "" {
			tb.push(r)
			tb.Save()
			return response
		}
	}
	return ""
}

func isrecent(r Responder, last []Responder) bool {
	for _, one := range last {
		if r == one {
			return true
		}
	}
	return false
}

func (tb *TextBot) push(r Responder) {
	last := tb.Get("_", "last")
	uuid := r.UUID()
	new := []string{uuid}
	if last != nil {
		for n, s := range last.([]interface{}) {
			if s.(string) != uuid {
				new = append(new, s.(string))
			}
			if n >= 10 {
				break
			}
		}
	}
	tb.Set("_", "last", new)
}

func (tb *TextBot) String() string {
	return tb.state.String()
}

func (tb *TextBot) Pretty() string {
	return tb.state.Pretty()
}

func (tb *TextBot) Print() {
	tb.state.Print()
}

//TODO delegate add Marshal and Unmarshal JSON to state
