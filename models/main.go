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
	UserId    UserKey    `json:"user_id"`
	ProblemId ProblemKey `json:"problem_id"`
	Timestamp string     `json:"timestamp"`
	Solution  *Solution  `json:"solution"`
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
	CreateUser(key UserKey, u *User) (error, *User)
	CreateSubmission(key UserKey, sub *Submission) (error, *Submission)
}

type InMemoryStore struct {
	Problems    map[ProblemKey]*Problem
	Solutions   map[ProblemKey]*Solution
	Users       map[UserKey]*User
	Submissions map[UserKey][]*Submission
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

func (store *InMemoryStore) CreateUser(key UserKey, u *User) (error, *User) {
	if u == nil {
		return nil, nil
	}
	if _, ok := store.Users[key]; ok {
		return fmt.Errorf("Key exists"), nil
	}
	store.Users[key] = u
	return nil, u
}

func (store *InMemoryStore) CreateSubmission(key UserKey, sub *Submission) (error, *Submission) {
	if sub == nil {
		return nil, nil
	}
	if arr, ok := store.Submissions[key]; ok {
		arr = append(arr, sub)
	} else {
		store.Submissions[key] = []*Submission{sub}
	}
	return nil, sub
}
