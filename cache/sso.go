package cache

import (
	"encoding"
	"encoding/json"
	"fmt"
	"time"
	"xueyigou_demo/global"
)

type SSOUserInfo struct {
	Name        string `redis:"name"`
	WeChatId    string `redis:"wechat_id"`
	GitHubId    string `redis:"github_id"`
	HeaderPhoto string `redis:"headerphoto"`
	QQId        string `redis:"qq_id"`
}

func (m *SSOUserInfo) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

func (m *SSOUserInfo) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

var _ encoding.BinaryMarshaler = new(SSOUserInfo)
var _ encoding.BinaryUnmarshaler = new(SSOUserInfo)

func temp_user_key(id string) string {
	return fmt.Sprintf("sso:uid:%s", id)
}

func UserInfoCache(id string, account *SSOUserInfo) error {
	global.Log.WithField("key", temp_user_key(id)).Info("cache")
	if err := SetByTTL(temp_user_key(id), account, time.Hour); err != nil {
		global.Log.WithError(err).Error("sso user info cache error")
		return err
	}
	return nil
}

func UserInfoGetFromCache(id string) (*SSOUserInfo, error) {
	ret := &SSOUserInfo{}
	if err := GetResult(temp_user_key(id)).Scan(ret); err != nil {
		global.Log.WithError(err).Error("sso get user info cache error")
		return nil, err
	}
	return ret, nil
}

type WeChatAppLetToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (m *WeChatAppLetToken) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

func (m *WeChatAppLetToken) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}
func GetWeChatAppLetToken(appid string) (*WeChatAppLetToken, error) {
	ret := &WeChatAppLetToken{}
	if err := GetResult(fmt.Sprintf("WeChatAppLetToken_%s", appid)).Scan(ret); err != nil {
		global.Log.WithError(err).Error("sso get user info cache error")
		return nil, err
	}
	return ret, nil
}

func CacheWeChatAppLetToken(appid string, token *WeChatAppLetToken) error {
	global.Log.WithField("key", fmt.Sprintf("WeChatAppLetToken_%s", appid)).Info("cache")
	if err := SetByTTL(fmt.Sprintf("WeChatAppLetToken_%s", appid), token, time.Hour*2); err != nil {
		global.Log.WithError(err).Error("sso WeChatAppLetToken_ cache error")
		return err
	}
	return nil
}
