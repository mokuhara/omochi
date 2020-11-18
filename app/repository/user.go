package repository

import "omochi/app/models"

type UserRepository struct {}

func (UserRepository) Create(user *models.User) error {
	db := DBCon()
	defer db.Close()
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (UserRepository) GetByEmail(email string) (*models.User, error){
	db := DBCon()
	defer db.Close()
	user := models.User{}
	if err := db.Order("updated_at desc").Where("email = ?", email).First(&user).Error; err != nil{
		return nil, err
	}
	return &user, nil
}

func (UserRepository) GetByUserId(userId string) (*models.User, error){
	db := DBCon()
	defer db.Close()
	user := models.User{}
	if err := db.Order("updated_at desc").Where("user_id = ?", userId).First(&user).Error; err != nil{
		return nil, err
	}
	return &user, nil
}

func (UserRepository) Update(editUser *models.User) error {
	db := DBCon()
	defer db.Close()
	if err := db.Model(&models.User{}).Updates(editUser).Error; err != nil {
		return err
	}
	return nil
}

func (UserRepository) Delete(userId int64) error {
	db := DBCon()
	defer db.Close()
	user := models.User{}
	if err := db.Order("updated_at desc").Where("user_id = ?", userId).First(&user).Error; err != nil{
		return err
	}
	db.Delete(&user)
	return nil
}

func (UserRepository) GetAll() ([]models.User, error){
	db := DBCon()
	defer db.Close()
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

