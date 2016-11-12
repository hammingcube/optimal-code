package main

import (
	"encoding/json"
	"fmt"
	"github.com/maddyonline/umpire"
	"io/ioutil"
	"log"
)

type UserKey string
type ProblemKey string

type User struct {
	Name string `json:"name"`
}

type Problem struct {
	Name      string            `json:"name"`
	Title     string            `json:"title"`
	ShortDesc string            `json:"short_desc"`
	FullDesc  string            `json:"full_desc"`
	Templates map[string]string `json:"templates"`
}

type Submission struct {
	Problem   string `json:"problem"`
	Timestamp string `json:"timestamp"`
	Solution  *umpire.Payload
}

type Schema struct {
	Problems    map[ProblemKey]*Problem
	Solutions   map[ProblemKey]*umpire.Payload
	Users       map[UserKey]*User
	Submissions map[UserKey][]*Submission
}

func main() {
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
