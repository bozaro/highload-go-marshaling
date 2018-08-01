package main

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type Go110 struct {
	*privateStruct
}

type privateStruct struct {
	Name string `json:"name"`
}

func main() {
	var go110 Go110
	//language=json
	if err := json.Unmarshal([]byte(`{
	"name": "Foo"
}`), &go110); err != nil {
		panic(err)
	}
	spew.Dump(go110)
	data, err := json.Marshal(go110)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
