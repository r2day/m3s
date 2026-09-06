package route

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "transport_"
	collectionNameSuffix = "_config"
	modelName            = "route"
)

// Model 线路。专线=固定班线；非专线=临时/同城散单。
type Model struct {
	model.Model `json:"_" bson:"_"`
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	Name          string `json:"name" bson:"name"`
	FromCity      string `json:"from_city" bson:"from_city,omitempty"`
	ToCity        string `json:"to_city" bson:"to_city,omitempty"`
	FromAddress   string `json:"from_address" bson:"from_address,omitempty"`
	ToAddress     string `json:"to_address" bson:"to_address,omitempty"`
	Dedicated     bool   `json:"dedicated" bson:"dedicated"`
	TwoWay        bool   `json:"two_way" bson:"two_way"`
	FleetID       string `json:"fleet_id" bson:"fleet_id,omitempty"`
	FleetName     string `json:"fleet_name" bson:"fleet_name,omitempty"`
	Fee           string `json:"fee" bson:"fee,omitempty"`
	BillingMethod string `json:"billing_method" bson:"billing_method,omitempty"`
	Enabled       bool   `json:"enabled" bson:"enabled"`
	Note          string `json:"note" bson:"note,omitempty"`
}

func (m *Model) ResourceName() string {
	return modelName
}

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
