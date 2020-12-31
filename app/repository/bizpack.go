package repository

import "omochi/app/models"
import "log"

type BizpackRepository struct {}

func (BizpackRepository) Create(bizpack *models.Bizpack) error {
	db := DBCon()
	defer db.Close()

	if err := db.Create(&bizpack).Error; err != nil {
		return err
	}
	return nil
}

func (BizpackRepository) Update(editBizpack *models.Bizpack) error {
	//部分更新に対応させてないので追々実装する
	db := DBCon()
	defer db.Close()
	bizpack := models.Bizpack{}
	if err := db.Model(&bizpack).Updates(editBizpack).Error; err != nil {
		return err
	}
	return nil
}

func (BizpackRepository) Delete(bizpackId int64) error {
	db := DBCon()
	defer db.Close()
	bizpack := models.Bizpack{}
	if err := db.Where("id = ?", bizpackId).Delete(&bizpack).Error; err != nil {
		return err
	}
	return nil
}

func (BizpackRepository) GetAll() (*[]models.Bizpack, error) {
	db := DBCon()
	defer db.Close()

	var bizpacks []models.Bizpack

	if err := db.Set("gorm:auto_preload", true).Find(&bizpacks).Error; err != nil {
		return nil, err
	}

	return &bizpacks, nil
}

func (BizpackRepository) Find(bizpackId int64) (*models.Bizpack, error){
	db := DBCon()
	defer db.Close()

	bizpack := models.Bizpack{}

	if err := db.Set("gorm:auto_preload", true).Where("id = ?", bizpackId).First(&bizpack).Error; err != nil {
		return nil, err
	}
	return &bizpack, nil
}

func (BizpackRepository) GetByUserId(userId int64) (*[]models.Bizpack, error){
	db := DBCon()
	defer db.Close()

	bizpacks:= []models.Bizpack{}

	log.Println(userId)
	if err := db.Set("gorm:auto_preload", true).Where("user_id = ?", userId).Find(&bizpacks).Error; err != nil {
		return nil, err
	}

	return &bizpacks, nil
}

func (BizpackRepository) GetByBizpackId(bizpackId int64) (*models.Bizpack, error) {
	db := DBCon()
	defer db.Close()
	bizpack := models.Bizpack{}

	if err := db.Set("gorm:auto_preload", true).Where("id = ?", bizpackId).First(&bizpack).Error; err != nil {
		return nil, err
	}
	return &bizpack, nil
}