package textbot

import ()

func ExampleSet_1() {
	d := map[string]interface{}{}
	Set(d, "name", "Mr. Rob")
	Print(d)

	// Output:
	// {
	//   "name": "Mr. Rob"
	// }
}

func ExampleSet_2() {
	d := map[string]interface{}{}
	Set(d, "people", "robmuh", "name", "Mr. Rob")
	Set(d, "people", "robmuh", "age", 50)
	Print(d)

	// Output:
	// {
	//   "people": {
	//     "robmuh": {
	//       "age": 50,
	//       "name": "Mr. Rob"
	//     }
	//   }
	// }
}

func ExamplePrint() {
	d := map[string]interface{}{}
	Set(d, "org", "people", "robmuh", "name", "Mr. Rob")
	Set(d, "org", "people", "robmuh", "age", 50)
	Set(d, "org", "people", "robmuh", "local", true)
	Set(d, "org", "people", "robmuh", "empty", nil)
	Set(d, "org", "people", "betropper", "name", "Ben")
	Set(d, "org", "people", "betropper", "local", false)

	Print(d)

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
