package std

import (
	"encoding/json"
	"fmt"
)

// The parts in backticks are struct tags, it tells json package to
// map Name to JSON key "name" and Age to "age"
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func JSONTest() {
	p := Person{Name: "Alice", Age: 30}

	// Encodes struct to JSON
	data, err := json.Marshal(p) // or MarshalIndent
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encodes person %v to JSON: %s\n", p, string(data))

	// Decodes JSON to struct
	// func Unmarshal(data []byte, v any) error
	// v must be a pointer so Unmarshal can modify the value we pass in
	var p2 Person
	if err := json.Unmarshal(data, &p2); err != nil {
		panic(err)
	}
	fmt.Printf("Decodes JSON bytes %v to struct %v\n", data, p2)

	// Decodes JSON into a flexible Go map
	dynamicJSON := `{"name":"Bob","age":25,"tags":["golang","stdlib"]}`
	var dynamic map[string]any
	if err := json.Unmarshal([]byte(dynamicJSON), &dynamic); err != nil {
		panic(err)
	}
	fmt.Printf("Decodes %s to %v\n", dynamicJSON, dynamic)

	name, ok := dynamic["name"].(string)
	if !ok {
		panic("failed to extract name")
	}
	fmt.Println("name: ", name)
	age, ok := dynamic["age"].(float64)
	if !ok {
		panic("failed to extract age")
	}
	fmt.Println("age: ", int(age))
	tags, ok := dynamic["tags"].([]any)
	if !ok {
		panic("failed to extract tags")
	}
	fmt.Println("tags: ", tags)
}
