package dao

import (
	"errors"
	"xueyigou_demo/db"
	"xueyigou_demo/models"
	"xueyigou_demo/serializer"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//	Add

func AddUser(user models.User, account models.Account) error {
	//mu.Lock()
	//defer mu.Unlock()
	return db.DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Model(&models.User{}).Create(&user).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.Model(&models.Account{}).Create(&account).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
}

// Get
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := db.DB.Table("users").Find(&users).Error
	return users, err
}
func GetUserById(userId int64) (models.User, error) {
	// 从db中获取user
	var user models.User
	err := db.DB.Where("user_id = ?", userId).First(&user).Error //find无匹配记录不报错 所以first
	return user, err
}

func GetAccountById(userId int64) (models.Account, error) {
	var account models.Account
	err := db.DB.Where("id = ?", userId).Find(&account).Error
	return account, err
}

// func GetUserByName(username string) bool {
//
//		err := db.DB.Where("name = ?", username).First(&models.User{}).Error //find无匹配记录不报错 所以first
//		if err == nil {
//			return true
//		}
//		return false
//	}
func GetUidByPhone(phone string) int64 {
	var userId int64
	if err := db.DB.Raw("select id from accounts where phone=?", phone).Scan(&userId).Error; err != nil {
		panic(err)
	}
	return userId
}
func AccountExist(username string) bool {
	var account models.Account
	var count int64
	db.DB.Table("accounts").Where("user_name = ?", username).Find(&account).Count(&count)
	return count == 0
}

func AccountExistPhoneNum(phone_num string) bool {
	var account models.Account
	var count int64
	db.DB.Table("accounts").Where("phone = ?", phone_num).Find(&account).Count(&count)
	return count == 0
}
func GetUserNameById(userId int64) string {
	var name string
	if err := db.DB.Table("accounts").Select("user_name").Where("id = ?", userId).Find(&name).Error; err != nil {
		panic(err)
	}
	return name
}
func GetAccount(phone string, password string) models.Account {
	var account models.Account
	db.DB.Table("accounts").Where("phone = ?", phone).Where("pass_word = ?", password).Find(&account)
	return account
}

func GetUsersAmount() int64 {
	// 从db中获取user的数量(ID)
	var count int64
	db.DB.Model(&models.User{}).Where("nick_name != ?", "").Count(&count)
	return count
}

func GetUserAttention(userId int64) []models.User {
	var user models.User
	var attentions []models.User
	if err := db.DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		if err = db.DB.Model(&user).Association("UserFollows").Find(&attentions); err != nil {
			panic(err)
		}
	}
	return attentions
}

func AddAttention(userId int64, otherId int64) error {
	var user models.User
	userAttention, err := GetUserById(otherId)
	if err != nil {
		return err
	}
	if err := db.DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.DB.Model(&user).Association("UserFollows").Append(&userAttention)
		return err
	}
	return nil
}
func AddFan(userId int64, otherId int64) error {
	var user models.User
	userFan, err := GetUserById(userId)
	if err != nil {
		return err
	}
	if err := db.DB.Where("user_id = ?", otherId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.DB.Model(&user).Association("UserFollowers").Append(&userFan)
		return err
	}
	return nil
}

func DeleteAttention(userId int64, otherId int64) error {
	var user models.User
	userAttention, err := GetUserById(otherId)
	if err != nil {
		return err
	}
	if err := db.DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.DB.Model(&user).Association("UserFollows").Delete(&userAttention)
		return err
	}
	return nil
}
func DeleteFan(userId int64, otherId int64) error {
	var user models.User
	userFan, err := GetUserById(userId)
	if err != nil {
		return err
	}
	if err := db.DB.Where("user_id = ?", otherId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.DB.Model(&user).Association("UserFollowers").Delete(&userFan)
		return err
	}
	return nil
}

func GetUserFans(userId int64) []models.User {
	var fansList []models.User
	var user models.User
	if err := db.DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		if err = db.DB.Model(&user).Association("UserFollowers").Find(&fansList); err != nil {
			panic(err)
		}
	}
	return fansList
}
func GetUserInfo(userId int64) models.User {
	var user models.User
	if err := db.DB.Where("user_id = ?", userId).Preload("LabelList").Find(&user).Error; err != nil {
		panic(err)
	}
	return user
}
func GetUserFollowCount(userId int64) int64 {
	var count int64
	if err := db.DB.Raw("select count(user_user_id = ? or null) from user_follows", userId).Find(&count).Error; err != nil {
		panic(err)
	}
	return count
}
func GetUserFollowerCount(userId int64) int64 {
	var count int64
	if err := db.DB.Raw("select count(user_user_id = ? or null) from user_followers", userId).Find(&count).Error; err != nil {
		panic(err)
	}
	return count
}

func GetIsFollow(userId, visitorId int64) bool {
	var count int64
	if err := db.DB.Raw("select count(user_user_id = ? and user_follow_user_id = ? or null) from user_follows", visitorId, userId).Find(&count).Error; err != nil {
		panic(err)
	}
	if count > 0 {
		return true
	}
	return false
}

func GetGoodsList(userId int64) []int64 {
	var goodsList []int64
	if err := db.DB.Raw("select goods_id from user_goods where user_id=?", userId).Scan(&goodsList).Error; err != nil {
		panic(err)
	}
	return goodsList
}

func GetUserWorkCountAndLikes(user_id int64) (like_cnt, work_cnt int64, err error) {
	var works []models.Work
	like_cnt = 0
	if err = db.DB.Where("post_user_id = ?", user_id).Find(&works).Error; err != nil {
		return 0, 0, err
	}
	for _, work := range works {
		like_cnt += work.Like
	}
	return like_cnt, int64(len(works)), err
}

// xueyigou/task/info 仅返回此路由需要的
func GetUserAllTasks(user_id int64, task_id int64) []models.Task {
	var tasks []models.Task
	if err := db.DB.Select("id", "task_name", "task_max", "task_min", "task_cover", "post_user_id").Where(
		"post_user_id = ?", user_id).Preload("PicturesUrlList").Find(&tasks).Error; err != nil {
		panic(err)
	}
	for index, task := range tasks {
		if task.ID == task_id {
			tasks = append(tasks[:index], tasks[index+1:]...)
		}
	}
	return tasks
}
func GetUserLables(user models.User) []models.UserLables {
	var lables []models.UserLables
	if err := db.DB.Model(&user).Association("LabelList").Find(&lables); err != nil {
		panic(err)
	}
	return lables
}

// Post修改数据库内容
func SetNewPassword(userid int64, password string) error {
	err := db.DB.Table("accounts").Where("id = ?", userid).Update("pass_word", password).Error
	return err
}
func SetUserInfo(user models.User, newuser models.User, labels []models.UserLables) error {
	err := db.DB.Model(&user).Updates(newuser).Error
	err = db.DB.Model(&user).Association("LabelList").Clear()
	err = db.DB.Model(&user).Association("LabelList").Replace(labels)
	return err
}
func SetUserAccount(account models.Account, newAccount models.Account) error {
	err := db.DB.Model(&account).Updates(newAccount).Error
	return err
}
func PostUserAddress(user models.User, address *models.Useraddress) {
	if err := db.DB.Model(&user).Association("Addresses").Append(address); err != nil {
		panic(err)
	}
}

func GetUserAllAddress(user models.User) []models.Useraddress {
	var addresses []models.Useraddress
	if err := db.DB.Model(&user).Association("Addresses").Find(&addresses); err != nil {
		panic(err)
	}
	return addresses
}

func PostIndentity(indentity *models.Indentity, Userid int64) {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Model(&models.Indentity{}).Create(indentity).Error; err != nil {
			// 返回任何错误都会回滚事务
			panic(err)
		}
		//if err := tx.Model(&models.Account{}).Where("id = ?", Userid).Update("user_real_name", indentity.UserName).Error; err != nil {
		//	panic(err)
		//}
		if err := tx.Model(&models.Account{}).Where("id = ?", Userid).Update("phone", indentity.UserTelenum).Error; err != nil {
			panic(err)
		}
		if err := tx.Model(&models.User{}).Where("user_id = ?", Userid).Update("telephone", indentity.UserTelenum).Error; err != nil {
			panic(err)
		}
		return nil
	})
	if err != nil {
		return
	}
}

func UpdateIndentity(indentity *models.Indentity, Userid int64) {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Account{}).Where("id = ?", Userid).Update("user_name", indentity.UserName).Error; err != nil {
			panic(err)
		}
		if err := tx.Model(&models.Account{}).Where("id = ?", Userid).Update("phone", indentity.UserTelenum).Error; err != nil {
			panic(err)
		}
		if err := tx.Model(&models.Indentity{}).Where("user_id = ?", Userid).Updates(*indentity).Error; err != nil {
			panic(err)
		}
		if err := tx.Model(&models.User{}).Where("user_id = ?", Userid).Update("telephone", indentity.UserTelenum).Error; err != nil {
			panic(err)
		}
		return nil
	})
	if err != nil {
		return
	}
}
func PostPeople(people *models.PeopleMid) {
	// Upsert功能，存在就更新，不存在就创建
	if err := db.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "people_id"}},
		UpdateAll: true,
	}).Create(&people).Error; err != nil {
		panic(err)
	}
	if err := db.DB.Model(&people).Where("people_id = ?", people.PeopleId).Association("PeopleSubject").
		Replace(people.PeopleSubject); err != nil {
		panic(err)
	}
}

func PostPeopleMid(people *models.People) {
	// Upsert功能，存在就更新，不存在就创建
	if err := db.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "people_id"}},
		UpdateAll: true,
	}).Create(&people).Error; err != nil {
		panic(err)
	}
	if err := db.DB.Model(&people).Where("people_id = ?", people.PeopleId).Association("PeopleSubject").
		Replace(people.PeopleSubject); err != nil {
		panic(err)
	}
}

func GetIndentity(user_id int64) (*serializer.IndentityItem, bool) {
	var indentity serializer.IndentityItem
	err := db.DB.Model(&models.Indentity{}).Where("user_id = ?", user_id).Find(&indentity).Error
	return &indentity, err == nil
}

func OrderTask(user_id int64, task_id int64) error {
	user, err := GetUserById(user_id)
	if err != nil {
		return err
	}
	var tasks []models.Task
	if err := db.DB.Model(&user).Association("OrderTask").Find(&tasks); err != nil {
		return err
	}
	for _, task := range tasks {
		if task.ID == task_id {
			return errors.New("have exist")
		}
	}
	//不更改status
	// updata := map[string]interface{}{"task_status": 1}
	// if err := db.DB.Table("tasks").Where("id = ?", task_id).Updates(updata).Error; err != nil {
	// 	return err
	// }
	if err := db.DB.Model(&user).Association("OrderTask").Append(&models.Task{ID: task_id}); err != nil {
		return nil
	}
	return nil
}

func FinishTask(user_id int64, task_id int64) error {
	updata := map[string]interface{}{"task_status": 2}
	if err := db.DB.Table("tasks").Where("id = ?", task_id).Updates(updata).Error; err != nil {
		return err
	}

	return nil
}

func GetOrderTaskList(userId int64) ([]int64, error) {
	var tasks []models.Task
	user, err := GetUserById(userId)
	if err != nil {
		panic(err)
	}
	if err = db.DB.Model(&user).Association("OrderTask").Find(&tasks); err != nil {
		return nil, err
	}
	var taskIds []int64
	for _, task := range tasks {
		taskIds = append(taskIds, task.ID)
	}
	return taskIds, nil
}

func GetUserByGitHubId(github_id int64) (*models.User, error) {
	user := &models.User{}
	if err := db.DB.Where("git_hub_id = ?", github_id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByWeChatId(wechat_id string) (*models.User, error) {
	user := &models.User{}
	if err := db.DB.Where("we_chat_id = ?", wechat_id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByQQId(qq_id string) (*models.User, error) {
	user := &models.User{}
	if err := db.DB.Where("qq_id = ?", qq_id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByPhoneNum(phone_num string) (*models.User, error) {
	user := &models.User{}
	if err := db.DB.Where("telephone = ?", phone_num).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByOpenid(openid string) (*models.User, error) {
	user := &models.User{}
	if err := db.DB.Where("app_let_openid = ?", openid).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByUnionid(unionid string) (*models.User, error) {
	user := &models.User{}
	if err := db.DB.Where("app_let_unionid = ?", unionid).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func UpdateUserInfoByStruct(user *models.User) error {
	if err := db.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderWorkList(userId int64) ([]int64, error) {
	var works []models.Work
	user, err := GetUserById(userId)
	if err != nil {
		panic(err)
	}
	if err = db.DB.Model(&user).Association("OrderWork").Find(&works); err != nil {
		return nil, err
	}
	var workIds []int64
	//for _, work := range works {
	//	workIds = append(workIds, work.ID)
	//}
	return workIds, nil
}

func PeopleChangeStatus(Id int64, status int) {
	if err := db.DB.Table("peoples").Where("people_id = ?", Id).Update("people_status", status).Error; err != nil {
		panic(err)
	}
}

func GetPeopleById(Id int64) (*models.People, error) {
	people := &models.People{}
	if err := db.DB.Where("people_id=?", Id).Preload(clause.Associations).First(people).Error; err != nil {
		return nil, err
	}
	return people, nil
}

func GetPeopleMidById(Id int64) (*models.PeopleMid, error) {
	people := &models.PeopleMid{}
	if err := db.DB.Where("people_id=?", Id).Preload(clause.Associations).First(people).Error; err != nil {
		return nil, err
	}
	return people, nil
}

func PostFeedback(feedback models.UserFeedback) error {
	err := db.DB.Create(&feedback).Error
	if err != nil {
		return err
	}
	return nil
}
