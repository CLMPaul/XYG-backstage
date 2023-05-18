package models

import (
	"gorm.io/gorm"
	"time"
)

type Welfare struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	//for me
	WelfareName          string `json:"welfare_name,omitempty"`
	WelfarePicture       string `json:"welfare_picture,omitempty"`
	WelfareDetails       string `json:"welfare_details,omitempty"`
	WelfareInfo          string `json:"welfare_info,omitempty"`
	WelfareView          int64  `json:"welfare_view,omitempty"`
	WelfareCollection    int64  `json:"welfare_collection,omitempty"`
	WelfareJoin          int64  `json:"welfare_join,omitempty"`
	WelfareAddress       string `json:"welfare_address,omitempty"`
	WelfareLikes         int64
	WelfareTime          int64  `json:"welfare_time,omitempty"`
	StartDate            string `json:"start_date,omitempty"`
	EndDate              string `json:"end_date,omitempty"`
	RecruitsPeople       string `json:"recruits_people,omitempty"`
	PostUserId           int64
	PostUser             User                `gorm:"foreignKey:PostUserId;references:UserId"`
	ConnectorName        string              `json:"connector_name"`
	ConnectorTelephone   string              `json:"connector_telephone"`
	ConnectorQrCodePhoto string              `json:"connector_qr_code_photo"`
	LikeUser             []*User             `gorm:"many2many:like_welfare"`
	PicturesUrlList      []WelfarePictureUrl `json:"pictures_url_list" gorm:"foreignKey:WelfareId;references:ID"` //照片url
	//for welfare append
	AppendUser []*User `gorm:"many2many:append_welfare"`
}

type WelfarePictureUrl struct {
	Id        int64
	WelfareId int64
	Url       string
}

// View 获取浏览数
//func (welfare *Welfare) View() uint64 {
//	countStr, _ := cache.GetClient().Get(cache.WelfareViewKey(welfare.ID)).Result()
//	count, _ := strconv.ParseUint(countStr, 10, 64)
//	return count
//}

// AddView 公益活动游览
//func (welfare *Welfare) AddView() {
//	// 增加公益活动浏览数
//	cache.GetClient().Incr(cache.WelfareViewKey(welfare.ID))
//	// 增加排行点击数
//	//cache.GetClient().ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(product.ID)))
//}

type WelfareHistory struct {
	HistoryCount  uint64 `json:"history_count"`
	HistoryPeople uint64 `json:"history_people"`
}

type WelfareMember struct {
	WelfareId int64 `json:"welfare_id"`
	UserId    int64 `json:"user_id"`
	Status    int   `json:"status"` // 用户是否完成公益
}
