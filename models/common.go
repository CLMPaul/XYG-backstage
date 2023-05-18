package models

type ObjectType int64

const (
	PWORK    string = "wo_"
	PTASK    string = "ta_"
	PWELFARE string = "we_"
	PEVENT   string = "ev_"
)

const (
	WORK ObjectType = iota
	TASK
	WELFARE
	EVENT
	//TODO:
)

type UserBasicInfo struct {
	UserId      int64  `json:"user_id"`
	Name        string `json:"user_name"`
	HeaderPhoto string `json:"header_photo"`
}
