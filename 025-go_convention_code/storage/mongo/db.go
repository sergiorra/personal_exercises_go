package mongo

import (
	"github.com/sergiorra/personal_exercises_go/025-go_convention_code"
)

type Db map[int]architecture.User

func (m Db) Save(n int, u architecture.User) {
	m[n] = u
}

func (m Db) Retrieve(n int) architecture.User {
	return m[n]
}
