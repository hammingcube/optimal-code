package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestDecoding(t *testing.T) {
	data, err := ioutil.ReadFile("optimal-code-export.json")
	if err != nil {
		log.Fatal(err)
	}
	v := &Schema{}
	err = json.Unmarshal(data, v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s %#v\n", "Problems", v.Problems)
	fmt.Printf("%s %#v\n", "Solutions", v.Solutions)
	fmt.Printf("%s %#v\n", "Users", v.Users)
	fmt.Printf("%s %#v\n", "Submissions", v.Submissions)
}
