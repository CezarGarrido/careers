package entities

import "errors"

type SuperHero struct {
	Base
	SuperHeroApiID int64        `json:"super-hero-api-id"`
	Name           string       `json:"name"`
	Powerstats     *Powerstats  `json:"powerstats"`
	Biography      *Biography   `json:"biography"`
	Appearance     *Appearance  `json:"appearance"`
	Work           *Work        `json:"work"`
	Connections    *Connections `json:"connections"`
	Image          *Image       `json:"image"`
}

//Validate: Simple validate required parameters
func (this *SuperHero) Validate() error {
	if this.Name == "" {
		return errors.New("Name is required")
	}
	return nil
}
