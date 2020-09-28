package harddrive

import (
	"github.com/sergiorra/personal_exercises_go/025-go_convention_code"
)

type Db map[int]architecture.User

func (hd Db) Save(n int, u architecture.User) {
	hd[n] = u
}

func (hd Db) Retrieve(n int) architecture.User {
	return hd[n]
}
