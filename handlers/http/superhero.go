package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	apiSuperHero "github.com/CezarGarrido/careers/services/superheroapi"
	"github.com/CezarGarrido/careers/util"
	"github.com/google/uuid"
)

func NewSuperHero(db *driver.DB, apiURL, Token string) *SuperHero {
	clienteApi, err := apiSuperHero.NewSuperHeroClient(apiURL, Token)
	if err != nil {
		panic(err)
	}
	return &SuperHero{
		superHeroRepo:   repoSuperHero.NewPgSQLSuperHeroRepo(db.SQL),
		appearanceRepo:  repoAppearance.NewPgSQLAppearanceRepo(db.SQL),
		biographyRepo:   repoBiography.NewPgSQLBiographyRepo(db.SQL),
		connectionsRepo: repoConnections.NewPgSQLConnectionsRepo(db.SQL),
		imageRepo:       repoImage.NewPgSQLImageRepo(db.SQL),
		powerstatsRepo:  repoPowerstats.NewPgSQLPowerstatsRepo(db.SQL),
		workRepo:        repoWork.NewPgSQLWorkRepo(db.SQL),
		superHeroApi:    clienteApi,
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
	superHeroApi    *apiSuperHero.SuperHeroClient
}

//Create: Creating a new super hero
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

	//Simple validate superhero name
	err = superHero.Validate()
	if err != nil {
		log.Println(err.Error())
		util.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	superHeroResponse, err := this.superHeroApi.FindSuperHeroesByName(superHero.Name)
	if err != nil {
		log.Println(err.Error())
		util.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var superHeroes []entities.SuperHero

	for _, superResponseApi := range superHeroResponse {

		id, _ := strconv.ParseInt(superResponseApi.ID, 10, 64)

		newSuperHero := entities.NewSuperHero(id, superResponseApi.Name)
		newSuperHero.CreatedAt = time.Now()
		newSuperHero.UUID = uuid.New().String()

		newSuperHeroID, err := this.superHeroRepo.Create(ctx, newSuperHero)
		if err != nil {
			log.Println(err.Error())
			util.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		newSuperHero.ID = newSuperHeroID

		//Power stats create
		powerstats := entities.NewPowerstats(newSuperHero.ID, superResponseApi.Powerstats.Intelligence, superResponseApi.Powerstats.Strength, superResponseApi.Powerstats.Speed, superResponseApi.Powerstats.Durability, superResponseApi.Powerstats.Power, superResponseApi.Powerstats.Combat)

		powerstats.CreatedAt = time.Now()
		powerstats.UUID = uuid.New().String()

		_, err = this.powerstatsRepo.Create(ctx, powerstats)
		if err != nil {
			log.Println(err.Error())
			util.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		//Bio stats create
		biography := entities.NewBiography(newSuperHero.ID, superResponseApi.Biography.FullName, superResponseApi.Biography.AlterEgos, superResponseApi.Biography.Aliases, superResponseApi.Biography.PlaceOfBirth, superResponseApi.Biography.FirstAppearance, superResponseApi.Biography.Publisher, superResponseApi.Biography.Alignment)

		biography.CreatedAt = time.Now()
		biography.UUID = uuid.New().String()

		_, err = this.biographyRepo.Create(ctx, biography)
		if err != nil {
			log.Println(err.Error())
			util.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		//Appearance create
		appearance := entities.NewAppearance(newSuperHero.ID, superResponseApi.Appearance.Gender, superResponseApi.Appearance.Race, superResponseApi.Appearance.Height, superResponseApi.Appearance.Weight, superResponseApi.Appearance.EyeColor, superResponseApi.Appearance.HairColor)
		appearance.CreatedAt = time.Now()
		appearance.UUID = uuid.New().String()
		_, err = this.appearanceRepo.Create(ctx, appearance)
		if err != nil {
			log.Println(err.Error())
			util.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		//Work create
		work := entities.NewWork(newSuperHero.ID, superResponseApi.Work.Occupation, superResponseApi.Work.Base)
		work.CreatedAt = time.Now()
		work.UUID = uuid.New().String()
		_, err = this.workRepo.Create(ctx, work)
		if err != nil {
			log.Println(err.Error())
			util.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		//Connections create
		connections := entities.NewConnections(newSuperHero.ID, superResponseApi.Connections.GroupAffiliation, superResponseApi.Connections.Relatives)
		connections.CreatedAt = time.Now()
		connections.UUID = uuid.New().String()
		_, err = this.connectionsRepo.Create(ctx, connections)
		if err != nil {
			log.Println(err.Error())
			util.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		//Image create
		image := entities.NewImage(newSuperHero.ID, superResponseApi.Image.URL)
		image.CreatedAt = time.Now()
		image.UUID = uuid.New().String()

		_, err = this.imageRepo.Create(ctx, image)
		if err != nil {
			log.Println(err.Error())
			util.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
			return
		}
		newSuperHero.Powerstats = powerstats
		newSuperHero.Biography = biography
		newSuperHero.Image = image
		newSuperHero.Work = work
		newSuperHero.Connections = connections
		newSuperHero.Appearance = appearance
		superHeroes = append(superHeroes, *newSuperHero)
	}

	util.ResponseJSON(w, http.StatusOK, superHeroes)
}
