package repository

import "omochi/app/models"

type VideoMeetingRepository struct {}

func (VideoMeetingRepository) Create(videoMeeting *models.VideoMeeting) error {
	db := DBCon()
	defer db.Close()

	if err := db.Create(&videoMeeting).Error; err != nil {
		return err
	}
	return nil
}

func (VideoMeetingRepository) Update(editVideoMeeting *models.VideoMeeting) error {
	db := DBCon()
	defer db.Close()
	videoMeeting := models.VideoMeeting{}
	if err := db.Model(&videoMeeting).Updates(editVideoMeeting).Error; err != nil {
		return err
	}
	return nil
}

func (VideoMeetingRepository) Delete (videoMeetingId int64) error {
	db := DBCon()
	defer db.Close()
	videoMeeting := models.VideoMeeting{}
	if err := db.Where("id = ?", videoMeetingId).Delete(&videoMeeting).Error; err != nil {
		return err
	}
	return nil
}

func (VideoMeetingRepository) GetByUserId(userId int64) (*[]models.VideoMeeting, error) {
	db := DBCon()
	defer db.Close()
	videoMeetings := []models.VideoMeeting{}

	if err := db.Set("gorm:auto_preload", true).Where("user_id = ?", userId).Find(&videoMeetings).Error; err != nil {
		return nil, err
	}
	return &videoMeetings, nil
}