package superheroapi

type SuperHeroResponse struct {
	Response string      `json:"response"`
	Results  []SuperHero `json:"results"`
	Error    string      `json:"error"`
}

type SuperHero struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Powerstats struct {
		Intelligence string `json:"intelligence"`
		Strength     string `json:"strength"`
		Speed        string `json:"speed"`
		Durability   string `json:"durability"`
		Power        string `json:"power"`
		Combat       string `json:"combat"`
	} `json:"powerstats"`
	Biography struct {
		FullName        string   `json:"full-name"`
		AlterEgos       string   `json:"alter-egos"`
		Aliases         []string `json:"aliases"`
		PlaceOfBirth    string   `json:"place-of-birth"`
		FirstAppearance string   `json:"first-appearance"`
		Publisher       string   `json:"publisher"`
		Alignment       string   `json:"alignment"`
	} `json:"biography"`
	Appearance struct {
		Gender    string   `json:"gender"`
		Race      string   `json:"race"`
		Height    []string `json:"height"`
		Weight    []string `json:"weight"`
		EyeColor  string   `json:"eye-color"`
		HairColor string   `json:"hair-color"`
	} `json:"appearance"`
	Work struct {
		Occupation string `json:"occupation"`
		Base       string `json:"base"`
	} `json:"work"`
	Connections struct {
		GroupAffiliation string `json:"group-affiliation"`
		Relatives        string `json:"relatives"`
	} `json:"connections"`
	Image struct {
		URL string `json:"url"`
	} `json:"image"`
}
