package transaction

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"omochi/app/models"
	"omochi/app/repository"
	"omochi/app/service"
	"omochi/config"
	"strconv"
)

func CreateVideoMeeting(c *gin.Context){
	type setting struct {
		Use_pmi string `json:"use_pmi"`
	}

	type zoomCreateRoom struct {
		Topic string  `json:"topic"`
		Type string `json:"type"`
		StartTime string `json:"start_time"`
		Timezone string `json:"timezone"`
		Settings setting `json:"settings"`
		TransactionId int64 `json:"transactionId`
	}

	zoomCreateRoomStruct := zoomCreateRoom{}
	err := c.BindJSON(&zoomCreateRoomStruct)
	if err != nil {
		log.Println("action=CreateVideoMeeting bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	endpoint := config.Config.ZoomEndpoint
	jwt := config.Config.ZoomJwt
	zoomCreateRoomJson, err := json.Marshal(zoomCreateRoomStruct)
	if err != nil {
		log.Println("action=CreateVideoMeeting zoomCreateRoom Marshal error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	req, _ := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(zoomCreateRoomJson))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + jwt)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("action=CreateVideoMeeting failed to create zoom room")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"data": "failed to create zoom room",
		})
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	type videoMeetingUrl struct {
		Url string `json:"join_url"`
	}
	Url := videoMeetingUrl{}
	err = json.Unmarshal(body, &Url)
	if err != nil {
		log.Println("action=CreateVideoMeeting failed to unmarshal videMeeting")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}

	videoMeeting := models.VideoMeeting{
		Name:zoomCreateRoomStruct.Topic,
		Url:Url.Url,
		StartedAt: zoomCreateRoomStruct.StartTime,
		TransactionID: zoomCreateRoomStruct.TransactionId,
	}

	videoMeetingRepository := repository.VideoMeetingRepository{}
	err = videoMeetingRepository.Create(&videoMeeting)
	if err != nil {
		log.Println("action=CreateVideoMeeting failed to create videoMeeting")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": videoMeeting,
	})
}

func UpdateVideoMeeting(c *gin.Context){
	videoMeeting := models.VideoMeeting{}
	err := c.BindJSON(&videoMeeting)
	if err != nil {
		log.Println("action=UpdateVideoMeeting bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	videoMeetingRepository := repository.VideoMeetingRepository{}
	err = videoMeetingRepository.Update(&videoMeeting)
	if err != nil {
		log.Println("action=UpdateVideoMeeting failed to update videoMeeting")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}
func DeleteVideoMeeting(c *gin.Context){
	paramVideoMeetingId, err := strconv.ParseInt(c.Param("videoMeetingId"), 10, 64)
	if err != nil {
		log.Println("action=DeleteVideoMeeting user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	videoMeetingRepository := repository.VideoMeetingRepository{}
	err = videoMeetingRepository.Delete(paramVideoMeetingId)
	if err != nil {
		log.Println("action=DeleteVideoMeeting failed to delete videoMeeting")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}
func GetVideoMeetingByUserId(c *gin.Context){
	tokenService := service.TokenService{}
	user, err := tokenService.Verify(c)
	if err != nil {
		log.Println("action=GetVideoMeetingByUserId user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	videoMeetingRepository := repository.VideoMeetingRepository{}
	videoMeetings, err := videoMeetingRepository.GetByUserId(user.UserId)
	if err != nil {
		log.Println("action=GetVideoMeetingByUserId failed to get videoMeetings")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": *videoMeetings,
	})
}

func GetVideoMeetingByTransactionId(c *gin.Context){
	paramTransactionId, err := strconv.ParseInt(c.Param("transactionId"), 10, 64)
	if err != nil {
		log.Println("action=GetVideoMeetingByTransactionId user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	videoMeetingRepository := repository.VideoMeetingRepository{}
	videoMeetings, err := videoMeetingRepository.GetByTransactionId(paramTransactionId)
	if err != nil {
		log.Println("action=GetVideoMeetingByTransactionId failed to get videoMeetings")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": *videoMeetings,
	})
}