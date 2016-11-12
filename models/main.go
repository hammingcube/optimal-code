package models

import (
	"github.com/maddyonline/umpire"
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
