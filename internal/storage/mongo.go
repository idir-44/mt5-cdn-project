package storage

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

func ConnectMongo() {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Erreur de connexion à MongoDB :", err)
	}

	MongoDB = client.Database(os.Getenv("DB_NAME"))
	log.Println("Connexion à MongoDB réussie")
}
