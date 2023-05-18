package service

import (
	"xueyigou_demo/cache"
	"xueyigou_demo/serializer"
)

func GetHonorList() interface{} {
	honor_list, err := cache.CacheGetHonorList()
	if err != nil {
		return serializer.BuildFailResponse("GetHonorList")
	}
	response := serializer.BuildHonorListResponse(honor_list)
	return response
}
