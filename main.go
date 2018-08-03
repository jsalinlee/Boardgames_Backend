package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	. "github.com/jsalinlee/boardgames_server/config"
	. "github.com/jsalinlee/boardgames_server/dao"
	. "github.com/jsalinlee/boardgames_server/models"
)

var config = Config{}
var dao = GamesDAO{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter()
	okHeaders := handlers.AllowedOrigins([]string{"*"})
	okOrigins := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	okMethods := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})

	myRouter.HandleFunc("/games", AllGamesEndPoint).Methods("GET")
	myRouter.HandleFunc("/games", CreateGameEndPoint).Methods("POST")
	myRouter.HandleFunc("/games", UpdateGameEndPoint).Methods("PUT")
	myRouter.HandleFunc("/games/{id}", DeleteGameEndPoint).Methods("DELETE")
	myRouter.HandleFunc("/games/{id}", FindGameEndPoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(okHeaders, okOrigins, okMethods)(myRouter)))
	// log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func AllGamesEndPoint(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "AllGamesEndPoint not implemented yet!")
	games, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, games)
}

func FindGameEndPoint(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "FindGameEndPoint not implemented yet!")
	params := mux.Vars(r)
	id := params["id"]
	game := FindGame(w, id)
	respondWithJSON(w, http.StatusOK, game)
}

func CreateGameEndPoint(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "CreateGameEndPoint not implemented yet!")
	defer r.Body.Close()
	var game Game
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	fmt.Println("Made it here!")
	game.ID = bson.NewObjectId()
	if err := dao.Insert(game); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, game)
}

func UpdateGameEndPoint(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "UpdateGameEndPoint not implemented yet!")
	defer r.Body.Close()
	var game Game
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(game); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteGameEndPoint(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "DeleteGameEndPoint not implemented yet!")
	// defer r.Body.Close()
	// var game Game
	// if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
	// 	respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	// 	return
	// }
	params := mux.Vars(r)
	id := params["id"]
	game := FindGame(w, id)
	if err := dao.Delete(game); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func FindGame(w http.ResponseWriter, id string) (game Game) {
	game, err := dao.FindByID(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Game ID")
		return
	}
	return game
}

// func returnAllGames(w http.ResponseWriter, r *http.Request) {
// 	games := Games{
// 		Game{Name: "Dominion", Desc: "Deckbuilder", Rating: 4.5, MinPlayers: 2, MaxPlayers: 4},
// 		Game{Name: "Hand of Fate Ordeals", Desc: "Dungeon Deckbuilder", Rating: 3, MinPlayers: 1, MaxPlayers: 4},
// 	}
// 	fmt.Println("Endpoint Hit: returnAllGames")

// 	json.NewEncoder(w).Encode(games)
// }

// func returnSingleGame(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key := vars["id"]
// 	var1 := vars["var1"]
// 	var2 := vars["var2"]

// 	fmt.Println("Var 1: " + var1)
// 	fmt.Println("Var 2: " + var2)
// 	fmt.Fprintf(w, "Key: "+key)
// }
