package file

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *Model) IsExistContent(contextMD5 string) (*Model, error) {

	result := &Model{}
	coll := m.Context.Handler.Collection(m.Context.Collection)
	filter := bson.D{{Key: "content_md_5", Value: contextMD5}}

	// 获取数据列表
	cursor, err := coll.Find(m.Context.Context, filter)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if err = cursor.Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}
