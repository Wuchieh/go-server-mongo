package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Model struct {
	Id       bson.ObjectID `bson:"_id,omitempty"`
	CreateAt time.Time     `bson:"create_at,omitempty"`
	UpdateAt time.Time     `bson:"update_at,omitempty"`
	DeleteAt *time.Time    `bson:"delete_at,omitempty"`
}

func (m *Model) Init() {
	now := time.Now()
	m.CreateAt = now
	m.UpdateAt = now
	m.DeleteAt = nil
}

func (m *Model) Collection() string {
	return "models"
}

func (m *Model) SetID(a any) {
	if id, ok := a.(bson.ObjectID); ok {
		m.Id = id
	}
}

func (m *Model) ID() bson.ObjectID {
	return m.Id
}
