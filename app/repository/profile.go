package repository

import "omochi/app/models"

type ProfileRepository struct {}

func (ProfileRepository) Create(profile *models.Profile) error {
	if err := DB.Create(&profile).Error; err != nil {
		return err
	}
	return nil
}

func (ProfileRepository) Get(userId int64) (*models.Profile, error) {
	profile := models.Profile{}
	if err := DB.Order("id desc").Where("user_id = ?",userId).First(&profile).Error; err != nil{
		return nil, err
	}
	return &profile, nil
}

func (ProfileRepository) Update(editProfile *models.Profile) error {
	if err := DB.Model(&models.Profile{}).Updates(editProfile).Error; err != nil {
		return err
	}
	return nil
}

func (ProfileRepository) Delete(id int64) error {
	profile := models.Profile{}
	if err := DB.Order("id desc").Where("id = ?", id).First(&profile).Error; err != nil{
		return err
	}
	DB.Delete(&profile)
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
	rows, err := DB.Raw("SELECT * FROM (SELECT *, rank() over(partition by user_id order by id desc) AS rank FROM profiles) AS a WHERE rank = 1").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var arr []Profile
	for rows.Next() {

		profile := Profile{}
		DB.ScanRows(rows, &profile)
		arr = append(arr, profile)
	}
	return arr, nil
}