package dao

import (
	"xueyigou_demo/db"
	"xueyigou_demo/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//	Add

func AddCommentL1(comment *models.FirstLevelComment) error {
	if err := db.DB.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

// 添加二级评论
func AddCommentL2(comment *models.SecondLevelComment) error {
	if err := db.DB.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

func FindL1CommentById(commentId int64) models.FirstLevelComment {
	var comment models.FirstLevelComment
	if err := db.DB.Where("id = ?", commentId).Preload("Commenter").First(&comment).Error; err != nil {
		panic(err)
	}
	return comment
}
func FindL2CommentById(commentId int64) models.SecondLevelComment {
	var comment models.SecondLevelComment
	if err := db.DB.Where("id = ?", commentId).First(&comment).Error; err != nil {
		panic(err)
	}
	return comment
}

//	Delete

func DeleteCommentById(commentId int64, object_type models.ObjectType, comment_level bool) error {
	//mu.Lock()
	//defer mu.Unlock()
	// 开启事务
	return db.DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		var err error
		//判断comment_id位于何表
		if comment_level {
			//在一级评论表
			if err = tx.Where("id = ? AND object_type = ?", commentId, object_type).Delete(&models.FirstLevelComment{}).Error; err == nil {
				if err1 := tx.Where("first_level_comment_id = ? AND object_type = ?",
					commentId, object_type).Delete(&models.SecondLevelComment{}).Error; err != nil {
					return err1
				}
				return nil
			}
		} else {
			//在二级评论表
			if err = tx.Where("id = ? AND object_type = ?", commentId, object_type).Delete(&models.SecondLevelComment{}).Error; err == nil {
				return nil
			}
		}
		return err
	})
}

func GetFirstLevelCommand(good_id int64, object_type models.ObjectType) []models.FirstLevelComment {
	var comments []models.FirstLevelComment
	// if err := db.DB.Raw("select * from first_level_comments where good_id=? order by post_date desc", good_id).Scan(&comment).Error; err != nil {
	// 	panic(err)
	// }
	if err := db.DB.Where("good_id = ? AND object_type = ?", good_id,
		object_type).Preload("Commenter").Order("created_at desc").Find(&comments).Error; err != nil {
		panic(err)
	}

	//get likes
	for index, comment := range comments {
		comments[index].CommentLikes = Likes_get(comment.ID, 0)
	}
	return comments
}

func GetSecondLevelCommand(first_level_comment_id int64, object_type models.ObjectType) []models.SecondLevelComment {
	var comments []models.SecondLevelComment
	// if err := db.DB.Raw("select * from second_level_comments where first_level_comment_id=? order by reply_date desc", first_level_comment_id).Scan(&comment).Error; err != nil {
	// 	panic(err)
	// }
	if err := db.DB.Where("first_level_comment_id = ? AND object_type = ?", first_level_comment_id,
		object_type).Preload(clause.Associations).Order("created_at").Find(&comments).Error; err != nil {
		panic(err)
	}

	//get likes
	for index, comment := range comments {
		comments[index].CommentLikes = Likes_get(comment.ID, 1)
	}
	return comments

}

func FindL1Commenter(flcid int64) models.User {
	var comment models.FirstLevelComment
	if err := db.DB.Select("id, commenter_id").Preload("Commenter").Where("id = ?", flcid).First(&comment).Error; err != nil {
		panic(err)
	}
	return comment.Commenter
}
