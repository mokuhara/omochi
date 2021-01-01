package repository

import (
	"omochi/app/models"
)

type SpecialistRepository struct {}

func (SpecialistRepository) Create(user *models.User) error {
	specialist := models.Specialist{}
	if err := DB.Create(&user).Error; err != nil {
		return err
	}
	specialist.UserID = int64(user.ID)

	if err := DB.Create(&specialist).Error; err != nil {
		return err
	}
	return nil
}

func (SpecialistRepository) GetByEmail(email string) (*models.Specialist, error){
	specialist := models.Specialist{}
	if err := DB.Where("email = ?", email).First(&specialist).Error; err != nil {
		return nil, err
	}
	return &specialist, nil
}

func (SpecialistRepository) GetByUserId(userId int64) (*models.Specialist, error) {
	specialist := models.Specialist{}
	if err := DB.Where("user_id = ?", userId).First(&specialist).Error; err != nil {
		return nil, err
	}
	return &specialist, nil
}

func (SpecialistRepository) Update(editSpecialist *models.Specialist) error {
	specialist := models.Specialist{}
	if err := DB.Model(&specialist).Updates(editSpecialist).Error; err != nil {
		return err
	}
	return nil
}

func (SpecialistRepository) Delete(userId int64) error {
	specialist := models.Specialist{}
	if err:= DB.Where("user_id = ?", userId).Delete(&specialist).Error; err != nil {
		return err
	}
	return nil
}

func (SpecialistRepository) GetAll() ([]models.Specialist, error) {
	var specialists []models.Specialist
	if err := DB.Find(&specialists).Error; err != nil {
		return nil, err
	}
	return specialists, nil
}