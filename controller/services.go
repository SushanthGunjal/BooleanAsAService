package controller

import (
	"fmt"
	"log"

	"github.com/SushanthGunjal/BooleanAsAService/database"
	"github.com/google/uuid"
)

func init() {
	database.SetRepoController(&database.MonitorRepo{})
	repository = database.GetRepoController()
}

func AddServiceIntoDB(postRequest *PostRequest) (*database.Services, error) {
	id := uuid.New().String()
	service := database.Services{
		ID:    id,
		Key:   postRequest.Key,
		Value: postRequest.Value,
	}
	err := repository.DatabaseCreate(&service)
	if err != nil {
		fmt.Println("Could not add the service into the databse")
		log.Fatal(err)
	}
	//DB.Create(service)
	return &service, nil
}

func GetServiceByID(id string) (*database.Services, bool) {
	service := database.Services{
		ID: id,
	}
	err := repository.DatabaseGet(&service)
	if err != nil {
		fmt.Println("Could get the service from the databse")
		//log.Fatal(err)
		return &service, false
	}
	// DB.First(&service)
	return &service, true
}

func UpdateServiceById(id string, patchRequest PostRequest) (*database.Services, bool) {

	service, isServicePresent := GetServiceByID(id)
	if !isServicePresent {
		return &database.Services{}, false
	}
	service.Value = patchRequest.Value
	service.Key = patchRequest.Key
	err := repository.DatabaseSave(service)
	if err != nil {
		log.Fatal(err)
	}
	return service, true

}

func DeleteById(id string) bool {
	service, isServicePresent := GetServiceByID(id)
	if !isServicePresent {
		return false
	}
	err := repository.DatabaseDelete(service)
	if err != nil {
		log.Fatal(err)
	}
	return true
}
