package repository

import "omochi/app/models"

type TransactionRepository struct {}

func (TransactionRepository) Create(transaction *models.Transaction) error {
	db := DBCon()
	defer db.Close()

	if err := db.Create(&transaction).Error; err != nil {
		return err
	}
	return nil
}

func (TransactionRepository) Update(editTransaction *models.Transaction) error {
	db := DBCon()
	defer db.Close()
	transaction := models.Transaction{}
	if err := db.Model(&transaction).Updates(editTransaction).Error; err != nil {
		return err
	}
	return nil
}

func (TransactionRepository) Delete (transactionId int64) error {
	db := DBCon()
	defer db.Close()
	transaction := models.Transaction{}
	if err := db.Where("id = ?", transactionId).Delete(&transaction).Error; err != nil {
		return err
	}
	return nil
}

func (TransactionRepository) GetByUserId(userId int64) (*[]models.Transaction, error) {
	db := DBCon()
	defer db.Close()
	transactions := []models.Transaction{}
	if err := db.Set("gorm:auto_preload", true).Where("specialist_user_id = ?", userId).Or("client_user_id = ?", userId).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return &transactions, nil
}