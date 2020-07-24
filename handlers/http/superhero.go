package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/CezarGarrido/careers/driver"
	"github.com/CezarGarrido/careers/entities"
	repoSuperHero "github.com/CezarGarrido/careers/repositories/superhero"
	"github.com/CezarGarrido/careers/util"
	"github.com/google/uuid"
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
		util.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = json.Unmarshal(b, &superHero)
	if err != nil {
		log.Println(err.Error())
		util.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	superHero.CreatedAt = time.Now()
	superHero.UUID = uuid.New().String()

	newSuperHeroID, err := this.superHeroRepo.Create(ctx, &superHero)
	if err != nil {
		log.Println(err.Error())
		util.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	superHero.ID = newSuperHeroID

	util.ResponseJSON(w, http.StatusOK, superHero)
}
