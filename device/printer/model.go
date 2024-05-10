package printer

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "device_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_config"
	// 这个需要用户根据具体业务完成设定
	modelName = "printer"
)

// Model 打印机
type Model struct {
	// 模型继承
	model.Model `json:"_" bson:"_"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	// Name 名称
	Name string `json:"name" bson:"name,omitempty"`
	// 类型
	Type string `json:"type" bson:"type,omitempty"`
	// 商户配置
	PrinterConf Printer `json:"printer_conf" bson:"printer_conf,omitempty"`
	// 打印机实时状态
	// 在线查询api接口，因此不存储到数据库
	Status bool `json:"status" bson:"-"`
}

type Printer struct {
	Sn      string `json:"sn"  bson:"sn"`
	User    string `json:"user"  bson:"user"`
	UserKey string `json:"user_key"  bson:"user_key"`
	Debug   string `json:"debug"  bson:"debug"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	//m.Meta = m.GetMeta()
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
