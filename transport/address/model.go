package address

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "transport_"
	collectionNameSuffix = "_config"
	modelName            = "address"
)

// Model 商户发货 / 取件地址，存 m3s.transport_address_config。
// 会员收货地址仍走 member 模块，不在这里管。
type Model struct {
	model.Model `json:"_" bson:"_"`
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	AcceptName string `json:"accept_name" bson:"accept_name"`
	Mobile     string `json:"mobile" bson:"mobile"`
	Address    string `json:"address" bson:"address"`
	Note       string `json:"note" bson:"note,omitempty"`
	Lat        string `json:"lat" bson:"lat,omitempty"`
	Lng        string `json:"lng" bson:"lng,omitempty"`
	Tag        string `json:"tag" bson:"tag,omitempty"`
	IsDefault  bool   `json:"is_default" bson:"is_default"`
	Enabled    bool   `json:"enabled" bson:"enabled"`
}

func (m *Model) ResourceName() string {
	return modelName
}

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
