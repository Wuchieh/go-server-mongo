package mongo

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ISetID interface {
	SetID(any)
}

type IID interface {
	ID() bson.ObjectID
}
