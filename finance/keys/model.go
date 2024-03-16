package keys

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "finance_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_config"
	// 这个需要用户根据具体业务完成设定
	modelName = "keys"
)

// Model 商品信息
type Model struct {
	// 模型继承
	model.Model `json:"meta" bson:"-"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	// Name 密钥名称
	Name string `json:"name" bson:"name,omitempty"`
	// 密钥
	Private string `json:"private" bson:"private,omitempty"`
	// 密钥类型
	Type string `json:"type" bson:"type,omitempty"`
	// 商户配置
	MerchantConf Merchant `json:"merchant_conf" bson:"merchant_conf,omitempty"`
	// Enabled 是否启用
	Enabled bool `json:"enabled" bson:"enabled,omitempty"`
}

type Merchant struct {
	// ID 编号
	ID string `json:"merchant_id" bson:"merchant_id,omitempty"`
	// 序列号
	CertSN string `json:"merchant_cert_sn" bson:"merchant_cert_sn,omitempty"`
	// 接口key
	APIKey string `json:"merchant_api_key" bson:"merchant_api_key,omitempty"`
	// AppID 应用id
	AppID string `json:"app_id" bson:"app_id,omitempty"`
	// Callback 回调地址
	Callback string `json:"callback" bson:"callback,omitempty"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	m.Meta = m.GetMeta()
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
