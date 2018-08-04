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
	d.Set("people", "robmuh", "name", "Mr. Rob")
	d.Set("people", "robmuh", "age", 50)
	fmt.Println(d)

	// Output:
	// {"people":{"robmuh":{"age":50,"name":"Mr. Rob"}}}
}

func ExampleData_Print() {
	d := Data{}
	d.Set("org", "people", "robmuh", "name", "Mr. Rob")
	d.Set("org", "people", "robmuh", "age", 50)
	d.Set("org", "people", "robmuh", "local", true)
	d.Set("org", "people", "robmuh", "empty", nil)
	d.Set("org", "people", "betropper", "name", "Ben")
	d.Set("org", "people", "betropper", "local", false)
	d.Print()

	// Output:
	// {
	//   "org": {
	//     "people": {
	//       "betropper": {
	//         "local": false,
	//         "name": "Ben"
	//       },
	//       "robmuh": {
	//         "age": 50,
	//         "empty": null,
	//         "local": true,
	//         "name": "Mr. Rob"
	//       }
	//     }
	//   }
	// }
}
