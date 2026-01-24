package interface_example

import (
	"encoding/json"
	"errors"
	"fmt"
)

// An empty interface (interface{} or any{}) can hold values of any type because every type implements at least zero methods.
type EmptyInterface interface{}

// Example function that accepts an empty interface
// generically and prints the value
func PrintValue(v EmptyInterface) string {
	return fmt.Sprintf("Value: %v", v)
}

// Another example for empty interface could be a container for any unknown data in JSON parsing: map[string]interface{}
// Original JSON:
// {
//   "code": 0,
//   "message": "ok",
//   "data": { // data is of unknown structure, could be anything
//     "id": 101,
//     "name": "Alice",
//     "age": 23
//   }
// }

type User struct {
	ID   int
	Name string
	Age  int
}

// Parse incoming JSON with unknown structure using empty interfaces
func ParseJSON(jsonData []byte) (User, error) {
	var raw map[string]interface{}

	// Unmarshal JSON into map with empty interface values
	// JSON → interface{}
	err1 := json.Unmarshal([]byte(jsonData), &raw)
	if err1 != nil {
		fmt.Println(err1)
		return User{}, err1
	}

	// Extract the "data" field
	data, ok := raw["data"]
	if !ok {
		fmt.Println("data field not found")
		return User{}, errors.New("data field not found")
	}

	// Remarshal the "data" field back to JSON
	// interface{} → JSON bytes
	bytes, err2 := json.Marshal(data)
	if err2 != nil {
		fmt.Println(err2)
		return User{}, err2
	}

	// Unmarshal the JSON bytes into a struct
	// JSON bytes → struct
	var user User
	if err3 := json.Unmarshal(bytes, &user); err3 != nil {
		fmt.Println(err3)
		return User{}, err3
	}
	return user, nil
}

// Outgoing Publishing with unknown payload using empty interface
func Publish(eventName string, payload interface{}) {
	fmt.Printf("event=%s payload=%v\n", eventName, payload)
}
