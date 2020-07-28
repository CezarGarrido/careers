package entities

type Biography struct {
	Base
	SuperID         int64    `json:"super-id"`
	FullName        string   `json:"full-name"`
	AlterEgos       string   `json:"alter-egos"`
	Aliases         []string `json:"aliases"`
	PlaceOfBirth    string   `json:"place-of-birth"`
	FirstAppearance string   `json:"first-appearance"`
	Publisher       string   `json:"publisher"`
	Alignment       string   `json:"alignment"`
}

func NewBiography(SuperID int64, FullName, AlterEgos string, Aliases []string, PlaceOfBirth, FirstAppearance, Publisher, Alignment string) *Biography {
	return &Biography{
		SuperID:         SuperID,
		FullName:        FullName,
		AlterEgos:       AlterEgos,
		Aliases:         Aliases,
		PlaceOfBirth:    PlaceOfBirth,
		FirstAppearance: FirstAppearance,
		Publisher:       Publisher,
		Alignment:       Alignment,
	}
}
