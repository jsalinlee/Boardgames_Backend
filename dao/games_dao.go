package dao

import (
	"log"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	. "github.com/jsalinlee/boardgames_server/models"
)

type GamesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "games"
)

// Connect Establish a connection to database
func (g *GamesDAO) Connect() {
	session, err := mgo.Dial(g.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(g.Database)
}

// FindAll Find list of games
func (g *GamesDAO) FindAll() ([]Game, error) {
	var games []Game
	err := db.C(COLLECTION).Find(bson.M{}).All(&games)
	return games, err
}

// FindByID Find a game by its id
func (g *GamesDAO) FindByID(id string) (Game, error) {
	var game Game
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&game)
	return game, err
}

// Insert a game into database
func (g *GamesDAO) Insert(game Game) error {
	err := db.C(COLLECTION).Insert(&game)
	return err
}

// Update an existing game
func (g *GamesDAO) Update(game Game) error {
	err := db.C(COLLECTION).UpdateId(game.ID, &game)
	return err
}

// Delete an existing game
func (g *GamesDAO) Delete(game Game) error {
	err := db.C(COLLECTION).Remove(&game)
	return err
}
