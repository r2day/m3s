package actor

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "transport_"
	collectionNameSuffix = "_config"
	modelName            = "actor"
)

// Shipper 货主联系方式。电话是履约号码，不是会员密文手机号。
type Shipper struct {
	ContactName  string `json:"contact_name,omitempty" bson:"contact_name,omitempty"`
	ContactPhone string `json:"contact_phone,omitempty" bson:"contact_phone,omitempty"`
}

// Driver 司机申请资料。Status 只由服务端写入。
type Driver struct {
	Status        string  `json:"status" bson:"status"`
	RealName      string  `json:"real_name,omitempty" bson:"real_name,omitempty"`
	Plate         string  `json:"plate,omitempty" bson:"plate,omitempty"`
	VehicleKind   string  `json:"vehicle_kind,omitempty" bson:"vehicle_kind,omitempty"`
	LoadKg        float64 `json:"load_kg,omitempty" bson:"load_kg,omitempty"`
	VolumeM3      float64 `json:"volume_m3,omitempty" bson:"volume_m3,omitempty"`
	LicenseFileID string  `json:"license_file_id,omitempty" bson:"license_file_id,omitempty"`
	IDFileID      string  `json:"id_file_id,omitempty" bson:"id_file_id,omitempty"`
	RejectReason  string  `json:"reject_reason,omitempty" bson:"reject_reason,omitempty"`
}

// Model 货主/司机档案，存 m3s.transport_actor_config。
// 一个商户下一个 account_id 一行。会员身份仍在 member_detail_data。
type Model struct {
	model.Model `json:"_" bson:"_"`
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	AccountID  string   `json:"account_id" bson:"account_id"`
	MemberID   string   `json:"member_id,omitempty" bson:"member_id,omitempty"`
	WxOpenID   string   `json:"-" bson:"wx_openid,omitempty"`
	Roles      []string `json:"roles" bson:"roles"`
	ActiveRole string   `json:"active_role,omitempty" bson:"active_role,omitempty"`
	Shipper    Shipper  `json:"shipper" bson:"shipper"`
	Driver     Driver   `json:"driver" bson:"driver"`
}

func (m *Model) ResourceName() string {
	return modelName
}

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
