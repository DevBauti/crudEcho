package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// coneccion mongo
func ConnectMongo() *mongo.Client {

	serverApioption := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().ApplyURI(EnvMongoUri()).SetServerAPIOptions(serverApioption)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client

}

// int DB
var DB *mongo.Client = ConnectMongo()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golangAPI").Collection(collectionName)
	return collection
}
