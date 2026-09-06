package platform

import (
	"strings"

	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "transport_"
	collectionNameSuffix = "_config"
	modelName            = "platform"
)

// Model 业主绑定的外卖/配送平台密钥，存 m3s.transport_platform_config。
// 每个租户每个 provider 一条。
type Model struct {
	model.Model `json:"_" bson:"_"`
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	Provider string `json:"provider" bson:"provider"`
	Name     string `json:"name" bson:"name,omitempty"`
	Kind     string `json:"kind" bson:"kind,omitempty"`

	AppID    string `json:"app_id" bson:"app_id,omitempty"`
	AppKey   string `json:"app_key" bson:"app_key,omitempty"`
	Token    string `json:"token" bson:"token,omitempty"`
	ShopID   string `json:"shop_id" bson:"shop_id,omitempty"`
	Callback string `json:"callback" bson:"callback,omitempty"`
	APIURL   string `json:"api_url" bson:"api_url,omitempty"`
	City     string `json:"city" bson:"city,omitempty"`

	Enabled bool   `json:"enabled" bson:"enabled"`
	Note    string `json:"note" bson:"note,omitempty"`
}

func (m *Model) ResourceName() string {
	return modelName
}

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}

func (m *Model) Bound() bool {
	if m == nil {
		return false
	}
	return strings.TrimSpace(m.AppID) != "" || strings.TrimSpace(m.AppKey) != "" || strings.TrimSpace(m.Token) != ""
}
