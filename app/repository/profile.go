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

func (ProfileRepository) Delete(id int64) error {
	db := DBCon()
	defer db.Close()
	profile := models.Profile{}
	if err := db.Order("id desc").Where("id = ?", id).First(&profile).Error; err != nil{
		return err
	}
	db.Delete(&profile)
	return nil
}

type Profile struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"userId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

func (ProfileRepository) GetAll() ([]Profile, error){
	db := DBCon()
	defer DBCon()
	rows, err := db.Raw("SELECT * FROM (SELECT *, rank() over(partition by user_id order by id desc) AS rank FROM profiles) AS a WHERE rank = 1").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var arr []Profile
	for rows.Next() {

		profile := Profile{}
		db.ScanRows(rows, &profile)
		arr = append(arr, profile)
	}
	return arr, nil
}