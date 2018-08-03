package models

import "github.com/globalsign/mgo/bson"

type Game struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `bson:"name" json:"name"`
	Desc       string        `bson:"desc" json:"desc"`
	Rating     float64       `bson:"rating" json:"rating,string"`
	MinPlayers int           `bson:"minPlayers" json:"minPlayers,string"`
	MaxPlayers int           `bson:"maxPlayers" json:"maxPlayers,string"`
}
