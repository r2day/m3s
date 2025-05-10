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
	// Background 背景颜色
	Background string `json:"background" bson:"background,omitempty"`
	// 推荐区域的标题（例如：火热的，热销的，招牌的）
	RecommendTitle string `json:"recommend_title" bson:"recommend_title,omitempty"`
	// 推荐列表展示
	RecommendShowList []*ImageDisplayConfig `json:"recommend_show_list" bson:"recommend_show_list,omitempty"`
	// Enable 启用与否
	Enable bool `json:"enable" bson:"enable,omitempty"`

	// Base 基本图标
	Base BaseIcon `json:"base" bson:"base,omitempty"`
	// 底部导航
	Tabs []string `json:"tabs" bson:"tabs,omitempty"`
	// HomeConfig 主页配置
	HomeConfig HomePageConfig `json:"home_config" bson:"home_config,omitempty"`
	// MenuConfig 新方案
	MenuConfig MenuPageConfig `json:"menu_config" bson:"menu_config,omitempty"`
	// OrderConfig 新方案
	OrderConfig OrderPageConfig `json:"order_config" bson:"order_config,omitempty"`
	// ProfileConfig 新方案
	ProfileConfig ProfilePageConfig `json:"profile_config" bson:"profile_config,omitempty"`
}

type BaseIcon struct {
	// 加载图标
	Loading string `json:"loading" bson:"loading,omitempty"`
	// 奖杯
	Cup string `json:"cup" bson:"cup,omitempty"`
	// Avatar 头像
	Avatar string `json:"avatar" bson:"avatar,omitempty"`
	// Role 吉祥物角色
	Role string `json:"role" bson:"role,omitempty"`
	// Level 等级图标，例如vip
	Level string `json:"level" bson:"level,omitempty"`
	// QRCode 二维码
	QRCode string `json:"qr_code" bson:"qr_code,omitempty"`
	// PreButton 向上箭头
	PreButton string `json:"pre_button" bson:"pre_button,omitempty"`
	// NextButton 向下箭头
	NextButton string `json:"next_button" bson:"next_button,omitempty"`
	// Service 客服图标
	Service string `json:"service" bson:"service,omitempty"`
	// Location 定位/地址图标
	Location string `json:"location" bson:"location,omitempty"`
	// Order 订单图标
	Order string `json:"order" bson:"order,omitempty"`
	// Packages 圈包图标
	Packages string `json:"packages" bson:"packages,omitempty"`
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
	// 操作的函数
	Tap string `json:"tap" bson:"tap,omitempty"`
	// 主标题
	Title string `json:"title" bson:"title,omitempty"`
	// 子标题
	SubTitle string `json:"sub_title" bson:"sub_title,omitempty"`
	// 英文名称
	SubName string `json:"sub_name" bson:"sub_name,omitempty"`
	// 图标
	Icon string `json:"icon" bson:"icon,omitempty"`
	// 背景图片或者颜色
	Background string `json:"background" bson:"background,omitempty"`
}

// ImageDisplayConfig 招牌推荐区域
//
//	 <view class="list" v-for="(item, index) in recommendShow" :key="index">
//	  <view class="item" @tap="clickAndRedirect(item.url)">
//	    <image mode="widthFix" :src="item.image_url"></image>
//	  </view>
//	</view>
type ImageDisplayConfig struct {
	// 标题
	Title string `json:"title" bson:"title,omitempty"`
	// 图片地址
	Image string `json:"image_url" bson:"image_url,omitempty"`
	// PackageType 包类型：main,public, private
	PackageType string `json:"package_type" bson:"package_type,omitempty"`
	// 根据包类型推荐不同的url
	Url string `json:"url" bson:"url,omitempty"`
	// 图片停留时长
	Stay int `json:"stay" bson:"stay,omitempty"`
	// 图片类型
	Type int `json:"type" bson:"type,omitempty"`
}

// HomePageConfig 主页页面配置管理
type HomePageConfig struct {
	// 轮播背景图片展示
	Background []*ImageDisplayConfig `json:"background" bson:"background,omitempty"`
	// Greeting 问候语，例如：Hello, Frank
	Greeting string `json:"greeting" bson:"greeting,omitempty"`
	// Remark 备注信息
	Remark string `json:"remark" bson:"remark,omitempty"`
	// ShowMerchantName 是否展示门店名称
	ShowMerchantName bool `json:"show_merchant_name" bson:"show_merchant_name,omitempty"`
	// 下单方式按钮 （例如：堂食、自提、外卖）
	Entrances []*Entrance `json:"entrances" bson:"entrances,omitempty"`
	// 推荐列表展示
	RecommendShowList []*ImageDisplayConfig `json:"recommend_show_list" bson:"recommend_show_list,omitempty"`
	// 快速导航
	Navigators []*Entrance `json:"navigators" bson:"navigators,omitempty"`
}

type OrderPageConfig struct {
	// ShowMerchantName 是否展示门店名称
	ShowMerchantName bool `json:"show_merchant_name" bson:"show_merchant_name,omitempty"`
	// ShowNavigateName 是否展示门店定位
	ShowNavigate bool `json:"show_navigate" bson:"show_navigate,omitempty"`
	// ShowPhone 是否展示门店定位
	ShowPhone bool `json:"show_phone" bson:"show_phone,omitempty"`
	// OrderProcess 订单状态进度
	OrderProcess OrderProcessConfig `json:"order_process" bson:"order_process,omitempty"`
}

// OrderProcessConfig 订单状态图标设置
type OrderProcessConfig struct {
	// OrderInit 订单初始化
	OrderInit string `json:"order_init" bson:"order_init,omitempty"`
	// OrderInitDone 订单初始化
	OrderInitDone string `json:"order_init_done" bson:"order_init_done,omitempty"`
	// OrderPay 订单支付
	OrderPay string `json:"order_pay" bson:"order_pay,omitempty"`
	// OrderPaid 订单支付完毕
	OrderPaid string `json:"order_paid" bson:"order_paid,omitempty"`
	// OrderReadyTake 订单待取餐
	OrderReadyTake string `json:"order_ready_take" bson:"order_ready_take,omitempty"`
	// OrderDone 订单完毕
	OrderDone string `json:"order_done" bson:"order_done,omitempty"`
}

type ProfilePageConfig struct {
	// Service 服务设置
	Service []ImageDisplayConfig `json:"service" bson:"service,omitempty"`
}

// MenuPageConfig 菜单页面配置管理
type MenuPageConfig struct {
	// ShowMerchantName 是否展示门店名称
	ShowMerchantName bool `json:"show_merchant_name" bson:"show_merchant_name,omitempty"`
	// ShowPickUpSwitch 是否展示自提/外卖切换按钮
	ShowPickUpSwitch bool `json:"show_pick_up_switch" bson:"show_pick_up_switch,omitempty"`
	// ShowPickUpSwitch 是否展示距离门店的距离
	ShowDistance bool `json:"show_distance" bson:"show_distance,omitempty"`
	// ShowCategory 是否展示左侧类型选择
	ShowCategory bool `json:"show_category" bson:"show_category,omitempty"`
	// 是否展示跑马灯
	ShowMarquee bool `json:"show_marquee" bson:"show_marquee,omitempty"`
	// 标题 home, menu, order, mine
	Title string `json:"title" bson:"title,omitempty"`
	// 轮播图片展示
	Carousel []*ImageDisplayConfig `json:"carousel" bson:"carousel,omitempty"`
	// 购物车图标
	CartIcon string `json:"cart_icon" bson:"cart_icon,omitempty"`
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
