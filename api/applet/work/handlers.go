package work

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"xueyigou_demo/global"
	"xueyigou_demo/internal/web"
	"xueyigou_demo/middleware"
	"xueyigou_demo/models"
	"xueyigou_demo/proto"
	"xueyigou_demo/service"
)

func getPaged(c *gin.Context) {
	var req proto.WorkRequest
	err := c.BindQuery(&req)
	if err != nil {
		web.BadRequestResponse(c, err.Error())
		return
	}
	state := 1
	req.State = &state
	userClaim, exist := c.Get("userClaim")
	if !exist {
		web.BadRequestResponse(c, "user claim not in token")
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	req.UserId = claim.Id
	data, err := service.WorkService.GetPaged(req)
	if err != nil {
		web.HandleError(c, err)
		return
	}
	web.SuccessResponse(c, data)
}

func getPagedByUser(c *gin.Context) {
	var req proto.WorkRequest
	err := c.BindQuery(&req)
	if err != nil {
		web.BadRequestResponse(c, err.Error())
		return
	}
	userClaim, exist := c.Get("userClaim")
	if !exist {
		web.BadRequestResponse(c, "user claim not in token")
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	data, err := service.WorkService.GetPagedByUser(req, claim.Id)
	if err != nil {
		web.HandleError(c, err)
		return
	}
	web.SuccessResponse(c, data)
}

func getInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		web.BadRequestResponse(c, "id is empty")
		return
	}
	userClaim, exist := c.Get("userClaim")
	if !exist {
		web.BadRequestResponse(c, "user claim not in token")
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	data, err := service.WorkService.FindById(id, claim.Id)
	if err != nil {
		web.HandleError(c, err)
		return
	}
	web.SuccessResponse(c, data)
}

func add(c *gin.Context) {
	var form proto.WorkForm
	err := c.BindJSON(&form)
	if err != nil {
		web.BadRequestResponse(c, err.Error())
		return
	}
	w := models.Work{
		Introduce: global.IllegalWords.Replace(form.Introduce, '*'),
		Title:     form.Title,
		TypeID:    form.TypeID,
	}
	for index, pic := range form.WorkPicture {
		if index == 0 {
			if pic == "" {
				for _, v := range global.WorkPhotourls {
					w.CoverPicture = v
					break
				}
			} else {
				w.CoverPicture = pic
			}

			continue
		}
		w.PicturesUrlList = append(w.PicturesUrlList, models.WorkPicturesUrl{Url: pic})
	}
	userClaim, exist := c.Get("userClaim")
	if !exist {
		web.BadRequestResponse(c, "user claim not in token")
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	w.PostUserId = claim.Id
	err = service.WorkService.Add(&w)
	if err != nil {
		web.HandleError(c, err)
		return
	}
	web.SuccessResponse(c, w.ID)
}

func _delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		web.BadRequestResponse(c, "id is empty")
		return
	}
	err := service.WorkService.Delete(id)
	if err != nil {
		web.HandleError(c, err)
		return
	}
	web.SuccessResponse(c, id)
}

func update(c *gin.Context) {
	var form proto.WorkForm
	err := c.BindJSON(&form)
	if err != nil {
		web.BadRequestResponse(c, err.Error())
		return
	}
	w := models.Work{
		Introduce: global.IllegalWords.Replace(form.Introduce, '*'),
		Title:     form.Title,
		TypeID:    form.TypeID,
	}
	for index, pic := range form.WorkPicture {
		if index == 0 {
			if pic == "" {
				for _, v := range global.WorkPhotourls {
					w.CoverPicture = v
					break
				}
			} else {
				w.CoverPicture = pic
			}

			continue
		}
		w.PicturesUrlList = append(w.PicturesUrlList, models.WorkPicturesUrl{Url: pic})
	}
	err = service.WorkService.Update(&w)
	if err != nil {
		web.HandleError(c, err)
		return
	}
	web.SuccessResponse(c, w.ID)
}

func TakeDown(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.BadRequestResponse(c, err.Error())
		return
	}
	if id == 0 {
		web.BadRequestResponse(c, "id is empty")
		return
	}
	err = service.WorkService.UpdateState(id, 3)
	if err != nil {
		web.HandleError(c, err)
		return
	}
	web.SuccessResponse(c, id)
}
