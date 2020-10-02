package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"../helper"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client *mongo.Client

type Team struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	TeamName string             `json:"team_name,omitempty" bson:"team_name,omitempty"`
	City     string             `json:"city,omitempty" bson:"city,omitempty"`
}
type Player struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	TeamID     primitive.ObjectID `json:"team_id,omitempty" bson:"team_id,omitempty"`
	PlayerName string             `json:"player_name,omitempty" bson:"player_name,omitempty"`
}

func CreatePlayerEndpoint(w http.ResponseWriter, r *http.Request) {
	//
	w.Header().Set("Content-Type", "application/json")

	var player Player

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&player)

	// connect db
	collection := helper.ConnectDB("players")

	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), player)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func GetPlayerEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var players []Player
	collection := helper.ConnectDB("players")
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		helper.GetError(err, w)
		return
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var player Player
		err := cur.Decode(&player) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		players = append(players, player)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(players) // encode similar to serialize process.
}

// team
func CreateTeamEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var team Team

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&team)

	// connect db
	collection := helper.ConnectDB("teams")

	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), team)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func GetTeamEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var teams []Team
	collection := helper.ConnectDB("teams")
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		helper.GetError(err, w)
		return
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var team Team
		err := cur.Decode(&team) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		teams = append(teams, team)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(teams) // encode similar to serialize process.
}
