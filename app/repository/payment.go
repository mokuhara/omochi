package repository

import "omochi/app/models"

type PaymentRepository struct {}

func (PaymentRepository) Create(payment *models.Payment) error {
	db := DBCon()
	defer db.Close()

	if err := db.Create(&payment).Error; err != nil {
		return err
	}
	return nil
}

func (PaymentRepository) Update(editPayment *models.Payment) error {
	db := DBCon()
	defer db.Close()
	payment := models.Payment{}
	if err := db.Model(&payment).Updates(editPayment).Error; err != nil {
		return err
	}
	return nil
}

func (PaymentRepository) Delete (paymentId int64) error {
	db := DBCon()
	defer db.Close()
	payment := models.Payment{}
	if err := db.Where("id = ?", paymentId).Delete(&payment).Error; err != nil {
		return err
	}
	return nil
}

func (PaymentRepository) GetByUserId(userId int64) (*[]models.Payment, error) {
	db := DBCon()
	defer db.Close()
	payments := []models.Payment{}

	if err := db.Set("gorm:auto_preload", true).Where("user_id = ?", userId).Find(&payments).Error; err != nil {
		return nil, err
	}
	return &payments, nil
}