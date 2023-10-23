package main

import (
	"fmt"

	encodingJson "encoding/json"
)

type TestType struct {
	Test int
}

func main() {
	json := []byte(`{"Test":1}`)

	var test = TestType{}

	encodingJson.Unmarshal(json, &test)

	fmt.Printf("%+v\n", test)
}
