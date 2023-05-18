package serializer

type SSOLoginResponse struct {
	ResultStatus int32  `json:"result_status"`
	ResultMsg    string `json:"result_msg,omitempty"`
	UserId       int64  `json:"user_id,omitempty"`
	Token        Token  `json:"token,omitempty"`
}
