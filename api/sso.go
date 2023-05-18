package api

import (
	"errors"
	"fmt"
	"net/http"
	"xueyigou_demo/global"
	"xueyigou_demo/service"
	"xueyigou_demo/tools"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GitHubLogin(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{
	// 	"result_status": 0,
	// 	"result_msg":    0,
	// 	"accessUrl": fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s",
	// 		service.Githubconf.ClientId, service.Githubconf.RedirectUrl),
	// })
	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&state=S",
		service.Githubconf.ClientId, service.Githubconf.RedirectUrl))
}

type loginform struct {
	Code string
}

func GitHubLoginCallback(c *gin.Context) {
	var form loginform
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	url := service.GetGitHubTokenAuthUrl(form.Code)
	token, err := service.GetThirdPartyToken(url, 2)
	if err != nil {
		global.Log.WithError(err).Error("get github token")
		c.JSON(http.StatusOK, ErrorResponse(errors.New("get github token err")))
		return
	}

	userinfo, err := service.GetThirdPartyUserInfo(token, 2, "")
	if err != nil {
		global.Log.WithError(err).Error("get github user info")
		c.JSON(http.StatusOK, ErrorResponse(errors.New("get github user info err")))
		return
	}
	github_id_f, exist := userinfo["id"].(float64)
	github_id := int64(github_id_f)
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("github id not exist")))
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	github_login, exist := userinfo["login"].(string)
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("github login not exist")))
		return
	}

	//github 的头像获取失败不用管
	github_head, exist := userinfo["avatar_url"].(string)
	if exist {
		github_head, err = service.GetRemoteFile(github_head, c.Request.Host)
		if err != nil {
			global.Log.Error(err.Error())
		}
	}
	global.Log.WithFields(logrus.Fields{
		"GitHubId":    github_id,
		"GitHubLogin": github_login,
	}).Info("github login info")
	response := service.LoginViaGitHub(github_id, github_login, github_head)
	c.JSON(http.StatusOK, response)
}

func WeChatLogin(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect,
		fmt.Sprintf("https://open.weixin.qq.com/connect/qrconnect?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_login&state=STATE#wechat_redirect", global.WeChatAppID, global.WeChatRedirectUri))
}

func WeChatLoginCallBack(c *gin.Context) {
	var form loginform
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	url := service.GetWeChatTokenAuthUrl(form.Code)
	token, err := service.GetThirdPartyToken(url, 0)
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("get wechat token err")))
		return
	}

	userinfo, err := service.GetThirdPartyUserInfo(token, 0, "")
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("get wechat user info err")))
		return
	}

	id, exist := userinfo["openid"].(string)
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("wechat id not exist")))
		return
	}

	name, exist := userinfo["nickname"].(string)
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("wechat name not exist")))
		return
	}

	//头像获取失败不用管
	head_photo, exist := userinfo["headimgurl"].(string)
	if exist {
		head_photo, err = service.GetRemoteFile(head_photo, c.Request.Host)
		if err != nil {
			global.Log.Error(err.Error())
		}
	}
	global.Log.WithFields(logrus.Fields{
		"GitHubId":    id,
		"GitHubLogin": name,
	}).Info("github login info")
	response := service.LoginViaWeChat(id, name, head_photo)
	c.JSON(http.StatusOK, response)
}

func QQLogin(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect,
		fmt.Sprintf("https://graph.qq.com/oauth2.0/authorize?client_id=%s&redirect_uri=%s&response_type=code&scope=get_user_info&state=STATE",
			global.QQAppId, global.QQRedirectUri))
}

func QQLoginCallBack(c *gin.Context) {
	var form loginform
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	url := service.GetQQTokenAuthUrl(form.Code)
	token, err := service.GetThirdPartyToken(url, 1)
	if err != nil {
		global.Log.WithError(err).Error("qq login")
		c.JSON(http.StatusOK, ErrorResponse(errors.New("get qq token err")))
		return
	}
	open_id, err := service.GetQQOpenId(token.(*service.QQToken))
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("get qq open id err")))
		return
	}
	userinfo, err := service.GetThirdPartyUserInfo(token, 1, open_id.OpenId)
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("get qq user info err")))
		return
	}
	global.Log.WithField("info", userinfo).Info("QQLogin")
	name, exist := userinfo["nickname"].(string)
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("qq name not exist")))
		return
	}

	//头像获取失败不用管
	head_photo, exist := userinfo["figureurl_qq_1"].(string)
	if exist {
		head_photo, err = service.GetRemoteFile(head_photo, c.Request.Host)
		if err != nil {
			global.Log.Error(err.Error())
		}
	}
	global.Log.WithFields(logrus.Fields{
		"QQId":   open_id,
		"QQName": name,
	}).Info("github login info")
	response := service.LoginViaQQ(open_id.OpenId, name, head_photo)
	c.JSON(http.StatusOK, response)
}

type weChatAppLetForm struct {
	CodePhone     string `json:"code_phone"`     // 通过 bindgetphonenumber 事件回调获取到动态令牌code，以获取手机号
	CodeSession   string `json:"code_session"`   // 调用 wx.login() 获取 临时登录凭证code
	EncryptedData string `json:"encrypted_data"` // @getphonenumber事件获取到的，由前端进行提供
	Iv            string `json:"iv"`             // 小程序的唯一标识，由前端进行提供
}

func WeChatAppLetCallBack(c *gin.Context) {
	var form weChatAppLetForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	global.Log.WithField("form", form).Info("WeChatAppLetCallBack")
	//获取Session
	session, err := service.Code2Session(form.CodeSession)
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	if session.Errcode != 0 {
		global.Log.WithError(errors.New(session.Errmsg)).Error("WeChatAppLetCallBack")
		c.JSON(http.StatusOK, ErrorResponse(errors.New(session.Errmsg)))
		return
	}
	//

	//获取token
	token, err := service.GetWeChatAppLetToken()
	if err != nil {
		global.Log.WithError(err).Error("GetWeChatAppLetToken")
		c.JSON(http.StatusOK, ErrorResponse(errors.New("get wechat token err")))
		return
	}
	//

	//获取手机号
	phone_info, err := service.GetWeChatAppLetPhone(form.CodePhone, token.AccessToken)
	if err != nil {
		global.Log.WithError(err).Error("GetWeChatAppLetPhone")
		c.JSON(http.StatusOK, ErrorResponse(errors.New("get phone err")))
		return
	}
	//
	response := service.WeChatAppLetBind(phone_info.PhoneInfo.PurePhoneNumber, session, form.EncryptedData, form.Iv)
	c.JSON(http.StatusOK, response)
}

type bindForm struct {
	PhoneNum string `json:"phonenum"`
	Code     string `json:"code"`
	PassWord string `json:"password"`
	UserId   string `json:"user_id"`
	Type     string `json:"type"`
}

func BindPhoneNum(c *gin.Context) {
	var form bindForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	valid := tools.StringValid(global.PhonecPattern, form.PhoneNum)
	if !valid {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("phone num error")))
		return
	}
	//for test
	// id := global.Worker.GetId()
	// account := &cache.SSOUserInfo{
	// 	Name:        "1",
	// 	WeChatId:    "1",
	// 	HeaderPhoto: "1",
	// }
	// cache.UserInfoCache(strconv.Itoa(int(id)), account)
	//
	res := service.SSOBindPhoneNum(form.PhoneNum, form.Code, form.PassWord, form.UserId, form.Type)
	c.JSON(http.StatusOK, res)
}
