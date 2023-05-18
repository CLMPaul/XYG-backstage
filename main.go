package main

import (
	"github.com/robfig/cron"
	"io"
	"os"
	_ "xueyigou_demo/config"
	"xueyigou_demo/dao"
	"xueyigou_demo/db"
	"xueyigou_demo/global"
	"xueyigou_demo/internal/utils"
	"xueyigou_demo/models"
	"xueyigou_demo/service"
	"xueyigou_demo/web"
)

func Init() {
	global.TaskPhotourls = make(map[int64]string)
	global.WorkPhotourls = make(map[int64]string)
	global.HeadPhotourls = make(map[int64]string)
	global.UserBackgroundurls = make(map[int64]string)
	//global.LikeChan = make(chan models.LikeForm, 100)
	global.IllegalWords.Load("config/sensitive.dat")
	//for i := 0; i < 30; i++ {
	//	go dao.LikeInLocal(global.LikeChan)
	//}
	service.InitDefaultPic()
	if err := service.DefaultPicRecover(); err != nil {
		panic(err)
	}
	//TODO initialize fs
	if _, err := os.Stat(global.LocalPath); os.IsNotExist(err) {
		os.MkdirAll(global.LocalPath, os.ModePerm)
	}
	writer1 := os.Stdout
	writer2, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		global.Log.Fatalf("create file log.txt failed: %v", err)
	}
	global.Log.SetOutput(io.MultiWriter(writer1, writer2))
}

func cornfunc(c *cron.Cron) {
	c.AddFunc("0 */30 * * * ?", dao.CreateHonorList)
	c.Start()
}
func main() {
	db.SetupFromEnv()
	models.SetupDatabase(db.NewMigrator(db.DB))
	Init()
	c := cron.New()
	cornfunc(c)

	a := utils.NewApp()
	a.Run(web.RunServer)
	a.Wait()
}
