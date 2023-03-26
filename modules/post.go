package modules

import (
	"github.com/bhavyaunacademy/goCRUD-MySQL/models"
	"gorm.io/gorm"
)

func GetPosts(db *gorm.DB, Posts *[]models.Post) (err error) {
	err = db.Find(Posts).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPost(db *gorm.DB, Post *models.Post, uid int) (err error) {
	err = db.Where("uid = ?", uid).First(Post).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPostsByUser(db *gorm.DB, Posts *[]models.Post, user_uid int) (err error) {
	err = db.Where("user_uid = ?", user_uid).Find(Posts).Error
	if err != nil {
		return err
	}
	return nil
}

func CreatePost(db *gorm.DB, Post *models.Post) (err error) {

	err = db.Create(Post).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdatePost(db *gorm.DB, UpdatedPost *models.Post, post_uid int) (err error) {
	var Post *models.Post
	err = db.Where("uid = ? AND user_uid = ?", post_uid, UpdatedPost.UserUID).First(&Post).Error
	if err != nil {
		return err
	}

	Post.Title = UpdatedPost.Title
	Post.Content = UpdatedPost.Content

	err = db.Save(Post).Error
	if err != nil {
		return err
	}
	return nil
}

func DeletePost(db *gorm.DB, Post *models.Post, post_uid int) (err error) {
	err = db.Where("uid = ? AND user_uid = ?", post_uid, Post.UserUID).Delete(Post).Error
	if err != nil {
		return err
	}
	return nil
}

func LikePost(db *gorm.DB, post_uid int) (err error) {
	var Post *models.Post
	err = db.Where("uid = ?", post_uid).First(&Post).Error
	if err != nil {
		return err
	}

	Post.Likes += 1

	err = db.Save(&Post).Error
	if err != nil {
		return err
	}

	return nil

}

func DislikePost(db *gorm.DB, post_uid int) (err error) {
	var Post *models.Post
	err = db.Where("uid = ?", post_uid).First(&Post).Error
	if err != nil {
		return err
	}

	Post.Dislikes += 1

	err = db.Save(&Post).Error
	if err != nil {
		return err
	}

	return nil

}
