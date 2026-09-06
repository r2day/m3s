package order

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *Model) GetByOriginID(originID string) error {
	if originID == "" {
		return mongo.ErrNoDocuments
	}
	coll := m.Context.Handler.Collection(m.Context.Collection)
	err := coll.FindOne(m.Context.Context, bson.M{"origin_id": originID}).Decode(m)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}
	return err
}

func (m *Model) GetByOrderCode(orderCode string) error {
	if orderCode == "" {
		return mongo.ErrNoDocuments
	}
	coll := m.Context.Handler.Collection(m.Context.Collection)
	err := coll.FindOne(m.Context.Context, bson.M{"order_code": orderCode}).Decode(m)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}
	return err
}
