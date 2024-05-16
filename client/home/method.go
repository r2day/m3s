package notification

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *Model) GetByStoreID(id string) ([]*Model, error) {

	results := make([]*Model, 0)
	coll := m.Context.Handler.Collection(m.Context.Collection)
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "meta.merchant_id", Value: objID}}

	// 获取数据列表
	cursor, err := coll.Find(m.Context.Context, filter)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if err = cursor.All(m.Context.Context, &results); err != nil {
		return nil, err
	}
	return results, nil
}
