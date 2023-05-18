package service

import (
	"xueyigou_demo/db"
	"xueyigou_demo/global"
	"xueyigou_demo/models"
)

func InitDefaultPic() error {
	err := db.DB.AutoMigrate(&models.DefaultPicture{}, &models.DefaultPictureId2Url{})
	if err != nil {
		return err
	}

	defaultpic := models.DefaultPicture{
		ObjectType: "task",
	}

	err = db.DB.Create(&defaultpic).Error
	if err != nil {
		return err
	}

	defaultpic.ObjectType = "work"
	err = db.DB.Create(&defaultpic).Error
	if err != nil {
		return err
	}

	defaultpic.ObjectType = "userbackground"
	err = db.DB.Create(&defaultpic).Error
	if err != nil {
		return err
	}

	defaultpic.ObjectType = "headphoto"
	err = db.DB.Create(&defaultpic).Error
	if err != nil {
		return err
	}
	return nil
}

func DefaultPicRecover() error {
	var default_pics []models.DefaultPicture
	err := db.DB.Preload("PicMap").Find(&default_pics).Error
	if err != nil {
		return err
	}
	for _, default_pic := range default_pics {
		switch default_pic.ObjectType {
		case "task":
			for _, pic_map := range default_pic.PicMap {
				global.TaskPhotourls[pic_map.Id] = pic_map.Url
			}
		case "work":
			for _, pic_map := range default_pic.PicMap {
				global.WorkPhotourls[pic_map.Id] = pic_map.Url
			}
		case "userbackground":
			for _, pic_map := range default_pic.PicMap {
				global.UserBackgroundurls[pic_map.Id] = pic_map.Url
			}
		case "headphoto":
			for _, pic_map := range default_pic.PicMap {
				global.HeadPhotourls[pic_map.Id] = pic_map.Url
			}

		}
	}
	return nil
}
