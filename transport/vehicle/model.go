package vehicle

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "transport_"
	collectionNameSuffix = "_config"
	modelName            = "vehicle"
)

// Model 自建运力车辆，存 m3s.transport_vehicle_config
type Model struct {
	model.Model `json:"_" bson:"_"`
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	Plate       string  `json:"plate" bson:"plate"`
	Kind        string  `json:"kind" bson:"kind"`
	Brand       string  `json:"brand" bson:"brand,omitempty"`
	Color       string  `json:"color" bson:"color,omitempty"`
	LoadKg      float64 `json:"load_kg" bson:"load_kg,omitempty"`
	VolumeM3    float64 `json:"volume_m3" bson:"volume_m3,omitempty"`
	DriverName  string  `json:"driver_name" bson:"driver_name,omitempty"`
	DriverPhone string  `json:"driver_phone" bson:"driver_phone,omitempty"`
	Status      string  `json:"status" bson:"status"`
	Enabled     bool    `json:"enabled" bson:"enabled"`
	Note        string  `json:"note" bson:"note,omitempty"`
}

func (m *Model) ResourceName() string {
	return modelName
}

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
