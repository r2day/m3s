package freight

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "transport_"
	collectionNameSuffix = "_order"
	modelName            = "freight"
)

// Model x11 货运单（货主发布、司机报价、线下付款），存 m3s.transport_freight_order。
// 与 UU 跑腿单 transport/order 分开：状态是字符串（draft/published/matched/in_transit/delivered/cancelled），
// 报价内嵌在 applications，付款状态在 pay。金额同时存「元」字符串和「分」int64，不用 float。
type Model struct {
	model.Model `json:"_" bson:"_"`
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	FreightNo string `json:"freight_no" bson:"freight_no"`
	// IdempotencyKey 发布幂等键。为空时不落库，配合 partial 唯一索引。
	IdempotencyKey string `json:"idempotency_key,omitempty" bson:"idempotency_key,omitempty"`
	Status         string `json:"status" bson:"status"`

	ShipperAccountID string `json:"shipper_account_id" bson:"shipper_account_id"`
	ShipperMemberID  string `json:"shipper_member_id" bson:"shipper_member_id,omitempty"`
	DriverAccountID  string `json:"driver_account_id" bson:"driver_account_id,omitempty"`
	DriverMemberID   string `json:"driver_member_id" bson:"driver_member_id,omitempty"`

	Origin      Place `json:"origin" bson:"origin"`
	Destination Place `json:"destination" bson:"destination"`

	CargoName   string  `json:"cargo_name" bson:"cargo_name"`
	CargoType   string  `json:"cargo_type" bson:"cargo_type"`
	WeightKg    float64 `json:"weight_kg" bson:"weight_kg"`
	VolumeM3    float64 `json:"volume_m3" bson:"volume_m3"`
	Count       int     `json:"count" bson:"count"`
	VehicleKind string  `json:"vehicle_kind" bson:"vehicle_kind"`
	// PickupFrom / PickupTo 装货时间窗，Unix 秒。
	PickupFrom int64 `json:"pickup_from" bson:"pickup_from"`
	PickupTo   int64 `json:"pickup_to" bson:"pickup_to"`

	// PriceMode fixed / negotiable。PriceYuan 是十进制字符串，PriceFen 是对应的分。
	PriceMode string `json:"price_mode" bson:"price_mode"`
	PriceYuan string `json:"price_yuan" bson:"price_yuan"`
	PriceFen  int64  `json:"price_fen" bson:"price_fen"`
	// AgreedYuan / AgreedFen 接受报价后的成交价。
	AgreedYuan string `json:"agreed_yuan" bson:"agreed_yuan"`
	AgreedFen  int64  `json:"agreed_fen" bson:"agreed_fen"`
	// PayTiming on_pickup / on_delivery，决定哪个状态可以确认付款。
	PayTiming string `json:"pay_timing" bson:"pay_timing"`

	Remark         string   `json:"remark" bson:"remark"`
	PhotoIDs       []string `json:"photo_ids" bson:"photo_ids"`
	PodPhotoIDs    []string `json:"pod_photo_ids" bson:"pod_photo_ids,omitempty"`
	ContactVisible string   `json:"contact_visible" bson:"contact_visible"`

	Applications []Application `json:"applications" bson:"applications"`
	Pay          Pay           `json:"pay" bson:"pay"`

	ShipperReviewID string `json:"shipper_review_id" bson:"shipper_review_id,omitempty"`
	DriverReviewID  string `json:"driver_review_id" bson:"driver_review_id,omitempty"`

	CancelBy     string `json:"cancel_by" bson:"cancel_by,omitempty"`
	CancelReason string `json:"cancel_reason" bson:"cancel_reason,omitempty"`

	// CreatedAt 创建时间 Unix 秒（顶层字段，供索引排序；_.meta 里也有一份）。
	CreatedAt int64 `json:"created_at" bson:"created_at"`
	// Provider 固定 x11，避免与 UU 单的 uu/self 混淆。
	Provider string `json:"provider" bson:"provider,omitempty"`
}

// Place 装货地或卸货地快照。坐标存字符串，不落 float。
type Place struct {
	Name    string `json:"name" bson:"name"`
	Phone   string `json:"phone" bson:"phone"`
	Region  string `json:"region" bson:"region"`
	Address string `json:"address" bson:"address"`
	Door    string `json:"door" bson:"door"`
	Lat     string `json:"lat" bson:"lat"`
	Lng     string `json:"lng" bson:"lng"`
	City    string `json:"city" bson:"city"`
}

// Application 司机报价，内嵌在运单 applications[]。
// Status 为 pending/accepted/rejected/withdrawn/cancelled。ID 是 ObjectID hex 字符串。
type Application struct {
	ID        string `json:"id" bson:"id"`
	AccountID string `json:"account_id" bson:"account_id"`
	MemberID  string `json:"member_id" bson:"member_id,omitempty"`
	QuoteYuan string `json:"quote_yuan" bson:"quote_yuan"`
	QuoteFen  int64  `json:"quote_fen" bson:"quote_fen"`
	Note      string `json:"note" bson:"note,omitempty"`
	Status    string `json:"status" bson:"status"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
}

// Pay 线下付款状态。Status 为 unpaid/settled/disputed；Ledger 为 none/pending/done，与 Status 分开。
// 时间字段都是 Unix 秒，0 表示还没发生。
type Pay struct {
	Status             string   `json:"status" bson:"status"`
	Timing             string   `json:"timing" bson:"timing"`
	Ledger             string   `json:"ledger" bson:"ledger"`
	EvidenceFileIDs    []string `json:"evidence_file_ids,omitempty" bson:"evidence_file_ids,omitempty"`
	EvidenceAt         int64    `json:"evidence_at,omitempty" bson:"evidence_at,omitempty"`
	ShipperConfirmedAt int64    `json:"shipper_confirmed_at,omitempty" bson:"shipper_confirmed_at,omitempty"`
	DriverConfirmedAt  int64    `json:"driver_confirmed_at,omitempty" bson:"driver_confirmed_at,omitempty"`
	TransactionID      string   `json:"transaction_id,omitempty" bson:"transaction_id,omitempty"`
	RetryStopped       bool     `json:"retry_stopped,omitempty" bson:"retry_stopped,omitempty"`
	DisputeReason      string   `json:"dispute_reason,omitempty" bson:"dispute_reason,omitempty"`
	DisputeNote        string   `json:"dispute_note,omitempty" bson:"dispute_note,omitempty"`
}

func (m *Model) ResourceName() string {
	return modelName
}

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
