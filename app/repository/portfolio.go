package repository

import "omochi/app/models"

type PortfolioRepository struct {}

func (PortfolioRepository) Create(portfolio *models.Portfolio) error {
	db := DBCon()
	defer db.Close()

	if err := db.Create(&portfolio).Error; err != nil {
		return err
	}
	return nil
}

func (PortfolioRepository) Update(editPortfolio *models.Portfolio) error {
	//部分更新に対応させてないので追々実装する
	db := DBCon()
	defer db.Close()
	portfolio := models.Portfolio{}
	if err := db.Model(&portfolio).Updates(editPortfolio).Error; err != nil {
		return err
	}
	return nil
}

func (PortfolioRepository) Delete(portfolioId int64) error {
	db := DBCon()
	defer db.Close()
	portfolio := models.Portfolio{}
	if err := db.Where("id = ?", portfolioId).Delete(&portfolio).Error; err != nil {
		return err
	}
	return nil
}

func (PortfolioRepository) GetAll() (*[]models.Portfolio, error) {
	db := DBCon()
	defer db.Close()
	var portfolios []models.Portfolio
	if err := db.Set("gorm:auto_preload", true).Find(&portfolios).Error; err != nil {
		return nil, err
	}
	return &portfolios, nil
}

func (PortfolioRepository) Find(portfolioId int64) (*models.Portfolio, error){
	db := DBCon()
	defer db.Close()
	portfolio := models.Portfolio{}

	if err := db.Set("gorm:auto_preload", true).Where("id = ?", portfolioId).First(&portfolio).Error; err != nil {
		return nil, err
	}
	return &portfolio, nil
}

func (PortfolioRepository) GetByUserId(userId int64) (*[]models.Portfolio, error){
	db := DBCon()
	defer db.Close()
	portfolios:= []models.Portfolio{}

	if err := db.Set("gorm:auto_preload", true).Where("user_id = ?", userId).Find(&portfolios).Error; err != nil {
		return nil, err
	}
	return &portfolios, nil
}

func (PortfolioRepository) GetByPortfolioId(portfolioId int64) (*models.Portfolio, error) {
	db := DBCon()
	defer db.Close()
	portfolio := models.Portfolio{}

	if err := db.Set("gorm:auto_preload", true).Where("id = ?", portfolioId).First(&portfolio).Error; err != nil {
		return nil, err
	}
	return &portfolio, nil
}