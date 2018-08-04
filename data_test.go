package textbot

import "fmt"

func ExampleData_String() {
	d := Data{"name": "Mr. Rob"}
	fmt.Println(d)

	// Output:
	// {"name":"Mr. Rob"}
}

func ExampleData_Set_1() {
	d := Data{}
	d.Set("name", "Mr. Rob")
	fmt.Println(d)

	// Output:
	// {"name":"Mr. Rob"}
}

func ExampleData_Set_2() {
	d := Data{}
	d.Set("people", "name", "Mr. Rob")
	fmt.Println(d)

	// Output:
	// {"people":{"name":"Mr. Rob"}}
}
