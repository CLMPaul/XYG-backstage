package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	TaskMax   int64          `json:"task_max"`  // 最高价格
	TaskMin   int64          `json:"task_min"`  // 最低价格
	TaskName  string         `json:"task_name"` // 任务名称
	//for me
	TaskCover     string            `json:"task_picture,omitempty"` // 任务封面图片url
	TaskStatus    int               `json:"task_status,omitempty"`  // 任务状态码，0-未接单，1-已接单，未完成，2-已完成
	TaskSubject   []TaskSubjectItem `json:"task_subject,omitempty"` // 任务科目
	TaskIntroduce string            `json:"task_introduce"`
	TaskMessage   string            `json:"task_message,omitempty"`  // 任务留言
	TaskNumber    string            `json:"task_number,omitempty"`   // 任务编号
	TaskProgress  string            `json:"task_progress,omitempty"` // 任务进度
	TaskRequire   string            `json:"task_require,omitempty"`  // 任务要求
	//for visitor
	TaskBedonated   string `json:"task_bedonated,omitempty"`  // 任务被捐赠对象
	TaskCollection  int64  `json:"task_collection,omitempty"` // 任务被收藏数
	TaskDetails     string `json:"task_details,omitempty"`    // 任务简介
	TaskDonation    string `json:"task_donation,omitempty"`   // 任务捐赠对象
	TaskForward     int64  `json:"task_forward,omitempty"`    // 任务转发数
	TaskLike        int64  `json:"task_like,omitempty"`       // 任务点赞量
	TaskPercent     int64  `json:"task_percent,omitempty"`    // 任务捐赠百分比
	TaskPrice       int64  `json:"task_price,omitempty"`      // 任务价格
	TaskView        int64  `json:"task_view,omitempty"`       // 任务被浏览量
	TaskType        string
	PostUserId      int64
	PostUser        User              `gorm:"foreignKey:PostUserId;references:UserId"`
	PicturesUrlList []TaskPicturesUrl `json:"pictures_url_list"` // 图像url
	LikeUser        []*User           `gorm:"many2many:like_task"`
	CollectUser     []*User           `gorm:"many2many:task_collection"`
	OrderUser       []*User           `gorm:"many2many:order_task"`
}

type TaskPicturesUrl struct {
	ID     int64
	TaskId int64
	Url    string
}

type TaskSubjectItem struct {
	ID     int64
	TaskId int64
	Item   string
}
type Candidate struct {
	TaskId int64 `json:"task_id"`
	UserId int64 `json:"user_id"`
}

//func (task *Task) View() uint64 {
//	countStr, _ := cache.GetClient().Get(cache.WelfareViewKey(task.ID)).Result()
//	count, _ := strconv.ParseUint(countStr, 10, 64)
//	return count
//}

// AddView 任务浏览
//func (task *Task) AddView() {
//	// 增加任务浏览数
//	cache.GetClient().Incr(cache.TaskViewKey(task.ID))
//	// 增加排行点击数
//	//cache.GetClient().ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(product.ID)))
//}
