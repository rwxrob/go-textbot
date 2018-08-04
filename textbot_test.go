package textbot

import (
	"fmt"
)

func ExampleState_SetDef_1() {
	s := New()
	s.Set("_", "prompt", "$ ")
	s.SetDef("_", "prompt", "> ")
	fmt.Println(s.Get("_", "prompt"))

	// Output:
	// $
}

func ExampleState_SetDef_2() {
	s := New()
	s.SetDef("_", "prompt", "$ ")
	s.Set("_", "prompt", "> ")
	fmt.Println(s.Get("_", "prompt"))

	// Output:
	// >
}
