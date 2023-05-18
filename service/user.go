package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"xueyigou_demo/dao"
	"xueyigou_demo/global"
	"xueyigou_demo/middleware"
	"xueyigou_demo/models"
	"xueyigou_demo/pkg/e"
	"xueyigou_demo/serializer"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Phone    string `json:"phone,omitempty"`
	Code     string `json:"code"`
}

func (service *UserService) Register(c *gin.Context) interface{} {
	if exist := dao.AccountExist(service.UserName); !exist {
		return serializer.Response{
			ResultStatus: 1,
			ResultMsg:    "UserName already exist",
		}
	}
	if exist := dao.AccountExistPhoneNum(service.Phone); !exist {
		return serializer.Response{
			ResultStatus: 1,
			ResultMsg:    "UserPhone already exist",
		}
	}
	if service.Password == "" || len(service.Password) < 6 || service.UserName == "" {
		return serializer.Response{
			ResultStatus: 1,
			ResultMsg:    "length of password is too short",
		}
	}
	err := VerifSmsCodeWithPhone(service.Phone, service.Code, global.TemplateCodeForRegister)
	if err != nil {
		return serializer.BuildFailResponse(err.Error())
	}
	service.Password = middleware.Md5Crypt(service.Password, "xueyigou")

	// 创建账户，生成用户
	id := global.Worker.GetId()
	newAccount := models.Account{
		UserName: service.UserName,
		PassWord: service.Password,
		Id:       id,
		Phone:    service.Phone,
	}
	newUser := models.User{
		Name:      service.UserName,
		UserId:    id,
		Telephone: service.Phone,
	}

	for _, v := range global.UserBackgroundurls {
		newUser.BackgroundPhoto = v
	}
	for _, v := range global.HeadPhotourls {
		newUser.HeaderPhoto = v
		break
	}
	// newUser.BackgroundPhoto = global.UserBackgroundurl
	// for _, t := range global.HeadPhotosMap {
	// 	newUser.HeaderPhoto = t
	// 	break
	// }
	err = dao.AddUser(newUser, newAccount) // 拉取当前账户及其用户的信息，并存储到数据库
	if err != nil {
		global.Log.WithError(err).Error("register")
		return serializer.Response{
			ResultStatus: 1,
			ResultMsg:    "Failed to save",
		}
	}
	newClaim := middleware.UserClaims{Id: id, Phone: service.Phone}
	token, refresh_token := middleware.GenerateToken(&newClaim)
	Token := serializer.Token{
		AccessToken:  token,
		RefreshToken: refresh_token,
	}
	return serializer.UserLoginResponse{
		ResultStatus: 0,
		ResultMsg:    "Success",
		Token:        Token,
		UserId:       strconv.FormatInt(id, 10),
	}
}

func (service UserService) Login(c *gin.Context) serializer.UserLoginResponse {
	service.Password = middleware.Md5Crypt(service.Password, "xueyigou")
	if !IdentityVerify(service.Phone, service.Password) {
		return serializer.UserLoginResponse{
			ResultStatus: 1,
			ResultMsg:    "Phone or password failed",
		}
	} else {
		userId := dao.GetUidByPhone(service.Phone)
		newClaim := middleware.UserClaims{Id: userId, Phone: service.Phone}
		token, refresh_token := middleware.GenerateToken(&newClaim)
		Token := serializer.Token{
			AccessToken:  token,
			RefreshToken: refresh_token,
		}
		return serializer.UserLoginResponse{
			ResultStatus: 0,
			ResultMsg:    "Success",
			UserId:       strconv.FormatInt(userId, 10),
			Token:        Token,
		}
	}
}

func LoginViaCode(phone string, code string) interface{} {
	exist := dao.AccountExistPhoneNum(phone)
	if exist {
		return serializer.BuildErrorResponse(e.ErrorNotExistUser)
	}
	err := VerifSmsCodeWithPhone(phone, code, global.TemplateCodeForLogin)
	if err != nil {
		return serializer.BuildErrorResponse(e.ErrorCode)
	}

	userId := dao.GetUidByPhone(phone)
	newClaim := middleware.UserClaims{Id: userId, Phone: phone}
	token, refresh_token := middleware.GenerateToken(&newClaim)
	Token := serializer.Token{
		AccessToken:  token,
		RefreshToken: refresh_token,
	}
	return serializer.UserLoginResponse{
		ResultStatus: 0,
		ResultMsg:    "Success",
		UserId:       strconv.FormatInt(userId, 10),
		Token:        Token,
	}
}

func IdentityVerify(phone string, password string) bool {
	//fmt.Println(phone, password, "sadasdas")
	account := dao.GetAccount(phone, password)
	//fmt.Println(account.UserName, account.PassWord)
	return account.Phone == phone && account.PassWord == password
}
func IdentifyPhonenumber(phone string, userid int64) bool {
	account, err := dao.GetAccountById(userid)
	return account.Phone == phone && err == nil
}
func (service UserService) GetFans(c *gin.Context) serializer.UserFansResponse {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	fansList := dao.GetUserFans(int64(userId))
	var userList []serializer.UserAttention
	for _, user := range fansList {
		userAttention := serializer.UserAttention{
			Id:          user.UserId,
			Name:        user.Name,
			Sign:        user.Sign,
			HeaderPhoto: user.HeaderPhoto,
		}
		userList = append(userList, userAttention)
	}
	return serializer.UserFansResponse{
		ResultStatus: 1,
		ResultMsg:    "Success",
		UserList:     userList,
	}
}

func (service UserService) GetAttention(c *gin.Context) serializer.UserAttentionResponse {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	attentionList := dao.GetUserAttention(int64(userId))
	var userList []serializer.UserAttention
	for _, user := range attentionList {
		userAttention := serializer.UserAttention{
			Id:          user.UserId,
			Name:        user.Name,
			Sign:        user.Sign,
			HeaderPhoto: user.HeaderPhoto,
		}
		userList = append(userList, userAttention)
	}
	return serializer.UserAttentionResponse{
		ResultStatus: 1,
		ResultMsg:    "Success",
		UserList:     userList,
	}
}

func GetUserInfo(userId, visitorId int64) interface{} {
	userInfo := dao.GetUserInfo(userId)
	followCount := dao.GetUserFollowCount(userId)
	followerCount := dao.GetUserFollowerCount(userId)
	isFollow := dao.GetIsFollow(userId, visitorId)
	work_cnt, like_cnt, err := dao.GetUserWorkCountAndLikes(userId)
	if err != nil {
		global.Log.WithError(err).Error("GetUserInfo %d", userId)
		return serializer.BuildFailResponse("internal err")
	}
	return serializer.BuildUserInfoResponse(userInfo, followCount, followerCount, isFollow, work_cnt, like_cnt)
}

type UserAttentionService struct {
	UserId     int64 `json:"user_id,omitempty"`
	OtherId    int64 `json:"other_id,omitempty"`
	ActionType int   `json:"action_type,omitempty"`
}

func (service *UserAttentionService) ParseAttentionAction(c *gin.Context) serializer.Response {
	if service.ActionType == 1 {
		dao.AddAttention(service.UserId, service.OtherId)
		dao.AddFan(service.UserId, service.OtherId)
		return serializer.Response{ResultStatus: 1, ResultMsg: "Success"}
	} else if service.ActionType == 2 {
		dao.DeleteAttention(service.UserId, service.OtherId)
		dao.DeleteFan(service.UserId, service.OtherId)
		return serializer.Response{ResultStatus: 1, ResultMsg: "Success"}
	} else {
		return serializer.Response{ResultStatus: 0, ResultMsg: "参数传入错误"}
	}
}

type UserGoodsService struct {
	UserId   int64 `json:"user_id"`
	ActionId int   `json:"action_id"`
}

type GoodsReturnService struct {
	UserId int64
	WordId int64
}

type ResetKeyService struct {
	VerifyCode  string `json:"code"`
	NewPassword string `json:"newPassword"`
	PhoneNumber string `json:"phone"`
}

func ResetKey(userid int64, newpassword string, Phone string, Code string) serializer.Response {
	err := VerifSmsCodeWithPhone(Phone, Code, global.TemplateCodeForResetPassword)
	if err != nil {
		return serializer.BuildFailResponse("verify code check wrong")
	}
	if newpassword == "" || len(newpassword) < 6 || newpassword == "" {
		return serializer.Response{
			ResultStatus: 1,
			ResultMsg:    "length of password is too short",
		}
	}
	newpassword = middleware.Md5Crypt(newpassword, "xueyigou")
	//fmt.Println(userid, "newone is", newpassword)
	err = dao.SetNewPassword(userid, newpassword)
	if err != nil {
		return serializer.BuildFailResponse("Reset wrong")
	}
	return serializer.BuildSuccessResponse("reset password success")
}

type SetUserInfoService struct {
	UserName            string   `json:"user_name"`
	UserHeadPhoto       string   `json:"user_header_photo"`
	UserBackgroundPhoto string   `json:"user_background_photo"`
	UserGender          string   `json:"user_gender"`
	UserHometown        string   `json:"user_hometown"`
	UserProvince        string   `json:"user_province"`
	UserCity            string   `json:"user_city"`
	UserSign            string   `json:"user_sign"`
	UserGraduate        string   `json:"user_graduate"`
	UserSubject         string   `json:"user_subject"`
	Userdetail          string   `json:"user_detail"`
	UserQQ              string   `json:"user_qq"`
	UserWechat          string   `json:"user_wechat"`
	UserLabels          []string `json:"user_labels"`
}

func (service SetUserInfoService) SetUserInfo(userid int64) serializer.Response {
	if userName := dao.GetUserNameById(userid); userName != service.UserName {
		if exist := dao.AccountExist(service.UserName); !exist {
			return serializer.Response{
				ResultStatus: 1,
				ResultMsg:    "UserName already exist",
			}
		}
	}
	user, err := dao.GetUserById(userid)
	account, err := dao.GetAccountById(userid)
	labels := make([]models.UserLables, len(service.UserLabels))
	for t, i := range service.UserLabels {
		labels[t].Label = i
		labels[t].UserId = userid
	}
	newaccount := models.Account{
		UserName: service.UserName,
	}
	newuser := models.User{
		Name:            service.UserName,
		Sex:             service.UserGender,
		HeaderPhoto:     service.UserHeadPhoto,
		BackgroundPhoto: service.UserBackgroundPhoto,
		Wechat:          service.UserWechat,
		Hometown:        service.UserHometown,
		Province:        service.UserProvince,
		Sign:            service.UserSign,
		City:            service.UserCity,
		Graduate:        service.UserGraduate,
		Subject:         service.UserSubject,
		QQ:              service.UserQQ,
		Detail:          service.Userdetail,
	}
	err = dao.SetUserAccount(account, newaccount)
	err = dao.SetUserInfo(user, newuser, labels)
	if err != nil {
		return serializer.BuildFailResponse("set failed")
	}
	return serializer.BuildSuccessResponse("set success")
}
func GetUserDetailInfo(userid int64) serializer.GetUserInfoResponse {
	user, _ := dao.GetUserById(userid)
	lables := dao.GetUserLables(user)
	return serializer.BuildGetUserDetailInfoResponse(user, lables)
}

// type ConnectCreatorService struct{
// 	UserId int64

// }
// func (service ConnectCreatorService)GetConnectCreator(c *gin.Context) serializer.Response{

// }

func GetAllAddress(user_id int64) interface{} {
	user, err := dao.GetUserById(user_id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return serializer.BuildUserDoesNotExitResponse("get address")
	}
	addresses := dao.GetUserAllAddress(user)
	return serializer.BuildGetAddress(addresses)
}

func PostUserAddress(Claim *middleware.UserClaims, address *models.Useraddress) interface{} {
	user, err := dao.GetUserById(Claim.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return serializer.BuildUserDoesNotExitResponse("post address")
	}
	dao.PostUserAddress(user, address)
	return serializer.BuildPostAddress()
}

func GetIndentity(user_id int64, Claim *middleware.UserClaims) interface{} {
	//TODO：在此处需要检测是否userid为token解析出的id，不是混淆数据
	if Claim.Id == user_id {
		_, err := dao.GetUserById(user_id)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.BuildUserDoesNotExitResponse("User don't exist")
		}
		indentity, result := dao.GetIndentity(user_id)
		//fmt.Println(*indentity)
		if !result {
			return serializer.BuildFailResponse("no Indentity exist")
		}
		return serializer.BuildGetIndentityResponse(*indentity)
	}
	return serializer.BuildFailResponse("token error")
}

type Indentityservice struct {
	UserID       int64
	UserPassword string  `json:"user_password"` // 用户密码
	Verifycode   string  `json:"verifycode"`    // 验证码
	UserTelenum  string  `json:"user_telenum"`
	UserIdentity string  `json:"user_identity"`
	UserIdnum    string  `json:"user_idnum"`
	UserName     string  `json:"user_name"`
	UserPhoto    *string `json:"user_photo,omitempty"`
}

func PostIndentity(indentity *Indentityservice, Claim *middleware.UserClaims) interface{} {
	//TODO: 在此处需要判断是否有已经挂起的资质审核
	New := models.Indentity{
		UserId:       Claim.Id,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		UserTelenum:  indentity.UserTelenum,
		UserIdentity: indentity.UserIdentity,
		UserIdnum:    indentity.UserIdnum,
		UserName:     indentity.UserName,
		UserPhoto:    indentity.UserPhoto,
	}
	account, err := dao.GetAccountById(Claim.Id)
	if New.UserTelenum != account.Phone {
		if !dao.AccountExistPhoneNum(New.UserTelenum) {
			return serializer.BuildFailResponse("phonenumber has been used")
		}
	}
	//真实姓名
	// if New.UserName != user.Name {
	// 	if dao.AccountExist(New.UserName) {
	// 		return serializer.BuildFailResponse("Name has been used")
	// 	}
	// }
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return serializer.BuildUserDoesNotExitResponse("post indentity")
	}
	err = VerifSmsCodeWithPhone(New.UserTelenum, indentity.Verifycode, global.TemplateCodeForIndentity)
	if err != nil {
		return serializer.BuildFailResponse("phone check wrong")
	}
	indentity.UserPassword = middleware.Md5Crypt(indentity.UserPassword, "xueyigou")
	if !IdentityVerify(account.Phone, indentity.UserPassword) {
		return serializer.BuildFailResponse("password failed")
	}
	if a, _ := dao.GetIndentity(Claim.Id); a.UserName == "" {
		dao.PostIndentity(&New, Claim.Id)
		return serializer.BuildPostIndentityResponse()
	}
	dao.UpdateIndentity(&New, Claim.Id)
	fmt.Println(New)
	return serializer.BuildPostIndentityResponse()
}

func PostPeople(people *models.PeopleMid, Claim *middleware.UserClaims) interface{} {
	_, err := dao.GetUserById(Claim.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return serializer.BuildUserDoesNotExitResponse("post people")
	}
	dao.PostPeople(people)
	return serializer.BuildPostPeopleResponse()
}

func PostPeopleMid(peopleId int64, ispassed int64) interface{} {
	if ispassed == 1 {
		peopleMid, err := dao.GetPeopleMidById(peopleId)
		if err != nil {
			return serializer.BuildFailResponse("Get False")
		}
		people := models.People{
			PeopleDetails:    peopleMid.PeopleDetails,
			PeopleMajorskill: peopleMid.PeopleMajorskill,
			PeopleType:       peopleMid.PeopleType,
			PeopleMax:        peopleMid.PeopleMax,
			PeopleMin:        peopleMid.PeopleMin,
			PeopleIntroduce:  peopleMid.PeopleIntroduce,
			PeopleIp:         peopleMid.PeopleIp,
			PeopleId:         peopleId,
			PeopleStatus:     0,
		}
		for _, subject := range peopleMid.PeopleSubject {
			people.PeopleSubject = append(people.PeopleSubject, models.PeopleSubjectItem{PeopleId: peopleId, Item: subject.Item})
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.BuildUserDoesNotExitResponse("post people")
		}
		dao.PostPeopleMid(&people)
	}
	if err := dao.DeletePeopleMid(peopleId); err == nil {
		return serializer.BuildSuccessResponse("Delete people success")
	} else {
		return serializer.BuildFailResponse("Delete people failed")
	}
	return serializer.BuildPostPeopleResponse()
}

func GetPeopleForMe(Id int64) interface{} {
	people, err := dao.GetPeopleById(Id)
	if err != nil {
		return serializer.BuildFailResponse("Get False")
	}
	return serializer.BuildMyPeople(*people)
}

func PostOrderTask(user_id int64, task_id int64) interface{} {
	err := dao.OrderTask(user_id, task_id)
	if err != nil {
		return serializer.BuildFailResponse(err.Error())
	}
	return serializer.BuildOrderTaskResponse()
}

func FinishTask(user_id int64, task_id int64) interface{} {
	err := dao.FinishTask(user_id, task_id)
	if err != nil {
		return serializer.BuildFailResponse(err.Error())
	}
	return serializer.BuildFinishTaskResponse()
}

func GetOrderTaskList(userId int64) interface{} {
	taskIds, err := dao.GetOrderTaskList(userId)
	if err != nil {
		return serializer.BuildFailResponse(err.Error())
	}
	return serializer.BuildGetOrderTaskListResponse(taskIds)
}

func PostOrderWork(userId int64, workId int64) interface{} {
	//if err := dao.PostOrderWork(userId, workId); err != nil {
	//	return serializer.BuildFailResponse(err.Error())
	//}
	return serializer.BuildSuccessResponse("Order Work")
}

func GetOrderWorkList(userId int64) interface{} {
	workIds, err := dao.GetOrderWorkList(userId)
	if err != nil {
		return serializer.BuildFailResponse(err.Error())
	}
	return serializer.BuildGetOrderWorkListResponse(workIds)
}

func DelOrderItem(uId int64, typeId int, oId int64) interface{} {
	//if typeId == 0 {
	//if err := dao.DelWorkItem(uId, oId); err != nil {
	//	return serializer.BuildFailResponse(err.Error())
	//}
	//} else if typeId == 1 {
	//	if err := dao.DelTaskItem(uId, oId); err != nil {
	//		return serializer.BuildFailResponse(err.Error())
	//	}
	//}
	ret := serializer.BuildSuccessResponse("delete order successfully")
	return ret
}
func Feedback(feedbcak models.UserFeedback) interface{} {
	err := dao.PostFeedback(feedbcak)
	if err != nil {
		return serializer.BuildFailResponse("post fail")
	} else {
		return serializer.BuildSuccessResponse("post success")
	}
}
