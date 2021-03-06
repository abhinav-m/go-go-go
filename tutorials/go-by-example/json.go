package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*

Go offers built-in support for JSON encoding and decoding,
 including to and from built-in and custom data types.
*/

// Structs to demonstrate encoding and decoding of custom types
type response1 struct {
	Page   int
	Fruits []string
}

// Only exported fields will be encoded/decoded in JSON.
// Fields must start with capital lettes to be exported.
type response2 struct {
	Page   int      `json:"page"` // Tags used to change json key names for encoding
	Fruits []string `json:"fruits"`
}

func main() {
	/*
		Encoding basic, atomic types
	*/
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	/*
		And here are some for slices and maps, which encode to JSON arrays
		 and objects as you’d expect.
	*/
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}

	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	/*
		You can use tags on struct field declarations to customize the encoded JSON key names.
		Check the definition of response2 above to see an example of such tags.
	*/
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Now let’s look at decoding JSON data into Go values.
	// Here’s an example for a generic data structure.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// We need to provide a variable where the JSON package can put the decoded data.
	// This map[string]interface{} will hold a map of strings to arbitrary data types.
	var dat map[string]interface{}

	// Decode, check for errors
	// Pass reference of interface
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	fmt.Println(dat)
	/*
		In order to use the values in the decoded map, we’ll need to convert them
		to their appropriate type.
		For example here we convert the value in num to the expected float64 type.
	*/
	num := dat["num"].(float64)

	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)

	fmt.Println(str1)

	//Accessing nested data requires a series of conversions.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}

	// decode JSON into custom data types
	// Grants additional type safety, eliminating need for assertions
	// while accessing decoded data
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	/*
		In the examples above we always used bytes and strings as intermediates
		between the data and JSON representation on standard out. We can also stream
		 JSON encodings directly to os.Writers like os.Stdout or even HTTP response bodies.
	*/
	// Streaming json encoders to os.stdout or HTTP response bodies
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}
