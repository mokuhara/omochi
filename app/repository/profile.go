package repository

import "omochi/app/models"

type ProfileRepository struct {}

func (ProfileRepository) Create(profile *models.Profile) error {
	db := DBCon()
	defer db.Close()
	if err := db.Create(&profile).Error; err != nil {
		return err
	}
	return nil
}

func (ProfileRepository) Get(userId int64) (*models.Profile, error) {
	db := DBCon()
	defer db.Close()
	profile := models.Profile{}
	if err := db.Order("id desc").Where("user_id = ?",userId).First(&profile).Error; err != nil{
		return nil, err
	}
	return &profile, nil
}

func (ProfileRepository) Update(editProfile *models.Profile) error {
	db := DBCon()
	defer db.Close()
	if err := db.Model(&models.Profile{}).Updates(editProfile).Error; err != nil {
		return err
	}
	return nil
}

func (ProfileRepository) Delete(userId int64) error {
	db := DBCon()
	defer db.Close()
	profile := models.Profile{}
	if err := db.Order("id desc").Where("user_id = ?", userId).First(&profile).Error; err != nil{
		return err
	}
	db.Delete(&profile)
	return nil
}

func (ProfileRepository) GetAll() ([]models.Profile, error){
	db := DBCon()
	defer DBCon()
	rows, err := db.Raw("SELECT * FROM (SELECT *, rank() over(partition by user_id order by id desc) AS rank FROM profile LIMIT 100 ) AS a WHERE rank = 1").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var arr []models.Profile
	for rows.Next() {
		profile := models.Profile{}
		rows.Scan(&profile)
		arr = append(arr, profile)
	}
	return arr, nil
}