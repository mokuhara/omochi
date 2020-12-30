package repository

import "omochi/app/models"

type ReviewRepository struct {}

func (ReviewRepository) Create(review *models.Review) error {
	db := DBCon()
	defer db.Close()

	if err := db.Create(&review).Error; err != nil {
		return err
	}
	return nil
}

func (ReviewRepository) Update(editReview *models.Review) error {
	db := DBCon()
	defer db.Close()
	review := models.Review{}
	if err := db.Model(&review).Updates(editReview).Error; err != nil {
		return err
	}
	return nil
}

func (ReviewRepository) Delete (reviewId int64) error {
	db := DBCon()
	defer db.Close()
	review := models.Review{}
	if err := db.Where("id = ?", reviewId).Delete(&review).Error; err != nil {
		return err
	}
	return nil
}

func (ReviewRepository) GetByUserId(userId int64) (*[]models.Review, error) {
	db := DBCon()
	defer db.Close()
	reviews := []models.Review{}

	if err := db.Set("gorm:auto_preload", true).Where("user_id = ?", userId).Find(&reviews).Error; err != nil {
		return nil, err
	}
	return &reviews, nil
}

func (ReviewRepository) GetByTransactionId(transactionId int64) (*[]models.Review, error) {
	db := DBCon()
	defer db.Close()
	reviews := []models.Review{}

	if err := db.Set("gorm:auto_preload", true).Where("transaction_id = ?", transactionId).Find(&reviews).Error; err != nil {
		return nil, err
	}
	return &reviews, nil
}