package entities

type SuperHero struct {
	Base
	SuperHeroApiID int64       `json:"super-hero-api-id"`
	Name           string      `json:"name"`
	Powerstats     Powerstats  `json:"powerstats"`
	Biography      Biography   `json:"biography"`
	Appearance     Appearance  `json:"appearance"`
	Work           Work        `json:"work"`
	Connections    Connections `json:"connections"`
	Image          Image       `json:"image"`
}
