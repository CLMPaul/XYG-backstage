package cache

import (
	"fmt"
	"strconv"
)

func WelfareViewKey(id int64) string {
	return fmt.Sprintf("view:welfare:%s", strconv.Itoa(int(id)))
}

func WorkViewKey(id int64) string {
	return fmt.Sprintf("view:work:%s", strconv.Itoa(int(id)))
}

func TaskViewKey(id int64) string {
	return fmt.Sprintf("view:task:%s", strconv.Itoa(int(id)))
}

func Comment_l1LikeKey(id int64) string {
	return fmt.Sprintf("likes:comment_l1:%s", strconv.Itoa(int(id)))
}

func Comment_l1LikeUserKey(obejctid int64) string {
	return fmt.Sprintf("like:commentl1-userid:%s", strconv.Itoa(int(obejctid)))
}
func Comment_l2LikeKey(id int64) string {
	return fmt.Sprintf("likes:comment_l2:%s", strconv.Itoa(int(id)))
}
func Comment_l2LikeUserKey(obejectid int64) string {
	return fmt.Sprintf("like:commentl2-userid:%s", strconv.Itoa(int(obejectid)))
}

func TaskLikeKey(id int64) string {
	return fmt.Sprintf("likes:task:%s", strconv.Itoa(int(id)))
}
func TaskLikeUserKey(obejectid int64) string {
	return fmt.Sprintf("like:task-userid:%s", strconv.Itoa(int(obejectid)))
}

func WorkLikeKey(id int64) string {
	return fmt.Sprintf("likes:work:%s", strconv.Itoa(int(id)))
}
func WorkLikeUserKey(obejectid int64) string {
	return fmt.Sprintf("like:work-userid:%s", strconv.Itoa(int(obejectid)))
}

func WelfareLikeKey(id int64) string {
	return fmt.Sprintf("likes:welfare:%s", strconv.Itoa(int(id)))
}
func WelfareLikeUserKey(obejectid int64) string {
	return fmt.Sprintf("like:welfare-userid:%s", strconv.Itoa(int(obejectid)))
}
func EventLikeKey(id int64) string {
	return fmt.Sprintf("likes:event:%s", strconv.Itoa(int(id)))
}
func EventLikeUserKey(obejectid int64) string {
	return fmt.Sprintf("like:event-userid:%s", strconv.Itoa(int(obejectid)))
}
