package entities

type Powerstats struct {
	Base
	SuperID      int64  `json:"super-id"`
	Intelligence string `json:"intelligence"`
	Strength     string `json:"strength"`
	Speed        string `json:"speed"`
	Durability   string `json:"durability"`
	Power        string `json:"power"`
	Combat       string `json:"combat"`
}

func NewPowerstats(SuperID int64, Intelligence, Strength, Speed, Durability, Power, Combat string) *Powerstats {
	return &Powerstats{
		SuperID:      SuperID,
		Intelligence: Intelligence,
		Strength:     Strength,
		Speed:        Speed,
		Durability:   Durability,
		Power:        Power,
		Combat:       Combat,
	}
}
