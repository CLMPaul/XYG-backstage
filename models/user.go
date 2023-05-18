package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    int64  `json:"user_id" gorm:"primaryKey"`
	Name      string `json:"user_name,omitempty"` //用户名
	//FollowCount    int64       `json:"follow_count,omitempty"`
	//FollowerCount  int64       `json:"follower_count,omitempty"`
	//IsFollow       bool        `json:"is_follow,omitempty"`
	Sex             string        `json:"sex,omitempty"`
	HeaderPhoto     string        `json:"header_photo,omitempty"`
	BackgroundPhoto string        `json:"background_photo,omitempty"`
	Addresses       []Useraddress `json:"address_list,omitempty" gorm:"foreignKey:UserID;references:UserId"`
	Email           string        `json:"email,omitempty"`
	Wechat          string        `json:"wechat,omitempty"`
	Hometown        string        `json:"hometown,omitempty"`
	Province        string        `json:"province,omitempty"`
	City            string        `json:"city,omitempty"`
	Sign            string        `json:"sign,omitempty"`
	Graduate        string        `json:"graduate,omitempty"`
	Subject         string        `json:"subject,omitempty"`
	QQ              string        `json:"qq,omitempty"`
	Detail          string        `json:"detail,omitempty"`
	LabelList       []UserLables  `json:"label_list"`
	Level           uint          `json:"level"`
	Hot             uint64        `json:"hot"`
	Contribution    uint64        `json:"contribution"`
	Medals          uint          `json:"medals"`
	Deals           int64         `json:"deals"`
	WelfareTime     int64         `json:"welfare_time"`
	//for attentions and fans
	UserFollows   []User `gorm:"many2many:user_follows"`
	UserFollowers []User `gorm:"many2many:user_followers"`

	UserIdentity string `json:"user_identity,omitempty"` // 用户身份，用户身份，如：普通用户

	//for collection
	CollectionWelfare []Welfare `gorm:"many2many:welfare_collection"`
	CollectionWork    []*Work   `gorm:"many2many:work_collection"`
	CollectionTask    []*Task   `gorm:"many2many:task_collection"`
	CollectionEvent   []*Event  `gorm:"many2many:event_collection"`

	// for like
	LikeWork  []*Work  `gorm:"many2many:like_work"`
	LikeTask  []*Task  `gorm:"many2many:like_task"`
	LikeEvent []*Event `gorm:"many2many:like_event"`

	//for order
	OrderWork []*Work `gorm:"many2many:order_work"`
	OrderTask []*Task `gorm:"many2many:order_task"`

	//for welfare
	Telephone   string `json:"telephone,omitempty"`
	QrCodePhoto string `json:"qr_code_photo"`

	//sso
	GitHubId      int64  `json:"github_id"`
	WeChatId      string `json:"wechat_id"`
	QqId          string `json:"qq_id"`
	AppLetOpenid  string
	AppLetUnionid string
	//for official_message
	OfficialMessageList []OfficialMessage `gorm:"many2many:official_message_user""` //推送官方消息
	//for welfare append
	AppendWelfare []*Welfare `gorm:"many2many:append_welfare"`
}

type UserLables struct {
	gorm.Model
	UserId int64
	Label  string
}
type Account struct {
	Id       int64  `json:"id,omitempty"`
	UserName string `json:"username,omitempty"`
	PassWord string `json:"password,omitempty"`
	Phone    string `json:"phone"`
}

type Useraddress struct {
	AddressDetails    string `json:"address_details"`    // 详细地址信息，如街道、小区、单元等
	AddressMsg        string `json:"address_msg"`        // 粗略地址信息，省/市/区
	AddressName       string `json:"address_name"`       // 收件人姓名
	AddressPostalcode string `json:"address_postalcode"` // 该地址当地的邮政编码
	AddressStatus     int64  `json:"address_status"`     // 地址状态码，0-设置为默认，1-设置为非默认
	AddressTelenum    string `json:"address_telenum"`    // 收件人电话号码
	UserID            int64  `json:"user_id"`            // 用户id
}

type Indentity struct {
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserId       int64
	UserTelenum  string  `json:"user_telenum"`         // 用户电话号码
	UserIdentity string  `json:"user_identity"`        // 学生身份认证，0-在校学生，1-非在校学生
	UserIdnum    string  `json:"user_idnum"`           // 用户学生证号
	UserName     string  `json:"user_name"`            // 用户姓名
	UserPhoto    *string `json:"user_photo,omitempty"` // 用户学生证照片地址url
}

type LikeForm struct {
	UserId     int64
	ObjectId   int64
	ObjectType int64
	ActType    bool
}
