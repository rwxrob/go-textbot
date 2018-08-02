package textbot

import (
	"fmt"
	"os"
	"strings"
)

type TextBot struct {
	responders []*Responder
	index      map[string]*Responder
	state      *State
}

func New() *TextBot {
	tb := new(TextBot)
	tb.index = map[string]*Responder{}
	tb.state = NewState()
	tb.state.Every = "10s"
	tb.state.Load()
	return tb
}

func (tb *TextBot) Add(responders ...*Responder) *TextBot {
	for _, r := range responders {
		tb.index[r.UUID] = r
		tb.responders = append(tb.responders, r)
	}
	return tb
}

func (tb *TextBot) Respond() {
	if len(os.Args) > 1 {
		tb.PrintResponseTo(strings.Join(os.Args[1:], " "))
	} else {
		tb.RespondToInput()
	}
	tb.state.Save()
}

func (tb *TextBot) PrintResponseTo(text string) {
	fmt.Println(tb.RespondTo(text))
}

func (tb *TextBot) RespondToInput() {
	//TODO
}

func (tb *TextBot) RespondToPort() {
}

// Eventually modify RespondTo to responders into blocks (arrays) of
// responders and run several goroutines to asynchronously find the
// first responder to answer. This should support thousands of
// responders.

func (tb *TextBot) RespondTo(text string) string {
	// TODO check the open session responder block first
	for _, r := range tb.responders {
		response := r.RespondTo(strings.Trim(text, " "), tb.state)
		if response != "" {
			return response
		}
	}
	return ""
}
