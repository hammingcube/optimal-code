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

func try_creating_users(store Store, user *User) error {
	err, stored := store.CreateUser(user.Id, user)
	if err != nil {
		return err
	}
	fmt.Printf("Stored: %#v", stored)
	return nil
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

func try_creating_submissions(store Store, sub *Submission) error {
	err, stored := store.CreateSubmission(sub.UserId, sub)
	if err != nil {
		return err
	}
	fmt.Printf("Stored: %#v", stored)
	return nil
}

func TestStoreImpl(t *testing.T) {
	store := &InMemoryStore{
		Problems:    map[ProblemKey]*Problem{},
		Solutions:   map[ProblemKey]*Solution{},
		Users:       map[UserKey]*User{},
		Submissions: map[UserKey][]*Submission{},
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
	user := &User{"abc", "james"}

	sub := &Submission{
		UserId:    user.Id,
		ProblemId: prob.Id,
		Solution:  soln,
	}

	try_creating_problems(store, prob)
	try_creating_solutions(store, soln)
	try_creating_users(store, user)
	try_creating_submissions(store, sub)
}
