package dao

import (
	"context"
	"strconv"
	"xueyigou_demo/cache"
)

func WorkView(workId int64) uint64 {
	countStr, _ := cache.GetClient().Get(context.Background(), cache.WorkViewKey(workId)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 商品游览
func AddWorkView(workId int64) {
	// 增加商品浏览数
	cache.GetClient().Incr(context.Background(), cache.WorkViewKey(workId))

}

// View 获取浏览数
func WelfareView(welfareId int64) uint64 {
	countStr, _ := cache.GetClient().Get(context.Background(), cache.WelfareViewKey(welfareId)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 公益活动游览
func AddWelfareView(welfareId int64) {
	// 增加公益活动浏览数
	cache.GetClient().Incr(context.Background(), cache.WelfareViewKey(welfareId))
	// 增加排行点击数
	//cache.GetClient().ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(product.ID)))
}

func TaskView(taskId int64) uint64 {
	countStr, _ := cache.GetClient().Get(context.Background(), cache.TaskViewKey(taskId)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 任务浏览
func AddTaskView(taskId int64) {
	// 增加任务浏览数
	cache.GetClient().Incr(context.Background(), cache.TaskViewKey(taskId))
	// 增加排行点击数
	//cache.GetClient().ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(product.ID)))
}
