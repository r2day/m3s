package subscribe

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "message_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_flow"
	// 这个需要用户根据具体业务完成设定
	modelName = "subscribe"
)

// MessageStatus represents the status of the message
type MessageStatus int

const (
	// Sent represents a message status of sent
	Sent MessageStatus = iota
	// Read represents a message status of read
	Read
	// Unread represents a message status of unread
	Unread
	// Deleted represents a message status of deleted
	Deleted
	// Add more statuses as needed
)

// MessageType represents the type of the message
type MessageType int

// Model 打印机
type Model struct {
	// 模型继承
	model.Model `json:"_" bson:"_"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	// Name 名称(默认系统: system）
	Operator string `json:"operator" bson:"operator,omitempty"`
	// 模版
	TemplateID string `json:"template_id" bson:"template_id,omitempty"`
	// 接受者
	Receiver string `json:"receiver" bson:"receiver,omitempty"`
	// 消息主体
	Payload MessagePayload `json:"payload" bson:"payload,omitempty"`
	// 消息状态
	Status MessageStatus `json:"message_status" bson:"message_status,omitempty"`
}

type MessagePayload struct {
	Thing5   string `json:"thing5" bson:"thing5"`
	Amount12 string `json:"amount12" bson:"amount12"`
	Thing7   string `json:"thing7" bson:"thing7"`
	Date4    string `json:"date4" bson:"date4"`
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
