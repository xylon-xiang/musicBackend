package module

import (
	"MusicBackend/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var MusicInfoCol *mongo.Collection

func connectMongo() {

	// create a client
	client, err := mongo.NewClient(options.Client().
		ApplyURI(config.Config.DatabaseConfig.MongoConfig.DBAddress))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*
		time.Duration(config.Config.DatabaseConfig.MongoConfig.TimeOut))
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	musicInfoDatabase := client.Database(config.Config.DatabaseConfig.MongoConfig.DBName)

	MusicInfoCol = musicInfoDatabase.
		Collection(config.Config.DatabaseConfig.MongoConfig.MusicInfoCollection)

}

//func init() {
//	connectMongo()
//}
