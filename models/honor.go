package models

import "encoding/json"

type HonorList struct {
	// HonorId   int64 `gorm:"primaryKey;AUTO_INCREMENT"`
	// UpdatedAt time.Time
	// Users     []User `gorm:"many2many:user_honor;foreignKey:HonorId;References:UserId"`
	Contribution int64  `json:"user_contribution"` // 用户的公益贡献值
	HeaderPhoto  string `json:"user_head_photo"`   // 用户头像的url
	UserID       int64  `json:"user_id"`           // 用户id
	Medals       int64  `json:"user_medal"`        // 用户获得的奖牌数
	Name         string `json:"user_name"`         // 用户名称
}

func (m *HonorList) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

func (m *HonorList) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}
