package file

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "f5s_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_flow"
	// 这个需要用户根据具体业务完成设定
	modelName = "file"
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
	// Format 格式 png; jpg; ..
	Format string `json:"format" bson:"format,omitempty"`
	// 大小 (单位byte)
	Size uint `json:"size" bson:"size,omitempty"`
	// 名称 md5值
	NameMD5 string `json:"name_md_5" bson:"name_md_5,omitempty"`
	// 内容 md5值 （防止重复上传）
	ContentMD5 string `json:"content_md_5" bson:"content_md_5,omitempty"`
	// 存储位置
	Path string `json:"path" bson:"path,omitempty"`
	// url地址
	Url string `json:"url" bson:"url,omitempty"`
	// 宽
	Width int `json:"width" bson:"width,omitempty"`
	// 高
	Height int `json:"height" bson:"height,omitempty"`
	// 请求访问次数
	Visits int `json:"visits" bson:"visits,omitempty"`
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
