package dao

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

func DefaultPicPersistence(object_types []string) error {
	for _, object_type := range object_types {
		defaultpic := models.DefaultPicture{
			ObjectType: object_type,
		}
		var id2url []models.DefaultPictureId2Url
		switch object_type {
		case "task":
			for k, v := range global.TaskPhotourls {
				id2url = append(id2url, models.DefaultPictureId2Url{
					Id:  k,
					Url: v,
				})
			}
		case "work":
			for k, v := range global.WorkPhotourls {
				id2url = append(id2url, models.DefaultPictureId2Url{
					Id:  k,
					Url: v,
				})
			}
		case "userbackground":
			for k, v := range global.UserBackgroundurls {
				id2url = append(id2url, models.DefaultPictureId2Url{
					Id:  k,
					Url: v,
				})
			}
		case "headphoto":
			for k, v := range global.HeadPhotourls {
				id2url = append(id2url, models.DefaultPictureId2Url{
					Id:  k,
					Url: v,
				})
			}
		}
		err := db.DB.Model(&defaultpic).Association("PicMap").Replace(id2url)
		global.Log.WithError(err).Error("DefaultPicPersistence")
	}

	db.DB.Where("object_type is null").Delete(&models.DefaultPictureId2Url{})
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
