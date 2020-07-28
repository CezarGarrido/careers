package entities

type Image struct {
	Base
	SuperID int64  `json:"super-id"`
	URL     string `json:"url"`
}

func NewImage(SuperID int64, URL string) *Image {
	return &Image{
		SuperID: SuperID,
		URL:     URL,
	}
}
