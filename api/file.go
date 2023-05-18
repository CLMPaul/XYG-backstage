package api

import (
	"fmt"
	"net/http"
	"xueyigou_demo/serializer"
	"xueyigou_demo/service"

	"github.com/gin-gonic/gin"
)

func UploadPicture(c *gin.Context) {
	// file, err := c.FormFile("pictures")
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	f := form.File["pictures"]
	public_path, res := service.Uploadfile(f, c.Request.Host)
	if res != nil {
		c.JSON(http.StatusOK, res)
	}
	// var filepaths []string
	// for _, file := range f {
	// 	fileExt := strings.ToLower(path.Ext(file.Filename))
	// 	filename := path.Base(file.Filename)
	// 	//TODO:make map<ext,allowed>
	// 	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
	// 		global.Log.WithFields(logrus.Fields{
	// 			"file name": filename,
	// 			"file ext":  fileExt,
	// 		}).Error()
	// 		c.JSON(200, gin.H{
	// 			"result_status": 1,
	// 			"result_msg":    "上传失败!只允许png,jpg,gif,jpeg文件",
	// 		})
	// 		return
	// 	}
	// 	filename = middleware.Md5Crypt(filename+time.Now().String(), "xueyigou")
	// 	filepath := path.Join("./public/pictures", filename+fileExt)
	// 	if tools.FileExist(filepath) {
	// 		continue
	// 	}
	// 	filepaths = append(filepaths, filepath)
	// 	err = c.SaveUploadedFile(file, filepath)
	// 	if err != nil {
	// 		c.JSON(http.StatusOK, err.Error())
	// 		return
	// 	}
	// }
	// public_path := tools.LocalPath2Public(filepaths, c.Request.Host)

	// if err != nil {
	// 	c.JSON(http.StatusOK, ErrorResponse(err))
	// }

	c.JSON(http.StatusOK, serializer.BuildAddWorkPicturesUrlResponse(public_path))
	//c.JSON(http.StatusOK, gin.H{"uploading": "done", "message": "success", "url": "http://" + c.Request.Host + "/public/pictures/" + file.Filename})
}

type Deletefileform struct {
	Urls []string
}

func DeletePicture(c *gin.Context) {
	var form Deletefileform
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	fmt.Println(form)
	urls := form.Urls
	response, _ := service.Deletefile(urls)
	c.JSON(http.StatusOK, response)
}
