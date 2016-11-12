package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Problem struct {
	Name      string            `json:"name"`
	Title     string            `json:"title"`
	ShortDesc string            `json:"short_desc"`
	FullDesc  string            `json:"full_desc"`
	Templates map[string]string `json:"templates"`
}

func main() {
	data, err := ioutil.ReadFile("optimal-code-export.json")
	if err != nil {
		log.Fatal(err)
	}
	v := &struct {
		Problems map[string]*Problem
	}{}
	err = json.Unmarshal(data, v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", v)

	for k := range v.Problems {
		fmt.Printf("%s:\n", k)
		fmt.Printf("%#v\n", *v.Problems[k])
	}
}
