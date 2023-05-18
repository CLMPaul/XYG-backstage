package api

import (
	"fmt"
	"net/http"
	"os"
	"xueyigou_demo/global"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

type uploadForm struct {
	Bucket   string `json:"bucket"`
	FileName string `json:"filename"`
}

func UploadFile(c *gin.Context) {
	var form uploadForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	access_key := global.AWSAccessKey
	secret_key := global.AWSSecretKey
	// end_point := "http://10.0.6.247:7480" //endpoint设置，不要动

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(access_key, secret_key, ""),
		// Endpoint:         aws.String(end_point),
		Region:           aws.String("cn-north-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(false), //virtual-host style方式，不要修改
	})

	// bucket := form.Bucket
	// filename := form.FileName

	file, err := os.Open("public/pictures/d5357fd60283c553b0d20c4a957e77f7.jpg")
	if err != nil {
		global.Log.WithError(err).Info("upload")
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	defer file.Close()

	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("image.xueyigou.cn"),
		Key:    aws.String("d5357fd60283c553b0d20c4a957e77f7.jpg"),
		Body:   file,
	})

	if err != nil {
		fmt.Errorf("failed to upload file, %v", err)
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	fmt.Printf("file uploaded to, %s\n", result.Location)
}
