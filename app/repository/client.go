package repository

import (
	"omochi/app/models"
)

type ClientRepository struct {}

func (ClientRepository) Create(user *models.User) error {
	db := DBCon()
	defer db.Close()
	client := models.Client{}

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	client.UserID = int64(user.ID)

	company := models.Company{Name: ""}
	if err := db.Create(&company).Error; err != nil {
		return err
	}
	client.CompanyID = int64(company.ID)
	if err := db.Create(&client).Error; err != nil {
		return err
	}
	return nil
}

func (ClientRepository) GetByEmail(email string) (*models.Client, error){
	db := DBCon()
	defer db.Close()
	client := models.Client{}
	if err := db.Where("email = ?", email).First(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (ClientRepository) GetByUserId(userId int64) (*models.Client, error) {
	db := DBCon()
	defer db.Close()
	client := models.Client{}
	if err := db.Where("user_id = ?", userId).First(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (ClientRepository) Update(editClient *models.Client) error {
	db := DBCon()
	defer db.Close()
	client := models.Client{}
	if err := db.Model(&client).Updates(editClient).Error; err != nil {
		return err
	}
	return nil
}

func (ClientRepository) Delete(userId int64) error {
	db := DBCon()
	defer db.Close()
	client := models.Client{}
	if err:= db.Where("user_id = ?", userId).Delete(&client).Error; err != nil {
		return err
	}
	return nil
}

func (ClientRepository) GetAll() ([]models.Client, error) {
	db := DBCon()
	defer db.Close()
	var clients []models.Client
	if err := db.Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}



