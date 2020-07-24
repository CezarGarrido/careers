package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/CezarGarrido/careers/driver"
	"github.com/CezarGarrido/careers/entities"
	repoSuperHero "github.com/CezarGarrido/careers/repositories/superhero"
)

func NewSuperHero(db *driver.DB) *SuperHero {
	return &SuperHero{
		superHeroRepo: repoSuperHero.NewPgSQLSuperHeroRepo(db.SQL),
	}
}

type SuperHero struct {
	superHeroRepo repoSuperHero.SuperHeroRepo
}

func (this *SuperHero) Create(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var superHero entities.SuperHero

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = json.Unmarshal(b, &superHero)
	if err != nil {
		log.Println(err.Error())
		return
	}

	newSuperHeroID, err := this.superHeroRepo.Create(ctx, &superHero)
	if err != nil {
		log.Println(err.Error())
		return
	}

}
