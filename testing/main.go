package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	file := "C:\\Users\\jjose\\Source\\go\\src\\github.com\\jjosephy\\foodapi\\testing\\test.json"

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

	// this is hte main map to all properties (map["list"])
	mlist := m["list"].(map[string]interface{})

	// Can access properties like this
	start := mlist["start"].(float64)
	end := mlist["end"].(float64)
	total := mlist["total"].(float64)

	fmt.Printf("%f %f %f", start, end, total)

	ml := mlist["item"].([]interface{})

	fmt.Print(ml)
}
