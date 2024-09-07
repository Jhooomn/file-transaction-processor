package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbUser   = os.Getenv("DB_USER")
	dbPsw    = os.Getenv("DB_PSW")
	dbDomain = os.Getenv("DB_DOMAIN")
	dbName   = os.Getenv("DB_NAME")
)

func NewConnection() (*mongo.Client, func()) {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	dbUrl := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=%s", dbUser, dbPsw, dbDomain, dbName)
	opts := options.Client().ApplyURI(dbUrl).SetServerAPIOptions(serverAPI)

	ctx := context.Background()

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(ctx, bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	close := func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}

	return client, close
}
