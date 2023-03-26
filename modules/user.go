package modules

import (
	"github.com/bhavyaunacademy/goCRUD-MySQL/models"
	"gorm.io/gorm"
)

//create a user
func CreateUser(db *gorm.DB, User *models.User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

//get users
func GetUsers(db *gorm.DB, Users *[]models.User) (err error) {
	err = db.Find(Users).Error
	if err != nil {
		return err
	}
	return nil
}

//get user by id
func GetUser(db *gorm.DB, User *models.User, uid int) (err error) {
	err = db.Where("uid = ?", uid).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

//update user
func UpdateUser(db *gorm.DB, User *models.User) (err error) {
	db.Save(User)
	return nil
}

//delete user
func DeleteUser(db *gorm.DB, User *models.User, id int) (err error) {
	db.Where("id = ?", id).Delete(User)
	return nil
}
