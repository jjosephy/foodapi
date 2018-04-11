package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	file := "/Users/jjosephy/Source/go/src/github.com/jjosephy/testing/test.json"

	dat, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("error reading the file")
		panic(1)
	}

	var i interface{}
	err = json.Unmarshal([]byte(dat), &i)
	if err != nil {
		fmt.Println("error unmarshaling json")
		panic(1)
	}

	m := i.(map[string]interface{})

	ml := m["list"].(map[string]interface{})["item"].([]interface{})

	fmt.Print(ml)
}
