package controller

import (
	"testing"

	"github.com/SushanthGunjal/BooleanAsAService/database"
	"github.com/SushanthGunjal/BooleanAsAService/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testPostRequest struct {
	Value bool
	Key   string
}

func TestUpdateServiceById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepositoryController(ctrl)
	var testRequest PostRequest
	testRequest.Value = true
	testRequest.Key = "fortesting"
	mockRepo.EXPECT().DatabaseSave(&database.Services{ID: "testID", Key: "fortesting", Value: true}).DoAndReturn(func(service *database.Services) error {
		service.Key = "fortesting"
		service.ID = "testID"
		service.Value = true
		return nil
	})
	mockRepo.EXPECT().DatabaseGet(&database.Services{ID: "testID"}).DoAndReturn(func(service *database.Services) error {
		service.ID = "testID"
		service.Key = "fortesting"
		service.Value = true
		return nil
	})
	database.SetRepoController(mockRepo)
	repository = database.GetRepoController()
	service, isServicePresent := UpdateServiceById("testID", testRequest)
	assert.Equal(t, service, &database.Services{ID: "testID", Key: "fortesting", Value: true})
	assert.Equal(t, isServicePresent, true)

}

func TestDeleteById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepositoryController(ctrl)
	mockRepo.EXPECT().DatabaseDelete(&database.Services{ID: "testID"}).DoAndReturn(func(service *database.Services) error {
		return nil
	})
	mockRepo.EXPECT().DatabaseGet(&database.Services{ID: "testID"}).DoAndReturn(func(service *database.Services) error {
		return nil
	})

	database.SetRepoController(mockRepo)
	repository = database.GetRepoController()
	isServicePresent := DeleteById("testID")
	assert.Equal(t, isServicePresent, true)

}

func TestGetServiceByValidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepositoryController(ctrl)

	mockRepo.EXPECT().DatabaseGet(&database.Services{ID: "testID"}).DoAndReturn(func(service *database.Services) error {
		service.Key = "testingkey"
		return nil
	})

	database.SetRepoController(mockRepo)
	repository = database.GetRepoController()

	service, isServicePresent := GetServiceByID("testID")
	assert.Equal(t, service, &database.Services{ID: "testID", Key: "testingkey"})
	assert.Equal(t, isServicePresent, true)
}

func TestGetServiceByInValidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepositoryController(ctrl)

	mockRepo.EXPECT().DatabaseGet(gomock.Any()).Return(nil)
	database.SetRepoController(mockRepo)
	repository = database.GetRepoController()

	service, isServicePresent := GetServiceByID("")
	assert.Equal(t, service, &database.Services{})
	assert.Equal(t, isServicePresent, true)
}
