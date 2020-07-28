package entities

type Appearance struct {
	Base
	SuperID   int64    `json:"super-id"`
	Gender    string   `json:"gender"`
	Race      string   `json:"race"`
	Height    []string `json:"height"`
	Weight    []string `json:"weight"`
	EyeColor  string   `json:"eye-color"`
	HairColor string   `json:"hair-color"`
}

func NewAppearance(SuperID int64, Gender string, Race string, Height, Weight []string, EyeColor string, HairColor string) *Appearance {
	return &Appearance{
		SuperID:   SuperID,
		Gender:    Gender,
		Race:      Race,
		Height:    Height,
		Weight:    Weight,
		EyeColor:  EyeColor,
		HairColor: HairColor,
	}
}
