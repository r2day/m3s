package fleet

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "transport_"
	collectionNameSuffix = "_config"
	modelName            = "fleet"
)

// Model 车队，存 m3s.transport_fleet_config
type Model struct {
	model.Model `json:"_" bson:"_"`
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	Name    string `json:"name" bson:"name"`
	Manager string `json:"manager" bson:"manager,omitempty"`
	Phone   string `json:"phone" bson:"phone,omitempty"`
	City    string `json:"city" bson:"city,omitempty"`
	Enabled bool   `json:"enabled" bson:"enabled"`
	Note    string `json:"note" bson:"note,omitempty"`
}

func (m *Model) ResourceName() string {
	return modelName
}

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
