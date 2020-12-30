package repository

import "omochi/app/models"

type IssueRepository struct {}

func (IssueRepository) Create(issue *models.Issue) error {
	db := DBCon()
	defer db.Close()

	if err := db.Create(&issue).Error; err != nil {
		return err
	}
	return nil
}

func (IssueRepository) Update(editIssue *models.Issue) error {
	//部分更新に対応させてないので追々実装する
	db := DBCon()
	defer db.Close()
	issue := models.Issue{}
	if err := db.Model(&issue).Updates(editIssue).Error; err != nil {
		return err
	}
	return nil
}

func (IssueRepository) Delete(issueId int64) error {
	db := DBCon()
	defer db.Close()
	issue := models.Issue{}
	if err := db.Where("id = ?", issueId).Delete(&issue).Error; err != nil {
		return err
	}
	return nil
}

func (IssueRepository) GetAll() (*[]models.Issue, error) {
	db := DBCon()
	defer db.Close()
	var issues []models.Issue
	if err := db.Set("gorm:auto_preload", true).Find(&issues).Error; err != nil {
		return nil, err
	}
	return &issues, nil
}

func (IssueRepository) Find(issueId int64) (*models.Issue, error){
	db := DBCon()
	defer db.Close()
	issue := models.Issue{}

	if err := db.Set("gorm:auto_preload", true).Where("id = ?", issueId).First(&issue).Error; err != nil {
		return nil, err
	}
	return &issue, nil
}

func (IssueRepository) GetByUserId(userId int64) (*[]models.Issue, error){
	db := DBCon()
	defer db.Close()
	issues:= []models.Issue{}

	if err := db.Set("gorm:auto_preload", true).Where("user_id = ?", userId).Find(&issues).Error; err != nil {
		return nil, err
	}
	return &issues, nil
}

