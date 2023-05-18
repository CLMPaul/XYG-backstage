package models

//// Account 平台账户
//type Account struct {
//	ID           string         `gorm:"size:36;primaryKey"`
//	Username     string         `gorm:"size:20;not null"`
//	PasswordHash datatypes.Blob `gorm:"not null"`
//	PasswordSalt datatypes.Blob `gorm:"not null"`
//	Name         string         `gorm:"size:20;not null;index"`
//	Locked       bool           `gorm:"not null;default:false;index"`
//	IsSuperuser  bool           `gorm:"not null;default:false"`
//
//	Email  *string `gorm:"size:64;index"`
//	Tel    *string `gorm:"size:64;index"`
//	Mobile *string `gorm:"size:64;index"`
//
//	// 更新 UpdatedTime 之后再删除 UserForAuthenticate 缓存，可以强制用户下线
//	UpdateTime *time.Time
//
//	// 新建用户为 false，重置密码时设为 false
//	// 用户自行修改密码时设为 true
//	// 如果为 false，前端在登录时要提示用户修改密码
//	PasswordChanged bool `gorm:"not null;default:false"`
//}
//
//var Accounts = new(Account)
