package service

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
	"xueyigou_demo/global"
	"xueyigou_demo/middleware"
	"xueyigou_demo/serializer"
	"xueyigou_demo/tools"

	"github.com/sirupsen/logrus"
)

func Uploadfile(f []*multipart.FileHeader, host string) ([]string, interface{}) {
	var filepaths []string
	for _, file := range f {
		fileExt := strings.ToLower(path.Ext(file.Filename))
		filename := path.Base(file.Filename)
		//TODO:make map<ext,allowed>
		if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
			global.Log.WithFields(logrus.Fields{
				"file name": filename,
				"file ext":  fileExt,
			}).Error()
			return nil, serializer.BuildFailResponse("上传失败!只允许png,jpg,gif,jpeg文件")
		}
		filename = middleware.Md5Crypt(filename+time.Now().String(), "xueyigou")
		filepath := path.Join("./public/pictures", filename+fileExt)
		if tools.FileExist(filepath) {
			continue
		}
		filepaths = append(filepaths, filepath)
		err := SaveUploadedFile(file, filepath)
		if err != nil {
			return nil, serializer.BuildFailResponse(err.Error())
		}
	}
	public_path := tools.LocalPath2Public(filepaths, host)
	return public_path, nil
}

func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// @param 访问路径 string
func Deletefile(paths []string) (serializer.Response, error) {
	global.Log.WithField("url", paths).Info("del file")
	local_paths := tools.PublicPath2Local(paths)
	global.Log.WithField("paths", local_paths).Info("del file")
	for _, path_ := range local_paths {

		if err := os.Remove(path_); err != nil {
			fmt.Println(path_, err)
			return serializer.BuildFailResponse("delete file"), err
		}
	}
	return serializer.BuildSuccessResponse("delete file"), nil
}

func GetRemoteFile(url string, host string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	filename := middleware.Md5Crypt(url+time.Now().String(), "xueyigou")
	filepath := path.Join("./public/pictures", filename+".png")
	out, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}
	public_path := tools.LocalPath2Public([]string{filepath}, host)
	if len(public_path) != 1 {
		return "", errors.New("path lenngth err")
	}
	return public_path[0], err
}
