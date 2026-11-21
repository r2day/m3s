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

// Model x11 货运小程序的货主/司机档案，存 m3s.transport_actor_config。
// 一个商户（_.meta.merchant_id）下一个 account_id 一行；roles 是集合，一人可同时是货主和司机。
type Model struct {
	model.Model `json:"_" bson:"_"`
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	// AccountID 会员登录账号，等于 JWT sub。服务端写入。
	AccountID string `json:"account_id" bson:"account_id"`
	// MemberID member_detail_data._id 的 hex。服务端写入。
	MemberID string `json:"member_id" bson:"member_id,omitempty"`
	// WxOpenID 最近一次请求的 JWT jti。只存库，不出现在 JSON 里。
	WxOpenID string `json:"-" bson:"wx_openid,omitempty"`

	// Roles shipper / driver 的子集，只能由服务端修改。
	Roles []string `json:"roles" bson:"roles"`
	// ActiveRole 小程序当前展示的身份，仅用于 UI，鉴权看 Roles。
	ActiveRole string `json:"active_role" bson:"active_role,omitempty"`

	Shipper Shipper `json:"shipper" bson:"shipper"`
	Driver  Driver  `json:"driver" bson:"driver"`
}

// Shipper 货主联系信息，履约用。
type Shipper struct {
	ContactName  string `json:"contact_name" bson:"contact_name,omitempty"`
	ContactPhone string `json:"contact_phone" bson:"contact_phone,omitempty"`
}

// Driver 司机资料。Status 为 none/pending/active/rejected，只由服务端和员工审核接口修改。
type Driver struct {
	Status        string  `json:"status" bson:"status"`
	RealName      string  `json:"real_name" bson:"real_name,omitempty"`
	Plate         string  `json:"plate" bson:"plate,omitempty"`
	VehicleKind   string  `json:"vehicle_kind" bson:"vehicle_kind,omitempty"`
	LoadKg        float64 `json:"load_kg" bson:"load_kg,omitempty"`
	VolumeM3      float64 `json:"volume_m3" bson:"volume_m3,omitempty"`
	LicenseFileID string  `json:"license_file_id" bson:"license_file_id,omitempty"`
	IDFileID      string  `json:"id_file_id" bson:"id_file_id,omitempty"`
	RejectReason  string  `json:"reject_reason,omitempty" bson:"reject_reason,omitempty"`
}

func (m *Model) ResourceName() string {
	return modelName
}

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
