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

// Place 装卸货快照。电话单独存放，不拼进 address。
type Place struct {
	Name    string `json:"name" bson:"name"`
	Phone   string `json:"phone,omitempty" bson:"phone,omitempty"`
	Region  string `json:"region,omitempty" bson:"region,omitempty"`
	Address string `json:"address" bson:"address"`
	Door    string `json:"door,omitempty" bson:"door,omitempty"`
	Lat     string `json:"lat,omitempty" bson:"lat,omitempty"`
	Lng     string `json:"lng,omitempty" bson:"lng,omitempty"`
	City    string `json:"city" bson:"city"`
}

// Application 内嵌报价。不另建集合。
type Application struct {
	ID        string `json:"id" bson:"id"`
	AccountID string `json:"account_id" bson:"account_id"`
	MemberID  string `json:"member_id,omitempty" bson:"member_id,omitempty"`
	QuoteYuan string `json:"quote_yuan" bson:"quote_yuan"`
	QuoteFen  int64  `json:"quote_fen" bson:"quote_fen"`
	Note      string `json:"note,omitempty" bson:"note,omitempty"`
	Status    string `json:"status" bson:"status"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
}

// Pay 线下付款。Status 与 Ledger 分开。没有 evidence 这一档状态，也没有 version。
type Pay struct {
	Status             string   `json:"status" bson:"status"`
	Timing             string   `json:"timing" bson:"timing"`
	EvidenceFileIDs    []string `json:"evidence_file_ids,omitempty" bson:"evidence_file_ids,omitempty"`
	EvidenceAt         int64    `json:"evidence_at,omitempty" bson:"evidence_at,omitempty"`
	ShipperConfirmedAt int64    `json:"shipper_confirmed_at,omitempty" bson:"shipper_confirmed_at,omitempty"`
	DriverConfirmedAt  int64    `json:"driver_confirmed_at,omitempty" bson:"driver_confirmed_at,omitempty"`
	TransactionID      string   `json:"transaction_id,omitempty" bson:"transaction_id,omitempty"`
	Ledger             string   `json:"ledger" bson:"ledger"`
	RetryStopped       bool     `json:"retry_stopped" bson:"retry_stopped"`
	DisputeReason      string   `json:"dispute_reason,omitempty" bson:"dispute_reason,omitempty"`
	DisputeNote        string   `json:"dispute_note,omitempty" bson:"dispute_note,omitempty"`
}

// Model 货主运单，存 m3s.transport_freight_order。不是餐饮订单，也不是 UU 跑腿单。
type Model struct {
	model.Model `json:"_" bson:"_"`
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	FreightNo        string `json:"freight_no" bson:"freight_no"`
	IdempotencyKey   string `json:"idempotency_key,omitempty" bson:"idempotency_key,omitempty"`
	Status           string `json:"status" bson:"status"`
	ShipperAccountID string `json:"shipper_account_id" bson:"shipper_account_id"`
	ShipperMemberID  string `json:"shipper_member_id,omitempty" bson:"shipper_member_id,omitempty"`
	DriverAccountID  string `json:"driver_account_id,omitempty" bson:"driver_account_id,omitempty"`
	DriverMemberID   string `json:"driver_member_id,omitempty" bson:"driver_member_id,omitempty"`

	Origin      Place `json:"origin" bson:"origin"`
	Destination Place `json:"destination" bson:"destination"`

	CargoName   string  `json:"cargo_name" bson:"cargo_name"`
	CargoType   string  `json:"cargo_type" bson:"cargo_type"`
	WeightKg    float64 `json:"weight_kg" bson:"weight_kg"`
	VolumeM3    float64 `json:"volume_m3" bson:"volume_m3"`
	Count       int     `json:"count" bson:"count"`
	VehicleKind string  `json:"vehicle_kind" bson:"vehicle_kind"`

	PickupFrom int64 `json:"pickup_from" bson:"pickup_from"`
	PickupTo   int64 `json:"pickup_to" bson:"pickup_to"`

	PriceMode  string `json:"price_mode" bson:"price_mode"`
	PriceYuan  string `json:"price_yuan,omitempty" bson:"price_yuan,omitempty"`
	PriceFen   int64  `json:"price_fen,omitempty" bson:"price_fen,omitempty"`
	AgreedYuan string `json:"agreed_yuan,omitempty" bson:"agreed_yuan,omitempty"`
	AgreedFen  int64  `json:"agreed_fen,omitempty" bson:"agreed_fen,omitempty"`
	PayTiming  string `json:"pay_timing" bson:"pay_timing"`

	Remark         string   `json:"remark,omitempty" bson:"remark,omitempty"`
	PhotoIDs       []string `json:"photo_ids,omitempty" bson:"photo_ids,omitempty"`
	PodPhotoIDs    []string `json:"pod_photo_ids,omitempty" bson:"pod_photo_ids,omitempty"`
	ContactVisible string   `json:"contact_visible" bson:"contact_visible"`

	Applications    []Application `json:"applications" bson:"applications"`
	ShipperReviewID string        `json:"shipper_review_id,omitempty" bson:"shipper_review_id,omitempty"`
	DriverReviewID  string        `json:"driver_review_id,omitempty" bson:"driver_review_id,omitempty"`

	Pay          Pay    `json:"pay" bson:"pay"`
	CancelReason string `json:"cancel_reason,omitempty" bson:"cancel_reason,omitempty"`
	CancelBy     string `json:"cancel_by,omitempty" bson:"cancel_by,omitempty"`

	// CreatedAt 是货盘排序用的 Unix 秒。商户与创建人在 _.meta。
	CreatedAt int64  `json:"created_at" bson:"created_at"`
	Provider  string `json:"provider,omitempty" bson:"provider,omitempty"`
}

func (m *Model) ResourceName() string {
	return modelName
}

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
