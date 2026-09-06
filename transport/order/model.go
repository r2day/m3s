package order

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "transport_"
	collectionNameSuffix = "_order"
	modelName            = "takeout"
)

// Model 物流订单（UU 跑腿），存 m3s.transport_takeout_order
type Model struct {
	model.Model `json:"_" bson:"_"`
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	OriginID       string `json:"origin_id" bson:"origin_id"`
	OrderCode      string `json:"order_code" bson:"order_code"`
	PriceToken     string `json:"price_token" bson:"price_token,omitempty"`
	OrderPrice     string `json:"order_price" bson:"order_price,omitempty"`
	NeedPaymoney   string `json:"need_paymoney" bson:"need_paymoney,omitempty"`
	Receiver       string `json:"receiver" bson:"receiver,omitempty"`
	ReceiverPhone  string `json:"receiver_phone" bson:"receiver_phone,omitempty"`
	Note           string `json:"note" bson:"note,omitempty"`
	State          string `json:"state" bson:"state,omitempty"`
	FromAddress    string `json:"from_address" bson:"from_address,omitempty"`
	FromLat        string `json:"from_lat" bson:"from_lat,omitempty"`
	FromLng        string `json:"from_lng" bson:"from_lng,omitempty"`
	ToAddress      string `json:"to_address" bson:"to_address,omitempty"`
	ToLat          string `json:"to_lat" bson:"to_lat,omitempty"`
	ToLng          string `json:"to_lng" bson:"to_lng,omitempty"`
	Distance       string `json:"distance" bson:"distance,omitempty"`
	DriverName     string `json:"driver_name" bson:"driver_name,omitempty"`
	DriverMobile   string `json:"driver_mobile" bson:"driver_mobile,omitempty"`
	DriverJobnum   string `json:"driver_jobnum" bson:"driver_jobnum,omitempty"`
	DriverPhoto    string `json:"driver_photo" bson:"driver_photo,omitempty"`
	AddTime        string `json:"add_time" bson:"add_time,omitempty"`
	FinishTime     string `json:"finish_time" bson:"finish_time,omitempty"`
	RoborderTime   string `json:"roborder_time" bson:"roborder_time,omitempty"`
	ExpectedArrive string `json:"expectedarrive_time" bson:"expectedarrive_time,omitempty"`
	ReturnCode     string `json:"return_code" bson:"return_code,omitempty"`
	ReturnMsg      string `json:"return_msg" bson:"return_msg,omitempty"`
	Provider       string `json:"provider" bson:"provider,omitempty"`
	MerchantID     string `json:"merchant_id" bson:"merchant_id,omitempty"`
}

func (m *Model) ResourceName() string {
	return modelName
}

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
