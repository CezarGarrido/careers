package entities

type Work struct {
	Base
	SuperID    int64  `json:"super-id"`
	Occupation string `json:"occupation"`
	BaseWork   string `json:"base"`
}

func NewWork(SuperID int64, Occupation, BaseWork string) *Work {
	return &Work{
		SuperID:    SuperID,
		Occupation: Occupation,
		BaseWork:   BaseWork,
	}
}
