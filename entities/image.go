package entities


type Image struct {
	Base
	SuperID int64  `json:"super-id"`
	URL     string `json:"url"`
}
