package entities

type Connections struct {
	Base
	SuperID          int64  `json:"super-id"`
	GroupAffiliation string `json:"group-affiliation"`
	Relatives        string `json:"relatives"`
}
