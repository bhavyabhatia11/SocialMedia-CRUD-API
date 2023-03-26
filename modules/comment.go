package modules

import (
	"github.com/bhavyaunacademy/goCRUD-MySQL/models"
	"gorm.io/gorm"
)

func CreateComment(db *gorm.DB, Comment *models.Comment) (err error) {

	err = db.Create(Comment).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCommentsOfPost(db *gorm.DB, Comments *[]models.Comment, post_uid int) (err error) {

	err = db.Where("post_uid = ?", post_uid).Find(Comments).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateComment(db *gorm.DB, UpdatedComment *models.Comment, comment_uid int) (err error) {

	var Comment models.Comment
	err = db.Where("uid = ? AND user_uid = ?", comment_uid, UpdatedComment.UserUID).First(&Comment).Error
	if err != nil {
		return err
	}

	Comment.Content = UpdatedComment.Content

	err = db.Save(Comment).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment(db *gorm.DB, Comment *models.Comment, comment_uid int) (err error) {
	err = db.Where("uid = ? AND user_uid = ?", comment_uid, Comment.UserUID).Delete(Comment).Error
	if err != nil {
		return err
	}
	return nil
}
