package models

type Country struct {
	Idx  int    `json:"idx" db:"idx"`
	Name string `json:"name" db:"name"`
}
