package dao

import (
	"xueyigou_demo/cache"
	"xueyigou_demo/db"
	"xueyigou_demo/models"
)

// maybe not used
func GetHonorList() models.HonorList {
	var honor_list models.HonorList
	//honor_id := CreateHonorList()
	db.DB.Preload("Users").Order("create_at desc").First(&honor_list)
	return honor_list
}

func CreateHonorList() {
	//global.Log.Info("create honor list")
	var honor_list []models.HonorList
	db.DB.Model(&models.User{}).Order("contribution desc, medals desc").Limit(10).Find(&honor_list)
	cache.CacheHonorList(honor_list)
	return
}
