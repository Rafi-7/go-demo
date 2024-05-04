package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	opts := options.Client().ApplyURI("mongodb+srv://<username>:<password>@cluster0.up2e714.mongodb.net/")
	client, _ := mongo.Connect(context.TODO(), opts)
	defer client.Disconnect(context.Background())

	db := client.Database("SongsCollections")
	songsCollection := db.Collection("songs")

	song1 := Songs{
		Name:          "Periyonne Rahmane",
		Starring:      "Prithivi Raj",
		Singer:        "A.R.RAHMAN",
		MusicDirector: "A.R.RAHMAN",
		Duration:      "05:25",
		Released:      "2024",
	}

	song2 := Songs{
		Name:          "Chaiyya Chaiyya",
		Starring:      "Shahrukh Khan",
		Singer:        "Sukhwinder Singh",
		MusicDirector: "A.R.RAHMAN",
		Duration:      "07:31",
		Released:      "1998",
	}

	song3 := Songs{
		Name:          "Manjal Veyil",
		Starring:      "Kamal Hassan, Jyothika",
		Singer:        "Nakkhul, Hariharan, Vijay",
		MusicDirector: "Harris Jayaraj",
		Duration:      "05:55",
		Released:      "2006",
	}

	song4 := Songs{
		Name:          "Yaaradi Nee Mohini",
		Starring:      "Dhanush, Nayanthara",
		Singer:        "Udit Narayan",
		MusicDirector: "Yuvan Shankar Raja",
		Duration:      "05:03",
		Released:      "2008",
	}

	song5 := Songs{
		Name:          "Life of Ram",
		Starring:      "Vijay Sethupathi",
		Singer:        "Pradeep Kumar",
		MusicDirector: "Govind Vasantha",
		Duration:      "06:11",
		Released:      "2018",
	}

	song6 := Songs{
		Name:          "Thanjavooru Mannu Eduthu",
		Starring:      "Murali",
		Singer:        "Deva and Krishnaraj",
		MusicDirector: "Deva",
		Duration:      "05.34",
		Released:      "1997",
	}

	songList := []interface{}{song1, song2, song3, song4, song5, songs6}
	name, _ := songsCollection.InsertMany(context.Background(), songList)
	fmt.Println(name)

}

type Songs struct {
	Name          string
	Starring      string
	Singer        string
	MusicDirector string
	Duration      string
	Released      string
}
