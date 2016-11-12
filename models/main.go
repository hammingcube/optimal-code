package models

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/maddyonline/umpire"
	"net/http"
	"strconv"
)

type UserKey string
type ProblemKey string
type SubmissionKey string

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
	Id        SubmissionKey `json:"id"`
	UserId    UserKey       `json:"user_id"`
	ProblemId ProblemKey    `json:"problem_id"`
	Timestamp string        `json:"timestamp"`
	Solution  *Solution     `json:"solution"`
}

type Schema struct {
	Problems  map[ProblemKey]*Problem
	Solutions map[ProblemKey]*Solution
	Users     map[UserKey]*User
	//Submissions map[UserKey]map[SubmissionKey]*Submission
}

func PostProblemHandler(store ProblemStore) func(echo.Context) error {
	return func(c echo.Context) error {
		prob := &Problem{}
		if err := c.Bind(prob); err != nil {
			return err
		}
		err, stored := store.CreateProblem(prob.Id, prob)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, stored)
	}
}

type ProblemStore interface {
	CreateProblem(key ProblemKey, p *Problem) (error, *Problem)
	GetProblem(key ProblemKey) (error, *Problem)
	//UpdateProblem(key ProblemKey, p *Problem) (error, *Problem)
	//GetAllProblems() (error, []*Problem)
}

type SolutionStore interface {
	CreateSolution(key ProblemKey, s *Solution) (error, *Solution)
	GetSolution(key ProblemKey) (error, *Solution)
	//UpdateSolution(key ProblemKey, s *Solution) (error, *Solution)
	//GetAllSolutions() (error, []*Solution)
}

type UserStore interface {
	CreateUser(key UserKey, u *User) (error, *User)
	GetUser(key UserKey) (error, *User)
	//UpdateUser(key UserKey, u *User) (error, *User)
	//GetAllUsers() (error, []*User)
}

type SubmissionStore interface {
	CreateSubmission(key UserKey, sub *Submission) (error, *Submission)
	GetSubmission(key UserKey, subKey SubmissionKey) (error, *Submission)
	//UpdateSubmission(key UserKey, sub *Submission) (error, *Submission)
	//GetAllSubmissions(key UserKey) (error, []*Submission)
}

type Store interface {
	ProblemStore
	SolutionStore
	UserStore
	SubmissionStore
}

type InMemoryStore struct {
	Problems    map[ProblemKey]*Problem
	Solutions   map[ProblemKey]*Solution
	Users       map[UserKey]*User
	Submissions map[UserKey]map[SubmissionKey]*Submission
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
	if subMap, ok := store.Submissions[key]; ok {
		newKey := SubmissionKey(strconv.Itoa(len(subMap) + 1))
		store.Submissions[key][newKey] = sub
		sub.Id = newKey
	} else {
		store.Submissions[key] = map[SubmissionKey]*Submission{SubmissionKey("0"): sub}
		sub.Id = "0"
	}
	return nil, sub
}

func (store *InMemoryStore) GetProblem(key ProblemKey) (error, *Problem) {
	if _, ok := store.Problems[key]; !ok {
		return fmt.Errorf("Key does not exist"), nil
	}
	return nil, store.Problems[key]
}

func (store *InMemoryStore) GetSolution(key ProblemKey) (error, *Solution) {
	if _, ok := store.Solutions[key]; !ok {
		return fmt.Errorf("Key does not exist"), nil
	}
	return nil, store.Solutions[key]
}

func (store *InMemoryStore) GetUser(key UserKey) (error, *User) {
	if _, ok := store.Users[key]; !ok {
		return fmt.Errorf("Key does not exist"), nil
	}
	return nil, store.Users[key]
}

func (store *InMemoryStore) GetSubmission(key UserKey, subKey SubmissionKey) (error, *Submission) {
	if _, ok := store.Submissions[key]; !ok {
		return fmt.Errorf("Key does not exist"), nil
	}
	if _, ok := store.Submissions[key][subKey]; !ok {
		return fmt.Errorf("Key does not exist"), nil
	}
	return nil, store.Submissions[key][subKey]
}
