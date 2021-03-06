package user

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}
