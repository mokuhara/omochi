package repository

import "omochi/app/models"

type UserInfoRepository struct{}

func (UserInfoRepository) Create(userInfo *models.UserInfo) error {
	if err := DB.Create(userInfo).Error; err != nil {
		return err
	}
	return nil
}

func (UserInfoRepository) GetByUserId(userId int64) (*models.UserInfo, error) {
	userInfo := models.UserInfo{}
	if err := DB.Order("updated_at desc").Where("user_id", userId).First(&userInfo).Error; err != nil {
		return nil, err
	}
	return &userInfo, nil
}