package repository

import (
	"omochi/app/models"
)

type ClientRepository struct {}

func (ClientRepository) Create(user *models.User) error {
	client := models.Client{}

	if err := DB.Create(&user).Error; err != nil {
		return err
	}
	client.UserID = int64(user.ID)

	company := models.Company{Name: ""}
	if err := DB.Create(&company).Error; err != nil {
		return err
	}
	client.CompanyID = int64(company.ID)
	if err := DB.Create(&client).Error; err != nil {
		return err
	}
	return nil
}

func (ClientRepository) GetByEmail(email string) (*models.Client, error){
	client := models.Client{}
	if err := DB.Where("email = ?", email).First(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (ClientRepository) GetByUserId(userId int64) (*models.Client, error) {
	client := models.Client{}
	if err := DB.Where("user_id = ?", userId).First(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (ClientRepository) Update(editClient *models.Client) error {
	client := models.Client{}
	if err := DB.Model(&client).Updates(editClient).Error; err != nil {
		return err
	}
	return nil
}

func (ClientRepository) Delete(userId int64) error {
	client := models.Client{}
	if err:= DB.Where("user_id = ?", userId).Delete(&client).Error; err != nil {
		return err
	}
	return nil
}

func (ClientRepository) GetAll() ([]models.Client, error) {
	var clients []models.Client
	if err := DB.Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}



