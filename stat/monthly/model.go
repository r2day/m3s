package monthly

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "stat_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_data"
	// 这个需要用户根据具体业务完成设定
	modelName = "monthly"
)

// MessageType represents the type of the message
type MessageType int

// Model 打印机
type Model struct {
	// 模型继承
	model.Model `json:"_" bson:"_"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	YearMonth     string             `bson:"year_month" json:"year_month"` // 格式: YYYY-MM
	OrderCount    int                `bson:"order_count" json:"order_count"`
	OrderAmount   float64            `bson:"order_amount" json:"order_amount"`
	RefundCount   int                `bson:"refund_count" json:"refund_count"`
	RefundAmount  float64            `bson:"refund_amount" json:"refund_amount"`
	CustomerCount int                `bson:"customer_count" json:"customer_count"` // 消费人次
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
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
