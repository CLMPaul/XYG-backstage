package api

import (
	"errors"
	"net/http"
	"strconv"
	"xueyigou_demo/dao"
	"xueyigou_demo/global"
	"xueyigou_demo/service"

	"github.com/gin-gonic/gin"
)

func UploadDefaultPic(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	f := form.File["pictures"]
	object_types := form.Value["object_type"]
	public_path, res := service.Uploadfile(f, c.Request.Host)
	if res != nil {
		c.JSON(http.StatusOK, res)
		return
	}
	for _, object_type := range object_types {
		switch object_type {
		case "task":
			for _, url := range public_path {
				global.TaskPhotourls[global.Worker.GetId()] = url
			}
		case "work":
			for _, url := range public_path {
				global.WorkPhotourls[global.Worker.GetId()] = url
			}
		case "userbackground":
			for _, url := range public_path {
				global.UserBackgroundurls[global.Worker.GetId()] = url
			}
		case "headphoto":
			for _, url := range public_path {
				global.HeadPhotourls[global.Worker.GetId()] = url
			}

		}
	}

	go dao.DefaultPicPersistence(object_types)
	c.JSON(http.StatusOK, gin.H{
		"result_status": 0,
		"reslut_msg":    "upload default pictures",
	})
}

type deleteForm struct {
	PicId      []int64 `json:"pic_id"`
	ObjectType string  `json:"object_type"`
}

func DeleteDefaultPic(c *gin.Context) {
	form := deleteForm{}
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	var del_urls []string
	switch form.ObjectType {
	case "task":
		del_urls = delDefaultPic(form.PicId, &global.TaskPhotourls)
	case "work":
		del_urls = delDefaultPic(form.PicId, &global.WorkPhotourls)
	case "userbackground":
		del_urls = delDefaultPic(form.PicId, &global.UserBackgroundurls)
	case "headphoto":
		del_urls = delDefaultPic(form.PicId, &global.HeadPhotourls)
	}
	res, err := service.Deletefile(del_urls)
	if err != nil {
		c.JSON(http.StatusOK, res)
		return
	}
	go dao.DefaultPicPersistence([]string{form.ObjectType})
	c.JSON(http.StatusOK, gin.H{
		"reslut_status": 0,
		"reslut_msg":    "DeleteDefaultPic",
	})
}

// return urls be deleted
func delDefaultPic(urls_id_2_del []int64, urls *map[int64]string) []string {
	cur := 0
	var del_urls []string
	for ; cur < len(urls_id_2_del); cur++ {
		for id, url := range *urls {
			if id == urls_id_2_del[cur] {
				delete(*urls, id)
				del_urls = append(del_urls, url)
				break
			}
		}
	}
	global.Log.WithField("del urls", del_urls).Info("delDefaultPic")
	return del_urls
}

type picid2urlmap struct {
	PicId  int64  `json:"pic_id"`
	PicUrl string `json:"pic_url"`
}

func GetDefaultPic(c *gin.Context) {
	o_type := c.Query("object_type")
	var picmap []picid2urlmap
	switch o_type {
	case "task":
		for k, v := range global.TaskPhotourls {
			picmap = append(picmap, picid2urlmap{
				PicId:  k,
				PicUrl: v,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"result_status": 0,
			"reslut_msg":    "get default pic",
			"picid2urlmap":  picmap,
		})
	case "work":
		for k, v := range global.WorkPhotourls {
			picmap = append(picmap, picid2urlmap{
				PicId:  k,
				PicUrl: v,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"result_status": 0,
			"reslut_msg":    "get default pic",
			"picid2urlmap":  picmap,
		})
	case "userbackground":
		for k, v := range global.UserBackgroundurls {
			picmap = append(picmap, picid2urlmap{
				PicId:  k,
				PicUrl: v,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"result_status": 0,
			"reslut_msg":    "get default pic",
			"picid2urlmap":  picmap,
		})
	case "headphoto":
		for k, v := range global.HeadPhotourls {
			picmap = append(picmap, picid2urlmap{
				PicId:  k,
				PicUrl: v,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"result_status": 0,
			"reslut_msg":    "get default pic",
			"picid2urlmap":  picmap,
		})
	default:
		c.JSON(http.StatusOK, ErrorResponse(errors.New("object type error")))
	}
}

func UpdateDefaultPic(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	f := form.File["pictures"]
	object_types := form.Value["object_type"]
	var old_pic_ids []int64
	for _, id_s := range form.Value["pic_id"] {
		id, err := strconv.ParseInt(id_s, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, ErrorResponse(err))
			return
		}
		old_pic_ids = append(old_pic_ids, id)
	}

	public_path, res := service.Uploadfile(f, c.Request.Host)
	if res != nil {
		c.JSON(http.StatusOK, res)
		return
	}
	var old_urls []string
	for _, object_type := range object_types {
		switch object_type {
		case "task":
			for i, url := range public_path {
				if _, exist := global.TaskPhotourls[old_pic_ids[i]]; !exist {
					c.JSON(http.StatusOK, ErrorResponse(errors.New("id not exist")))
				}
				old_urls = append(old_urls, global.TaskPhotourls[old_pic_ids[i]])
				delete(global.TaskPhotourls, old_pic_ids[i])
				global.TaskPhotourls[global.Worker.GetId()] = url
			}
		case "work":
			for i, url := range public_path {
				if _, exist := global.WorkPhotourls[old_pic_ids[i]]; !exist {
					c.JSON(http.StatusOK, ErrorResponse(errors.New("id not exist")))
				}
				old_urls = append(old_urls, global.WorkPhotourls[old_pic_ids[i]])
				delete(global.WorkPhotourls, old_pic_ids[i])
				global.WorkPhotourls[global.Worker.GetId()] = url
			}
		case "userbackground":
			for i, url := range public_path {
				if _, exist := global.UserBackgroundurls[old_pic_ids[i]]; !exist {
					c.JSON(http.StatusOK, ErrorResponse(errors.New("id not exist")))
				}
				old_urls = append(old_urls, global.UserBackgroundurls[old_pic_ids[i]])
				delete(global.UserBackgroundurls, old_pic_ids[i])
				global.UserBackgroundurls[global.Worker.GetId()] = url
			}
		case "headphoto":
			for i, url := range public_path {
				if _, exist := global.HeadPhotourls[old_pic_ids[i]]; !exist {
					c.JSON(http.StatusOK, ErrorResponse(errors.New("id not exist")))
				}
				old_urls = append(old_urls, global.HeadPhotourls[old_pic_ids[i]])
				delete(global.HeadPhotourls, old_pic_ids[i])
				global.HeadPhotourls[global.Worker.GetId()] = url
			}
		default:
			c.JSON(http.StatusOK, ErrorResponse(errors.New("object type error")))
			return
		}
		res, err = service.Deletefile(old_urls)
		if err != nil {
			c.JSON(http.StatusOK, res)
			return
		}
	}

	go dao.DefaultPicPersistence(object_types)
	c.JSON(http.StatusOK, gin.H{
		"reslut_status": 0,
		"reslut_msg":    "updateDefaultPic",
	})
}
