package repository

import "omochi/app/models"
import "log"

type BizpackRepository struct {}

func (BizpackRepository) Create(bizpack *models.Bizpack) error {
	if err := DB.Create(&bizpack).Error; err != nil {
		return err
	}

	return nil
}

//部分更新に対応させてないので追々実装する
func (BizpackRepository) Update(editBizpack *models.Bizpack) error {
	bizpack := models.Bizpack{}

	if err := DB.Model(&bizpack).Updates(editBizpack).Error; err != nil {
		return err
	}

	return nil
}

// TODO: deleteと合わせてリファクタ
func (BizpackRepository) CheckUserBizpack(userId int64, bizpackId int64) bool {
	bizpack := models.Bizpack{}

	if err := DB.Where("id = ?", bizpackId).First(&bizpack).Error; err != nil {
		return false
	}

	return bizpack.UserID == userId
}

func (BizpackRepository) Delete(bizpackId int64) error {
	bizpack := models.Bizpack{}

	if err := DB.Where("id = ?", bizpackId).Delete(&bizpack).Error; err != nil {
		return err
	}

	return nil
}

func (BizpackRepository) GetAll() (*[]models.Bizpack, error) {
	var bizpacks []models.Bizpack

	if err := DB.Set("gorm:auto_preload", true).Find(&bizpacks).Error; err != nil {
		return nil, err
	}

	return &bizpacks, nil
}

func (BizpackRepository) Find(bizpackId int64) (*models.Bizpack, error){
	bizpack := models.Bizpack{}

	if err := DB.Set("gorm:auto_preload", true).Where("id = ?", bizpackId).First(&bizpack).Error; err != nil {
		return nil, err
	}
	return &bizpack, nil
}

func (BizpackRepository) GetByUserId(userId int64) (*[]models.Bizpack, error){
	bizpacks:= []models.Bizpack{}

	log.Println(userId)
	if err := DB.Set("gorm:auto_preload", true).Where("user_id = ?", userId).Find(&bizpacks).Error; err != nil {
		return nil, err
	}

	return &bizpacks, nil
}

func (BizpackRepository) GetByBizpackId(bizpackId int64) (*models.Bizpack, error) {
	bizpack := models.Bizpack{}

	if err := DB.Set("gorm:auto_preload", true).Where("id = ?", bizpackId).First(&bizpack).Error; err != nil {
		return nil, err
	}
	return &bizpack, nil
}

// TODO: 要リファクタ
func (BizpackRepository) GetByUserIDAndBizpackId(userId int64, bizpackId int64) (*models.Bizpack, error) {
	bizpack := models.Bizpack{}

	if err := DB.Set("gorm:auto_preload", true).Where("id = ? AND user_id = ?", bizpackId, userId).First(&bizpack).Error; err != nil {
		return nil, err
	}

	return &bizpack, nil
}
