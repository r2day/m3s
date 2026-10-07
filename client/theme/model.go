package theme

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Status draft | published
type Status string

// Kind preset | clone
type Kind string

// Tokens 十个色值。页面底色只活在 background，不进 launch 根 background。
type Tokens struct {
	Primary     string `json:"primary" bson:"primary,omitempty"`
	Secondary   string `json:"secondary" bson:"secondary,omitempty"`
	Background  string `json:"background" bson:"background,omitempty"`
	Price       string `json:"price" bson:"price,omitempty"`
	TabSelected string `json:"tab_selected" bson:"tab_selected,omitempty"`
	TabNormal   string `json:"tab_normal" bson:"tab_normal,omitempty"`
	Button      string `json:"button" bson:"button,omitempty"`
	ButtonText  string `json:"button_text" bson:"button_text,omitempty"`
	Text        string `json:"text" bson:"text,omitempty"`
	TextMuted   string `json:"text_muted" bson:"text_muted,omitempty"`
}

// Asset 槽位上的一张图。URL 存相对路径。FileID 是 Mongo _id 的 hex，不出现在公开 URL 里。
type Asset struct {
	Slot     string `json:"slot" bson:"slot,omitempty"`
	FileID   string `json:"file_id" bson:"file_id,omitempty"`
	PublicID string `json:"public_id" bson:"public_id,omitempty"` // 32 hex，URL 用这个
	URL      string `json:"url" bson:"url,omitempty"`             // /v1/system/fs/client/image/{public_id}
	Width    int    `json:"width" bson:"width,omitempty"`
	Height   int    `json:"height" bson:"height,omitempty"`
	Bytes    int    `json:"bytes" bson:"bytes,omitempty"`
}

// EntranceIcon 只换 icon。没有 background 字段，避免空值清掉手写的入口底。
type EntranceIcon struct {
	Tap  string `json:"tap" bson:"tap,omitempty"` // go2scan | takein | takeout
	Icon string `json:"icon" bson:"icon,omitempty"`
}

// Slide 轮播一条。
type Slide struct {
	ImageURL string `json:"image_url" bson:"image_url,omitempty"`
	Type     int    `json:"type" bson:"type,omitempty"`
	Stay     int    `json:"stay" bson:"stay,omitempty"`
}

// HomeSkin 首页皮肤切片。不含 popups、navigators、entrance.background。
type HomeSkin struct {
	Background []Slide        `json:"background,omitempty" bson:"background,omitempty"`
	Entrances  []EntranceIcon `json:"entrances,omitempty" bson:"entrances,omitempty"`
}

// LoginButtonSkin 只写 color。不包含 background、icon。
type LoginButtonSkin struct {
	Color string `json:"color,omitempty" bson:"color,omitempty"`
}

// LoginBenefits 是登录页底部五个权益图的 URL。
// 顺序与页面从左到右一致。不要写进 login.bottom。
type LoginBenefits struct {
	Join     string `json:"join,omitempty" bson:"join,omitempty"`         // 入会有礼
	Redeem   string `json:"redeem,omitempty" bson:"redeem,omitempty"`     // 积分兑换
	Service  string `json:"service,omitempty" bson:"service,omitempty"`   // 客服咨询
	Birthday string `json:"birthday,omitempty" bson:"birthday,omitempty"` // 生日特权
	Mall     string `json:"mall,omitempty" bson:"mall,omitempty"`         // 积分商城
}

// LoginSkin 登录页皮肤。Background 是字标 URL，写入 login.background。
type LoginSkin struct {
	Background string          `json:"background,omitempty" bson:"background,omitempty"`
	Login      LoginButtonSkin `json:"login,omitempty" bson:"login,omitempty"`
	Benefits   LoginBenefits   `json:"benefits,omitempty" bson:"benefits,omitempty"`
}

// MenuSkin 菜单页皮肤切片。
type MenuSkin struct {
	Carousel []Slide `json:"carousel,omitempty" bson:"carousel,omitempty"`
	CartIcon string  `json:"cart_icon,omitempty" bson:"cart_icon,omitempty"`
}

// OrderProcessSkin 只含 order_init 与 order_done。
// 不含 order_pay、order_paid、order_ready_take、order_init_done。
type OrderProcessSkin struct {
	OrderInit string `json:"order_init,omitempty" bson:"order_init,omitempty"`
	OrderDone string `json:"order_done,omitempty" bson:"order_done,omitempty"`
}

// OrderSkin 订单页皮肤切片。
type OrderSkin struct {
	OrderProcess OrderProcessSkin `json:"order_process,omitempty" bson:"order_process,omitempty"`
}

// ProfileSkin 我的页顶图。
type ProfileSkin struct {
	Cover string `json:"cover,omitempty" bson:"cover,omitempty"`
}

// BaseSkin 不含 cup。
type BaseSkin struct {
	Location   string `json:"location,omitempty" bson:"location,omitempty"`
	NextButton string `json:"next_button,omitempty" bson:"next_button,omitempty"`
	QRCode     string `json:"qr_code,omitempty" bson:"qr_code,omitempty"`
}

// SpinnerSkin 只带 color，嵌在 spinner 下。不写 image / type。
type SpinnerSkin struct {
	Color string `json:"color,omitempty" bson:"color,omitempty"`
}

// LoadingSkin 只给模拟器留颜色。不写 spinner.image / spinner.type。
type LoadingSkin struct {
	BackgroundColor string      `json:"backgroundColor,omitempty" bson:"backgroundColor,omitempty"`
	Spinner         SpinnerSkin `json:"spinner,omitempty" bson:"spinner,omitempty"`
}

// Payload 是要合并进 client_home_config 的皮肤切片。
// 不出现的字段：根 background、根 logo、popups、navigators、entrance.background、
// login.login.icon、login.login.background、login.bottom、base.cup、
// order_pay、order_paid、order_ready_take、order_init_done、
// invite 的任何字段、loading.spinner.image、loading.spinner.type、
// remarks、dine_time、api_refresh_configs。
// 空字符串和空数组的意思是「不要 $set 这条路径」，不是「把租户的值写成空」。
// 页面底色只活在 theme.tokens.background，不进根 background。
type Payload struct {
	HomeConfig    HomeSkin    `json:"home_config,omitempty" bson:"home_config,omitempty"`
	Login         LoginSkin   `json:"login,omitempty" bson:"login,omitempty"`
	MenuConfig    MenuSkin    `json:"menu_config,omitempty" bson:"menu_config,omitempty"`
	OrderConfig   OrderSkin   `json:"order_config,omitempty" bson:"order_config,omitempty"`
	ProfileConfig ProfileSkin `json:"profile_config,omitempty" bson:"profile_config,omitempty"`
	Base          BaseSkin    `json:"base,omitempty" bson:"base,omitempty"`
	Loading       LoadingSkin `json:"loading,omitempty" bson:"loading,omitempty"`
	ShareImage    string      `json:"share_image,omitempty" bson:"share_image,omitempty"`
}

// Model 小程序视觉模版（预设或克隆）。
type Model struct {
	// 模型继承
	model.Model `json:"_" bson:"_"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Kind             Kind               `json:"kind" bson:"kind"`
	Key              string             `json:"key" bson:"key,omitempty"`
	Industry         string             `json:"industry" bson:"industry,omitempty"`
	Name             string             `json:"name" bson:"name,omitempty"`
	Blurb            string             `json:"blurb" bson:"blurb,omitempty"`
	Version          int                `json:"version" bson:"version"`
	Status           Status             `json:"status" bson:"status"`
	Tokens           Tokens             `json:"tokens" bson:"tokens,omitempty"`
	Assets           []Asset            `json:"assets" bson:"assets,omitempty"`
	Payload          Payload            `json:"payload" bson:"payload,omitempty"`
	SourceTemplateID string             `json:"source_template_id,omitempty" bson:"source_template_id,omitempty"`
	SourceVersion    int                `json:"source_version,omitempty" bson:"source_version,omitempty"`
	CopiedAt         int64              `json:"copied_at,omitempty" bson:"copied_at,omitempty"`
	ApplyStatus      string             `json:"apply_status,omitempty" bson:"apply_status,omitempty"` // copying|copied|merged|failed
	LaunchID         string             `json:"launch_id,omitempty" bson:"launch_id,omitempty"`
	Error            string             `json:"error,omitempty" bson:"error,omitempty"`
	OrphanIDs        []string           `json:"orphan_ids,omitempty" bson:"orphan_ids,omitempty"`
}

// CollectionName 返回表名称。
// 不要用 client_ + name + _config。那会得到 client_theme_config。
func (m *Model) CollectionName() string {
	return "client_theme_template"
}
