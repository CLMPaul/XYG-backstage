package serializer

import (
	"net/http"
	"xueyigou_demo/models"
)

type AppletActivityPagedResp struct {
	Response
	Data  []models.Event `json:"data"`
	Total int64          `json:"total"`
}

func BuildAppletActivityPagedResp(data []models.Event, total int64, err error) AppletActivityPagedResp {
	resp := AppletActivityPagedResp{
		Response: Response{
			ResultStatus: http.StatusOK,
		},
		Data:  data,
		Total: total,
	}

	if err != nil {
		resp.Response.Error = err.Error()
	}

	return resp
}
