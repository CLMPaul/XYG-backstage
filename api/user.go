package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"xueyigou_demo/dao"
	"xueyigou_demo/global"
	"xueyigou_demo/middleware"
	"xueyigou_demo/models"
	"xueyigou_demo/pkg/e"
	"xueyigou_demo/serializer"
	"xueyigou_demo/service"
	"xueyigou_demo/tools"

	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts

func UserRegister(c *gin.Context) {
	var userService service.UserService
	if err := c.ShouldBind(&userService); err == nil {
		vaild := tools.StringValid(global.PhonecPattern, userService.Phone)
		if !vaild {
			c.JSON(200, ErrorResponse(errors.New("phone format error")))
			return
		}
		res := userService.Register(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

func UserLogin(c *gin.Context) {
	var userService service.UserService
	if err := c.ShouldBind(&userService); err == nil {
		res := userService.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

type userLoginViaCodeForm struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

func UserLoginViaCode(c *gin.Context) {
	var form userLoginViaCodeForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
	vaild := tools.StringValid(global.PhonecPattern, form.Phone)
	if !vaild {
		c.JSON(200, ErrorResponse(e.PhonePatternError{Phone: form.Phone}))
		return
	}
	res := service.LoginViaCode(form.Phone, form.Code)
	c.JSON(http.StatusOK, res)
}
func UserAttention(c *gin.Context) {
	var userService service.UserService
	if err := c.ShouldBind(&userService); err == nil {
		res := userService.GetAttention(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

func UserAttentionAction(c *gin.Context) {
	var userService service.UserAttentionService
	if err := c.ShouldBind(&userService); err == nil {
		res := userService.ParseAttentionAction(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

func UserFans(c *gin.Context) {
	var userService service.UserService
	if err := c.ShouldBind(&userService); err == nil {
		res := userService.GetFans(c)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}
func UserInfo(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	visitorId, _ := strconv.ParseInt(c.Query("visitor_id"), 10, 64)
	response := service.GetUserInfo(userId, visitorId)
	c.JSON(http.StatusOK, response)
}

func Resetkey(c *gin.Context) { //账号安全（重设密码）
	var newservice service.ResetKeyService
	if err := c.ShouldBind(&newservice); err == nil {
		userid := dao.GetUidByPhone(newservice.PhoneNumber)
		if userid == 0 {
			c.JSON(200, serializer.BuildFailResponse("phone number has not been registered"))
			return
		}
		if !tools.StringValid(global.PhonecPattern, newservice.PhoneNumber) {
			c.JSON(200, serializer.BuildFailResponse("phone number form is wrong"))
			return
		}
		res := service.ResetKey(userid, newservice.NewPassword, newservice.PhoneNumber, newservice.VerifyCode)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}
func SetUserInfo(c *gin.Context) {
	var newservice service.SetUserInfoService
	if err := c.ShouldBind(&newservice); err == nil {
		userClaim, exist := c.Get("userClaim")
		if !exist {
			c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
			return
		}
		claim := userClaim.(*middleware.UserClaims)
		res := newservice.SetUserInfo(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}
func GetUserInfo(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	res := service.GetUserDetailInfo(userId)
	c.JSON(200, res)
}

// func ConnectCreator(c *gin.Context){
// 	var userService service.GoodsReturnService
// 	if err := c.ShouldBind(&userService); err == nil {
// 		res := userService.GetConnectCreator(c)
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(400, ErrorResponse(err))
// 	}
// }

type postAddressForm struct {
	AddressDetails    string `json:"address_details"`    // 详细地址信息，如街道、小区、单元等
	AddressMsg        string `json:"address_msg"`        // 粗略地址信息，省/市/区
	AddressName       string `json:"address_name"`       // 收件人姓名
	AddressPostalcode string `json:"address_postalcode"` // 该地址当地的邮政编码
	AddressStatus     int64  `json:"address_status"`     // 地址状态码，0-设置为默认，1-设置为非默认
	AddressTelenum    string `json:"address_telenum"`    // 收件人电话号码
}

func GetAllAddress(c *gin.Context) {
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	response := service.GetAllAddress(claim.Id)
	c.JSON(http.StatusOK, response)
}

func PostAddress(c *gin.Context) {
	var addressform postAddressForm
	if err := c.ShouldBind(&addressform); err != nil {
		panic(err)
	}
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	user_id := claim.Id
	address := models.Useraddress{
		AddressDetails:    addressform.AddressDetails,
		AddressMsg:        addressform.AddressMsg,
		AddressName:       addressform.AddressName,
		AddressPostalcode: addressform.AddressPostalcode,
		AddressStatus:     addressform.AddressStatus,
		AddressTelenum:    addressform.AddressTelenum,
		UserID:            user_id,
	}
	response := service.PostUserAddress(claim, &address)
	c.JSON(http.StatusOK, response)
}

func GetIndentity(c *gin.Context) {
	user_id, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	response := service.GetIndentity(user_id, claim)
	c.JSON(http.StatusOK, response)
}

func PostIndentity(c *gin.Context) {
	var indentity service.Indentityservice
	if err := c.ShouldBind(&indentity); err != nil {
		panic(err)
	}
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	valid := tools.StringValid(global.PhonecPattern, indentity.UserTelenum)
	if !valid {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("phone format error")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	indentity.UserID = claim.Id
	response := service.PostIndentity(&indentity, claim)
	c.JSON(http.StatusOK, response)
}

type peopleform struct {
	PeopleDetails    string   `json:"people_details"` // 人才具体说明
	PeopleIntroduce  string   `json:"people_introduce"`
	PeopleMajorskill string   `json:"people_majorskill"` // 人才主要技能
	PeopleSubject    []string `json:"people_subject"`    // 人才具体擅长标签
	PeopleType       string   `json:"people_type"`       // 人才专业类型
	PeopleMax        int64    `json:"people_max"`        //人才能接受的最高报价
	PeopleMin        int64    `json:"people_min"`        //人才能接受的最低报价
	PeopleIp         string   `json:"people_ip"`
}

func PostPeople(c *gin.Context) {
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	var form peopleform
	if err := c.ShouldBind(&form); err != nil {
		panic(err)
	}
	if form.PeopleMax < form.PeopleMin {
		c.String(http.StatusOK, "Illegal data")
		return
	}
	people := models.PeopleMid{
		//PeopleDetails: global.IllegalWords.Replace(form.PeopleDetails, '*'),
		PeopleDetails:    form.PeopleDetails,
		PeopleMajorskill: form.PeopleMajorskill,
		PeopleType:       form.PeopleType,
		PeopleMax:        form.PeopleMax,
		PeopleMin:        form.PeopleMin,
		//PeopleIntroduce:  global.IllegalWords.Replace(form.PeopleIntroduce, '*'),
		PeopleIntroduce: form.PeopleIntroduce,
		PeopleIp:        form.PeopleIp,
		PeopleId:        claim.Id,
		PeopleStatus:    0,
	}
	if exist := strings.Contains(people.PeopleIp, "/"); !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("peopleIp is illegal")))
		return
	}
	for _, subject := range form.PeopleSubject {
		people.PeopleSubject = append(people.PeopleSubject, models.PeopleMidSubjectItem{PeopleId: claim.Id, Item: subject})
	}
	response := service.PostPeople(&people, claim)
	c.JSON(http.StatusOK, response)
}

func PostPeopleMid(c *gin.Context) {
	peopleId, _ := strconv.ParseInt(c.Query("people_id"), 10, 64)
	ispassed, _ := strconv.ParseInt(c.Query("ispassed"), 10, 64)
	response := service.PostPeopleMid(peopleId, ispassed)
	c.JSON(http.StatusOK, response)
}

func GetPeopleForMe(c *gin.Context) {
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	response := service.GetPeopleForMe(claim.Id)
	c.JSON(http.StatusOK, response)
}

func GetOrderTaskList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	println(userId)
	response := service.GetOrderTaskList(userId)
	c.JSON(http.StatusOK, response)
}

func PostOrderTask(c *gin.Context) {
	var taskId int64
	taskId, err := strconv.ParseInt(c.Query("task_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	var userId int64
	userId, err = strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	response := service.PostOrderTask(userId, taskId)
	c.JSON(http.StatusOK, response)

}

func FinishOrderTask(c *gin.Context) {
	var taskId int64
	taskId, err := strconv.ParseInt(c.Query("task_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	response := service.FinishTask(claim.Id, taskId)
	c.JSON(http.StatusOK, response)
}

func PostOrderWork(c *gin.Context) {
	var userId int64
	userId, _ = strconv.ParseInt(c.Query("user_id"), 10, 64)
	var workId int64
	workId, _ = strconv.ParseInt(c.Query("work_id"), 10, 64)
	response := service.PostOrderWork(userId, workId)
	c.JSON(http.StatusOK, response)
}

func GetOrderWorkList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	response := service.GetOrderWorkList(userId)
	c.JSON(http.StatusOK, response)
}

type DelItemForm struct {
	UserId   string `json:"user_id"`
	TypeId   int    `json:"type_id"`
	ObjectId string `json:"object_id"`
}

func DelOrderItem(c *gin.Context) {
	var form DelItemForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(400, ErrorResponse(err))
		return
	}
	if form.TypeId != 0 && form.TypeId != 1 {
		c.JSON(http.StatusOK, "Item is illegal")
		return
	}
	uId, _ := strconv.ParseInt(form.UserId, 10, 64)
	oId, _ := strconv.ParseInt(form.ObjectId, 10, 64)
	//fmt.Println(form)
	response := service.DelOrderItem(uId, form.TypeId, oId)
	c.JSON(http.StatusOK, response)
}

func Feedback(c *gin.Context) {
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)

	var feedbackform models.UserFeedbackForm
	if err := c.ShouldBind(&feedbackform); err == nil {
		feedback := models.UserFeedback{
			FeedbackContent:  feedbackform.FeedbackContent,
			FeedbackPictures: nil,
			UserFeedbackID:   claim.Id,
		}
		feedback.FeedbackPictures = make([]models.UserFeedbackPicturesURL, 0)
		for _, i := range feedbackform.FeedbackPictures {
			feedback.FeedbackPictures = append(feedback.FeedbackPictures, models.UserFeedbackPicturesURL{
				Url: i})
		}
		response := service.Feedback(feedback)
		c.JSON(http.StatusOK, response)
	}
}
