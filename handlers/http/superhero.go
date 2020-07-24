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
	repoAppearance "github.com/CezarGarrido/careers/repositories/superhero/appearance"
	repoBiography "github.com/CezarGarrido/careers/repositories/superhero/biography"
	repoConnections "github.com/CezarGarrido/careers/repositories/superhero/connections"
	repoImage "github.com/CezarGarrido/careers/repositories/superhero/image"
	repoPowerstats "github.com/CezarGarrido/careers/repositories/superhero/powerstats"
	repoWork "github.com/CezarGarrido/careers/repositories/superhero/work"
	"github.com/CezarGarrido/careers/util"
	"github.com/google/uuid"
)

func NewSuperHero(db *driver.DB) *SuperHero {
	return &SuperHero{
		superHeroRepo: repoSuperHero.NewPgSQLSuperHeroRepo(db.SQL),
	}
}

type SuperHero struct {
	appearanceRepo  repoAppearance.AppearanceRepo
	biographyRepo   repoBiography.BiographyRepo
	connectionsRepo repoConnections.ConnectionsRepo
	imageRepo       repoImage.ImageRepo
	powerstatsRepo  repoPowerstats.PowerstatsRepo
	superHeroRepo   repoSuperHero.SuperHeroRepo
	workRepo        repoWork.WorkRepo
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

	err = superHero.Validate()
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
	//Power stats create
	superHero.Powerstats.SuperID = superHero.ID
	_, err = this.powerstatsRepo.Create(ctx, superHero.Powerstats)
	if err != nil {
		log.Println(err.Error())
		util.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	//Bio stats create
	superHero.Biography.SuperID = superHero.ID
	_, err = this.biographyRepo.Create(ctx, superHero.Biography)
	if err != nil {
		log.Println(err.Error())
		util.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	//Biography create
	superHero.Biography.SuperID = superHero.ID
	_, err = this.biographyRepo.Create(ctx, superHero.Biography)
	if err != nil {
		log.Println(err.Error())
		util.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	//Appearance create
	superHero.Appearance.SuperID = superHero.ID
	_, err = this.appearanceRepo.Create(ctx, superHero.Appearance)
	if err != nil {
		log.Println(err.Error())
		util.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	//Work create
	superHero.Work.SuperID = superHero.ID
	_, err = this.workRepo.Create(ctx, superHero.Work)
	if err != nil {
		log.Println(err.Error())
		util.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	//Connections create
	superHero.Connections.SuperID = superHero.ID
	_, err = this.connectionsRepo.Create(ctx, superHero.Connections)
	if err != nil {
		log.Println(err.Error())
		util.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	//Image create
	superHero.Image.SuperID = superHero.ID
	_, err = this.imageRepo.Create(ctx, superHero.Image)
	if err != nil {
		log.Println(err.Error())
		util.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	util.ResponseJSON(w, http.StatusOK, superHero)
}
