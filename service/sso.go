package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"xueyigou_demo/cache"
	"xueyigou_demo/dao"
	"xueyigou_demo/global"
	"xueyigou_demo/middleware"
	"xueyigou_demo/models"
	"xueyigou_demo/serializer"
	"xueyigou_demo/tools"
)

type GithubConf struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
}

var Githubconf = GithubConf{

	// test
	// ClientId:     "7a70727e08e78b92c281",
	// ClientSecret: "3689b4633f04672bc2375f4a7535b6728a0dece5",
	// RedirectUrl:  "http://127.0.0.1:8080/xueyigou/sso/githubcallback",

	//
	RedirectUrl:  "http://120.77.85.52:30001/xueyigou/sso/githubcallback",
	ClientId:     "051066dca60430d2e8c6",
	ClientSecret: "b1d4d8943ba239bbaaebad3d2f69f83a06860fed",
}

func GetGitHubTokenAuthUrl(code string) string {
	return fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		Githubconf.ClientId, Githubconf.ClientSecret, code,
	)
}

type GitHubToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func LoginViaGitHub(github_id int64, github_name string, github_head string) interface{} {
	user, err := dao.GetUserByGitHubId(github_id)
	if err != nil {
		id := global.Worker.GetId()
		account := &cache.SSOUserInfo{
			Name:        github_name,
			GitHubId:    strconv.Itoa(int(github_id)),
			HeaderPhoto: github_head,
		}
		cache.UserInfoCache(strconv.Itoa(int(id)), account)
		return serializer.UserLoginResponse{
			ResultStatus: 1,
			ResultMsg:    "未绑定",
			UserId:       strconv.FormatInt(id, 10),
		}
	}
	newClaim := middleware.UserClaims{Id: user.UserId, Name: user.Name}
	token, refresh_token := middleware.GenerateToken(&newClaim)
	Token := serializer.Token{
		AccessToken:  token,
		RefreshToken: refresh_token,
	}
	return serializer.UserLoginResponse{
		ResultStatus: 0,
		ResultMsg:    "login",
		UserId:       strconv.FormatInt(user.UserId, 10),
		Token:        Token,
	}
}

func GetQQTokenAuthUrl(code string) string {
	return fmt.Sprintf(
		"https://graph.qq.com/oauth2.0/token?grant_type=authorization_code&client_id=%s&client_secret=%s&code=%s&redirect_uri=%s&fmt=json",
		global.QQAppId, global.QQAppSecret, code, global.QQRedirectUri,
	)
}

type QQToken struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        string `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	Error            int    `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func GetWeChatTokenAuthUrl(code string) string {
	return fmt.Sprintf(
		"https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		global.WeChatAppID, global.WeChatAPPSecret, code,
	)
}

type WeChatToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
}

func GetWeChatAppLetTokenAuthUrl() string {
	return fmt.Sprintf(
		"https://api.weixin.qq.com/cgi-bin/token?appid=%s&secret=%s&grant_type=client_credential",
		global.WeChatAppID, global.WeChatAPPSecret,
	)
}

type WeChatAppLetPhoneInfo struct {
	Errcode   int64     `json:"errcode"`
	Errmsg    string    `json:"errmsg"`
	PhoneInfo phoneInfo `json:"phone_info"`
}

type phoneInfo struct {
	CountryCode     string    `json:"countryCode"`
	PhoneNumber     string    `json:"phoneNumber"`
	PurePhoneNumber string    `json:"purePhoneNumber"`
	Watermark       watermark `json:"watermark"`
}

func GetWeChatAppLetPhone(code string, access_token string) (*WeChatAppLetPhoneInfo, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s", access_token)
	// 形成请求
	var req *http.Request
	var err error
	body := make(map[string]string)
	body["code"] = code
	byte_body, _ := json.Marshal(body)
	if req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(byte_body))); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		return nil, err
	}
	global.Log.WithField("res", (*res).Body).Info("get tp token")
	// 将响应体解析为 token，并返回
	var phone_info WeChatAppLetPhoneInfo
	if err = json.NewDecoder(res.Body).Decode(&phone_info); err != nil {
		return nil, err
	}
	return &phone_info, nil
}

type AppLetSession struct {
	Sessionkey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Errmsg     string `json:"errmsg"`
	Openid     string `json:"openid"`
	Errcode    int32  `json:"errcode"`
}

func Code2Session(code string) (*AppLetSession, error) {
	url := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		global.WeChatAPPLETAppID, global.WeChatAPPLETAppSecret, code,
	)
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		return nil, err
	}
	global.Log.WithField("res", (*res).Body).Info("Code2Session")
	// 将响应体解析为 token，并返回
	var session AppLetSession
	if err = json.NewDecoder(res.Body).Decode(&session); err != nil {
		global.Log.WithError(err).Error("Code2Session")
		return nil, err
	}
	return &session, nil
}

func GetWeChatAppLetToken() (*cache.WeChatAppLetToken, error) {
	token, err := cache.GetWeChatAppLetToken(global.WeChatAPPLETAppID)
	if err == nil {
		return token, nil
	}
	token = &cache.WeChatAppLetToken{}
	url := fmt.Sprintf(
		"https://api.weixin.qq.com/cgi-bin/token?appid=%s&secret=%s&grant_type=client_credential",
		global.WeChatAPPLETAppID, global.WeChatAPPLETAppSecret,
	)
	var req *http.Request
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return nil, err
	}
	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		return nil, err
	}
	global.Log.WithField("res", (*res).Body).Info("GetWeChatAppLetToken")
	// 将响应体解析为 token，并返回
	if err = json.NewDecoder(res.Body).Decode(token); err != nil {
		return nil, err
	}
	cache.CacheWeChatAppLetToken(global.WeChatAPPLETAppID, token)
	return token, nil
}

// func WeChatTokenRefresh(refresh_token string) (*WeChatToken, error) {
// 	url := fmt.Sprintf(
// 		"https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s",
// 		global.WeChatAppID, refresh_token,
// 	)

// 	// 形成请求
// 	var req *http.Request
// 	var err error
// 	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
// 		return nil, err
// 	}
// 	req.Header.Set("accept", "application/json")

// 	// 发送请求并获得响应
// 	var httpClient = http.Client{}
// 	var res *http.Response
// 	if res, err = httpClient.Do(req); err != nil {
// 		return nil, err
// 	}

// 	// 将响应体解析为 token，并返回
// 	var token WeChatToken
// 	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
// 		return nil, err
// 	}
// 	return &token, nil
// }

func LoginViaWeChat(wechat_id string, wechat_name string, wechat_head string) interface{} {
	user, err := dao.GetUserByWeChatId(wechat_id)
	if err != nil {
		id := global.Worker.GetId()
		account := &cache.SSOUserInfo{
			Name:        wechat_name,
			WeChatId:    wechat_id,
			HeaderPhoto: wechat_head,
		}
		cache.UserInfoCache(strconv.Itoa(int(id)), account)
		return serializer.UserLoginResponse{
			ResultStatus: 1,
			ResultMsg:    "未绑定",
			UserId:       strconv.FormatInt(id, 10),
		}
	}
	newClaim := middleware.UserClaims{Id: user.UserId, Name: user.Name}
	token, refresh_token := middleware.GenerateToken(&newClaim)
	Token := serializer.Token{
		AccessToken:  token,
		RefreshToken: refresh_token,
	}
	return serializer.UserLoginResponse{
		ResultStatus: 0,
		ResultMsg:    "login",
		UserId:       strconv.FormatInt(user.UserId, 10),
		Token:        Token,
	}
}

func GetThirdPartyToken(url string, aid int) (interface{}, error) {
	// 形成请求
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		return nil, err
	}
	global.Log.WithField("res", (*res).Body).Info("get tp token")
	// 将响应体解析为 token，并返回
	switch aid {
	case 0:
		var token WeChatToken
		if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
			return nil, err
		}
		return &token, nil
	case 1:
		var token QQToken
		if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
			return nil, err
		}
		// if token.Error != 0 {
		// 	global.Log.WithField("token", token).Info("QQLogin")
		// }
		global.Log.WithField("token", token).Info("QQLogin")
		return &token, nil
	case 2:
		var token GitHubToken
		if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
			return nil, err
		}
		global.Log.WithField("token", token).Info("github token")
		return &token, nil
	case 3:
		var token cache.WeChatAppLetToken
		if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
			return nil, err
		}
		global.Log.WithField("token", token).Info("github token")
		return &token, nil
	}

	return nil, errors.New("unkonw aid")
}

func wechatUserInfoReq(token *WeChatToken) (*http.Request, error) {

	var userInfoUrl = fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s", token.AccessToken, token.OpenId)
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))
	return req, nil
}

func gitHubUserInfoReq(token *GitHubToken) (*http.Request, error) {

	var userInfoUrl = "https://api.github.com/user"
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))
	return req, nil
}

type QQOpenId struct {
	ClientId string `json:"client_id"`
	OpenId   string `json:"openid"`
}

func qqUserInfoReq(token *QQToken, open_id string) (*http.Request, error) {
	var userInfoUrl = fmt.Sprintf("https://graph.qq.com/user/get_user_info?access_token=%s&oauth_consumer_key=%s&openid=%s",
		token.AccessToken, global.QQAppId, open_id)
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))
	return req, nil
}

func GetQQOpenId(token *QQToken) (*QQOpenId, error) {
	var userInfoUrl = fmt.Sprintf("https://graph.qq.com/oauth2.0/me?access_token=%s&fmt=json", token.AccessToken)
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))

	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, err
	}
	openid := QQOpenId{}
	if err = json.NewDecoder(res.Body).Decode(&openid); err != nil {
		return nil, err
	}
	return &openid, nil
}

func LoginViaQQ(qq_id string, qq_name string, qq_head string) interface{} {
	user, err := dao.GetUserByQQId(qq_id)
	if err != nil {
		id := global.Worker.GetId()
		account := &cache.SSOUserInfo{
			Name:        qq_name,
			QQId:        qq_id,
			HeaderPhoto: qq_head,
		}
		cache.UserInfoCache(strconv.Itoa(int(id)), account)
		return serializer.UserLoginResponse{
			ResultStatus: 1,
			ResultMsg:    "未绑定",
			UserId:       strconv.FormatInt(id, 10),
		}
	}
	newClaim := middleware.UserClaims{Id: user.UserId, Name: user.Name}
	token, refresh_token := middleware.GenerateToken(&newClaim)
	Token := serializer.Token{
		AccessToken:  token,
		RefreshToken: refresh_token,
	}
	return serializer.UserLoginResponse{
		ResultStatus: 0,
		ResultMsg:    "login",
		UserId:       strconv.FormatInt(user.UserId, 10),
		Token:        Token,
	}
}

// 获取用户信息
// IN aid:第三方id, open_id for qq only
func GetThirdPartyUserInfo(token interface{}, aid int, open_id string) (map[string]interface{}, error) {

	var req *http.Request
	var err error
	switch aid {
	case 0:
		req, err = wechatUserInfoReq(token.(*WeChatToken))
		if err != nil {
			return nil, err
		}
	case 1:
		req, err = qqUserInfoReq(token.(*QQToken), open_id)
		if err != nil {
			return nil, err
		}
	case 2:
		req, err = gitHubUserInfoReq(token.(*GitHubToken))
		if err != nil {
			return nil, err
		}
	}
	// 发送请求并获取响应
	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, err
	}

	// 将响应的数据写入 userInfo 中，并返回
	var userInfo = make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}

// func WeChatAppLetLogin(unionid string, openid string) interface{} {
// 	user, err := dao.GetUserByOpenid(openid)
// 	if err != nil {
// 		global.Log.WithError(err).Error("WeChatAppLetLogin")
// 	}
// }

type AppLetUserInfo struct {
	AvatarURL string    `json:"avatarUrl"`
	City      string    `json:"city"`
	Country   string    `json:"country"`
	Gender    string    `json:"gender"`
	NickName  string    `json:"nickName"`
	OpenID    string    `json:"openId"`
	Province  string    `json:"province"`
	UnionID   string    `json:"unionId"`
	Watermark watermark `json:"watermark"`
}

type watermark struct {
	Appid     string `json:"appid"`
	Timestamp int    `json:"timestamp"`
}

func WeChatAppLetBind(phone_num string, session *AppLetSession, encrypted_data string, iv string) interface{} {
	user, err := dao.GetUserByPhoneNum(phone_num)

	if err != nil {
		//未绑定
		user_info_byte, err := tools.AesDecrypt(encrypted_data, session.Sessionkey, iv)
		if err != nil {
			global.Log.WithError(err).Error("WeChatAppLetBind")
			return serializer.BuildFailResponse("解密用户信息错误")
		}
		user_info := &AppLetUserInfo{}
		err = json.Unmarshal(user_info_byte, user_info)
		if err != nil {
			global.Log.WithError(err).Error("WeChatAppLetBind")
			return serializer.BuildFailResponse("json解码错误")
		}
		if user_info.Watermark.Appid != global.WeChatAPPLETAppID {
			err = errors.New("Appid错误")
			global.Log.WithError(err).Error("WeChatAppLetBind")
			return serializer.BuildFailResponse("Appid错误")
		}
		global.Log.WithField("user_info", user_info).Info("WeChatAppLetBind")
		id := global.Worker.GetId()
		user = &models.User{
			UserId:      id,
			Name:        user_info.NickName,
			City:        user_info.City,
			HeaderPhoto: user_info.AvatarURL,
			Telephone:   phone_num,
		}
		account := models.Account{
			Id:       id,
			UserName: user_info.NickName,
			Phone:    phone_num,
		}
		dao.AddUser(*user, account)

	}
	user.AppLetOpenid = session.Openid
	user.AppLetUnionid = session.Unionid
	if dao.UpdateUserInfoByStruct(user) != nil {
		return serializer.BuildFailResponse("save error")
	}
	newClaim := middleware.UserClaims{Id: user.UserId, Phone: user.Telephone}
	token, refresh_token := middleware.GenerateToken(&newClaim)
	Token := serializer.Token{
		AccessToken:  token,
		RefreshToken: refresh_token,
	}
	return serializer.UserLoginResponse{
		ResultStatus: 0,
		ResultMsg:    "Success",
		Token:        Token,
		UserId:       strconv.FormatInt(user.UserId, 10),
	}
}

func SSOBindPhoneNum(phone_num string, code string, password string, user_id string, type_ string) interface{} {
	err := VerifSmsCodeWithPhone(phone_num, code, global.TemplateCodeForRegister)
	if err != nil {
		return serializer.BuildFailResponse(err.Error())
	}
	userinfo, err := cache.UserInfoGetFromCache(user_id)
	if err != nil {
		global.Log.WithError(err).Error("sso bind")
		return serializer.BuildFailResponse("user_id expire")
	}

	var user *models.User
	user, err = dao.GetUserByPhoneNum(phone_num)
	if err != nil {
		//手机号没绑定过
		user = &models.User{}
		id, err := strconv.ParseInt(user_id, 10, 64)
		if err != nil {
			global.Log.WithError(err).Error("sso bind")
		}
		user.UserId = id
		user.Name = userinfo.Name
		password = middleware.Md5Crypt(password, "xueyigou")
		account := models.Account{
			Id:       id,
			UserName: userinfo.Name,
			PassWord: password,
			Phone:    phone_num,
		}
		user.GitHubId, err = strconv.ParseInt(userinfo.GitHubId, 10, 64)
		if err != nil {
			global.Log.WithError(err).Error("sso bind")
		}

		user.WeChatId = userinfo.WeChatId

		user.QqId = userinfo.QQId

		user.Telephone = phone_num

		user.HeaderPhoto = userinfo.HeaderPhoto
		dao.AddUser(*user, account)
	} else {
		if password != "" {
			password = middleware.Md5Crypt(password, "xueyigou")
			dao.SetNewPassword(user.UserId, password)
		}
		if type_ == "0" && user.WeChatId != "" {
			return serializer.BuildFailResponse("当前账号已绑定微信")
		}
		if type_ == "1" && user.QqId != "" {
			return serializer.BuildFailResponse("当前账号已绑定QQ")
		}
		if type_ == "2" && user.GitHubId != 0 {
			return serializer.BuildFailResponse("当前账号已绑定github")
		}

		if type_ == "0" {
			user.WeChatId = userinfo.WeChatId
		} else if type_ == "1" {
			user.QqId = userinfo.QQId
		} else if type_ == "2" {
			user.GitHubId, err = strconv.ParseInt(userinfo.GitHubId, 10, 64)
			if err != nil {
				global.Log.WithError(err).Error("sso bind")
			}
		}
		user.Telephone = phone_num

		user.HeaderPhoto = userinfo.HeaderPhoto

		if dao.UpdateUserInfoByStruct(user) != nil {
			return serializer.BuildFailResponse("save error")
		}
	}
	newClaim := middleware.UserClaims{Id: user.UserId, Phone: user.Telephone}
	token, refresh_token := middleware.GenerateToken(&newClaim)
	Token := serializer.Token{
		AccessToken:  token,
		RefreshToken: refresh_token,
	}
	return serializer.UserLoginResponse{
		ResultStatus: 0,
		ResultMsg:    "Success",
		Token:        Token,
		UserId:       strconv.FormatInt(user.UserId, 10),
	}
}
