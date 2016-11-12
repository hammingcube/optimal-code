package models

import (
	"fmt"
	"github.com/maddyonline/umpire"
)

type UserKey string
type ProblemKey string

type User struct {
	Id   UserKey `json:"id"`
	Name string  `json:"name"`
}

type Problem struct {
	Id        ProblemKey        `json:"id"`
	Name      string            `json:"name"`
	Title     string            `json:"title"`
	ShortDesc string            `json:"short_desc"`
	FullDesc  string            `json:"full_desc"`
	Templates map[string]string `json:"templates"`
}

type Solution struct {
	ProblemId ProblemKey `json:"problem_id"`
	*umpire.Payload
}

type Submission struct {
	ProblemId string `json:"problem_id"`
	Timestamp string `json:"timestamp"`
	Solution  *umpire.Payload
}

type Schema struct {
	Problems    map[ProblemKey]*Problem
	Solutions   map[ProblemKey]*Solution
	Users       map[UserKey]*User
	Submissions map[UserKey][]*Submission
}

type Store interface {
	CreateProblem(key ProblemKey, p *Problem) (error, *Problem)
	CreateSolution(key ProblemKey, s *Solution) (error, *Solution)
}

type InMemoryStore struct {
	Problems  map[ProblemKey]*Problem
	Solutions map[ProblemKey]*Solution
}

func (store *InMemoryStore) CreateProblem(key ProblemKey, p *Problem) (error, *Problem) {
	if p == nil {
		return nil, nil
	}
	if _, ok := store.Problems[key]; ok {
		return fmt.Errorf("Key exists"), nil
	}
	store.Problems[key] = p
	return nil, p
}

func (store *InMemoryStore) CreateSolution(key ProblemKey, s *Solution) (error, *Solution) {
	if s == nil {
		return nil, nil
	}
	if _, ok := store.Solutions[key]; ok {
		return fmt.Errorf("Key exists"), nil
	}
	store.Solutions[key] = s
	return nil, s
}
