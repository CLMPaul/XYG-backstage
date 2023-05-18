package global

import (
	"time"
	"xueyigou_demo/models"
	"xueyigou_demo/tools"
	"xueyigou_demo/tools/illegal_word/Words"

	"github.com/sirupsen/logrus"
)

const (
	PublicPath = "/static/"
	LocalPath  = "./public/pictures/"
)

var Host = "localhost:8080"

var Log = logrus.New()

var Worker, _ = tools.NewWorker(1) //每台机器部署代码不一致

var Timelayout = "2006-01-02 15:04"

var Expires time.Time

var RTokenExpires time.Time

// 默认的任务，头像，作品图片
var TaskPhotourl = "https://i.postimg.cc/8fNVYYmz/image.png"

var WorkPhotourl = "https://i.postimg.cc/rzk0cFhK/image.png"

var UserBackgroundurl = "https://i.postimg.cc/jSbdsm8m/image.png"

var HeadPhotosMap map[string]string

var TaskPhotourls map[int64]string

var WorkPhotourls map[int64]string

var UserBackgroundurls map[int64]string

var HeadPhotourls map[int64]string

var LikeChan chan models.LikeForm

var IllegalWords Words.StringSearchEx

var SlatSuperUserPassword = "xueyigou_superuser"

var AWSAccessKey = "AKIAV262STBDOEZF7CMF"

var AWSSecretKey = "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"
