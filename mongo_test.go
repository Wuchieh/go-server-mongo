package mongo_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	mongo "github.com/Wuchieh/go-server-mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	mongo.Model `bson:",inline"`
	Name        string
}

func TestDB(t *testing.T) {
	if err := mongo.Setup(mongo.Config{
		Host:     "192.168.213.13",
		Port:     27017,
		Database: "wuchieh-dev",
		Username: "mongo_KTtKCG",
		Password: "mongo_NjJcc2",
	}); err != nil {
		t.Fatal(err)
	}

	u := User{
		Name: strconv.FormatInt(time.Now().UnixNano(), 36),
	}
	u.Init()

	if one, err := mongo.
		GetDatabase().
		Collection("users").
		InsertOne(context.Background(), &u); err != nil {
		return
	} else {
		t.Log(u.ID())
		u.SetID(one.InsertedID)
		t.Log(u.ID())
	}

	var result User
	if err := mongo.
		GetDatabase().
		Collection("users").
		FindOne(context.Background(), bson.M{"_id": u.ID()}).
		Decode(&result); err != nil {
		t.Fatal(err)
	}

	t.Log(result)
	t.Log(u == result)
	t.Log(&u == &result)
}
