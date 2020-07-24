package superheroapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var ErrFailedSearch = errors.New("Character with given name not found")
var ErrInvalidApiURL = errors.New("Invalid api url")

type SuperHeroClient struct {
	baseURL string
	conn    *http.Client
}

//NewSuperHeroClient: Configure and return new SuperHeroClient
func NewSuperHeroClient(apiURL, Token string) (*SuperHeroClient, error) {
	if !isValidUrl(apiURL) {
		return nil, ErrInvalidApiURL
	}

	return &SuperHeroClient{
		conn:    &http.Client{},
		baseURL: fmt.Sprintf(apiURL+"/%s", Token),
	}, nil
}

//FindSuperHeroesByName: Search super heroes by name
func (this *SuperHeroClient) FindSuperHeroesByName(value string) (superHero []SuperHero, err error) {

	response, err := this.conn.Get(fmt.Sprintf(
		this.baseURL+"/search/%s",
		value,
	))

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	var superResponse SuperHeroResponse
	err = json.Unmarshal(body, &superResponse)
	if err != nil {
		return
	}
	if superResponse.Error != "" {
		return nil, errors.New(superResponse.Error)
	}
	return superResponse.Results, nil
}

//isValidUrl: tests a string to determine if it is a well-structured url or not.
func isValidUrl(toTest string) bool {
	// https://golangcode.com/how-to-check-if-a-string-is-a-url/
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
