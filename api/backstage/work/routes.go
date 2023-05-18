package work

import (
	"github.com/gin-gonic/gin"
)

// InitRouterBackstageWork 后台管理-作品
func SetupRouter(apiRouter gin.IRouter) {
	//apiRouter.GET("list", api.WorkSliceList)       // 已审核列表
	//apiRouter.GET("auditlist", api.GetWorkMidList) // 待审核列表
	//apiRouter.POST("audit", api.PostWorkMid)       // 审核

	apiRouter.GET("paged", getPaged)
	apiRouter.GET(":id", getInfo)
	apiRouter.PUT("audit/approved/:id", auditApproved)
	apiRouter.PUT("audit/not_approved/:id", auditNotApproved)
	apiRouter.PUT("audit/approved/batch", auditApprovedBatch)
	apiRouter.PUT("audit/not_approved/batch", auditNotApprovedBatch)
	apiRouter.PUT("audit/take_down/:id", auditTakeDown)
	apiRouter.PUT("audit/take_down/batch", auditTakeDownBatch)
	apiRouter.DELETE(":id", _delete)
}
