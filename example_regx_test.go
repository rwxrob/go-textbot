package textbot_test

import (
	"fmt"

	tb "github.com/robmuh/go-textbot"
)

var greet1 = tb.X(`How's it goin there, Mr. Rob?`)
var greet2 = tb.X(`(?i)how's it goin there, mr. rob?`)
var greet3 = tb.X(`(?i)how('|\si)s it goin(g|')? there,? mr.? rob`)

func Example() {
	text := "How's   it goin  there, Mr. Rob?"

	fmt.Println(greet1.MatchString(text))
	fmt.Println(greet1.M(text))
	fmt.Println(greet2.M(text))
	fmt.Println(greet3.M(text))

	fmt.Println()

	text = "how's it going there mr rob"
	fmt.Println(greet1.M(text))
	fmt.Println(greet2.M(text))
	fmt.Println(greet3.M(text))

	fmt.Println()

	text = "There is   a lot   of space here."
	fmt.Println(tb.CrunchSpace(text))
	fmt.Println(tb.SpaceToRegx(text))

	// Output:
	// true
	// true
	// true
	// true
	//
	// false
	// false
	// true
	//
	// There is a lot of space here.
	// There\s+is\s+a\s+lot\s+of\s+space\s+here.
}
