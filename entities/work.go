package entities

type Work struct {
	Base
	SuperID    int64  `json:"super-id"`
	Occupation string `json:"occupation"`
	BaseWork   string `json:"base"`
}
