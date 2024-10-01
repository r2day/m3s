package notification

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "client_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_config"
	// 这个需要用户根据具体业务完成设定
	modelName = "home"
)

// AppType 小程序类型
type AppType int

const (
	// WxApp 微信小程序
	WxApp AppType = iota
	// AlipayApp 支付宝
	AlipayApp
	// DouYin 抖音小程序
	DouYin
)

// Model 小程序主页的配置
type Model struct {
	// 模型继承
	model.Model `json:"_" bson:"_"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	// Name 名称
	Name string `json:"name" bson:"name,omitempty"`
	// 小程序类型
	Type AppType `json:"type" bson:"type,omitempty"`
	// logo
	Logo string `json:"logo" bson:"logo,omitempty"`
	// Background 背景图片
	Background string `json:"background" bson:"background,omitempty"`
	// 下单方式按钮 （例如：堂食、自提、外卖）
	Entrances []Entrance `json:"entrances" bson:"entrances,omitempty"`
	// 推荐区域的标题（例如：火热的，热销的，招牌的）
	RecommendTitle string `json:"recommend_title" bson:"recommend_title,omitempty"`
	// 推荐列表展示
	RecommendShowList []RecommendShow `json:"recommend_show_list" bson:"recommend_show_list,omitempty"`
	// Enable 启用与否
	Enable bool `json:"enable" bson:"enable,omitempty"`
}

// Entrance 使用方式如下:
// <view class="item" @tap="go2scan">
//
//	  <view class="title">堂食2</view>
//	  <view class="sub-title">Tang Shi</view>
//	  <view class="sub-name">扫描餐桌二维码点单</view>
//	  <image src="/static/images/hlj/3-1.png" class="icon"></image>
//	</view>
type Entrance struct {
	Tap      string `json:"tap" bson:"tap,omitempty"`
	Title    string `json:"title" bson:"title,omitempty"`
	SubTitle string `json:"sub_title" bson:"sub_title,omitempty"`
	SubName  string `json:"sub_name" bson:"sub_name,omitempty"`
	Icon     string `json:"icon" bson:"icon,omitempty"`
}

// RecommendShow 招牌推荐区域
//
//	 <view class="list" v-for="(item, index) in recommendShow" :key="index">
//	  <view class="item" @tap="clickAndRedirect(item.url)">
//	    <image mode="widthFix" :src="item.image_url"></image>
//	  </view>
//	</view>
type RecommendShow struct {
	Title string `json:"title" bson:"title,omitempty"`
	Url   string `json:"url" bson:"url,omitempty"`
	Image string `json:"image_url" bson:"image_url,omitempty"`
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
