package entities

type Connections struct {
	Base
	SuperID          int64  `json:"super-id"`
	GroupAffiliation string `json:"group-affiliation"`
	Relatives        string `json:"relatives"`
}

func NewConnections(SuperID int64, GroupAffiliation, Relatives string) *Connections {
	return &Connections{
		SuperID:          SuperID,
		GroupAffiliation: GroupAffiliation,
		Relatives:        Relatives,
	}
}
