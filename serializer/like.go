package serializer

type LikeListResponse struct {
	Response
	LikeList []int64
}

func BuildLikeListResponse(likelist []int64) LikeListResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get like list success",
	}
	likeListResponse := LikeListResponse{
		response,
		likelist,
	}
	return likeListResponse
}
