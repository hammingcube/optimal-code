package models

import (
	"encoding/json"
	"fmt"
	"github.com/maddyonline/umpire"
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

func try_creating_problems(store Store, prob *Problem) error {
	err, stored := store.CreateProblem(prob.Id, prob)
	if err != nil {
		return err
	}
	fmt.Printf("Stored: %#v", stored)
	return nil
}

func try_creating_solutions(store Store, soln *Solution) error {
	err, stored := store.CreateSolution(soln.ProblemId, soln)
	if err != nil {
		return err
	}
	fmt.Printf("Stored: %#v", stored)
	return nil
}

func TestStoreImpl(t *testing.T) {
	store := &InMemoryStore{
		Problems:  map[ProblemKey]*Problem{},
		Solutions: map[ProblemKey]*Solution{},
	}
	prob := &Problem{
		Id:    "abc",
		Name:  "prob-1",
		Title: "Sort numbers in array",
	}
	soln := &Solution{
		prob.Id,
		&umpire.Payload{
			Language: "cpp",
			Stdin:    "",
			Files: []*umpire.InMemoryFile{
				&umpire.InMemoryFile{"abc.cpp", ""},
			},
		},
	}
	try_creating_problems(store, prob)
	try_creating_solutions(store, soln)
}
