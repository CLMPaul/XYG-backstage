package errcodes

import (
	"net/http"
	"xueyigou_demo/internal/errcode"
)

var (
	// 用户 IP 与会话不一致
	SessionIPChanged = errcode.NewHttpError(http.StatusUnauthorized, "session ip changed")
	// 由于用户被禁用、用户密码修改等原因，强制下线
	SessionExpired = errcode.NewHttpError(http.StatusUnauthorized, "session expired")
	// csrf token 无效
	InvalidCsrfToken = errcode.NewHttpError(http.StatusUnauthorized, "invalid csrf token")
)
