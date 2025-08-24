package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var (
	client   *mongo.Client
	database *mongo.Database
)

func Setup(cfg Config) error {
	option := options.Client()
	option.ApplyURI(cfg.getUri())
	option.SetAuth(options.Credential{
		Username: cfg.Username,
		Password: cfg.Password,
	})

	var err error
	client, err = mongo.Connect(option)
	if err != nil {
		return err
	}

	if cfg.Database != "" {
		database = client.Database(cfg.Database)
	}

	return client.Ping(context.Background(), readpref.Primary())
}

func GetClient() *mongo.Client {
	return client
}

func GetDatabase() *mongo.Database {
	return database
}
