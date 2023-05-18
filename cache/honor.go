package cache

import (
	"encoding"
	"fmt"
	"time"
	"xueyigou_demo/global"
	"xueyigou_demo/models"
)

var _ encoding.BinaryMarshaler = new(models.HonorList)
var _ encoding.BinaryUnmarshaler = new(models.HonorList)

func honorKey(id int) string {
	return fmt.Sprintf("honor_list:%d", id)
}

func CacheHonorList(honor_list []models.HonorList) error {
	//global.Log.WithField("honorlist", honor_list).Info("cache")
	for id, item := range honor_list {
		if err := SetByTTL(honorKey(id), &item, time.Minute*30); err != nil {
			global.Log.WithError(err).Error("cache")
			return err
		}
	}
	return nil
}

func CacheGetHonorList() ([]models.HonorList, error) {
	honor_list := []models.HonorList{}

	for i := 0; i < 10; i++ {
		var item models.HonorList
		if err := GetResult(honorKey(i)).Scan(&item); err != nil {
			global.Log.WithError(err).Error("cache")
			return nil, err
		} else {
			honor_list = append(honor_list, item)
		}
	}
	//global.Log.WithField("honor list", honor_list).Info("cache")
	return honor_list, nil
}
