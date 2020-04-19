package models

import "fmt"

type Employee struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (u Employee) String() string {
	return fmt.Sprintf("Employee<%d %s %v>", u.Id, u.LastName, u.FirstName)
}
