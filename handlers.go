package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SushanthGunjal/BooleanAsAService/controller"
	"github.com/gin-gonic/gin"
)

//Adds a Service into the database.
func AddService(c *gin.Context) {
	var postRequest controller.PostRequest
	err := c.BindJSON(&postRequest)
	if err != nil {
		fmt.Println("incorrect post request format")
	}
	service, err := controller.AddServiceIntoDB(&postRequest)
	if err != nil {
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":    service.ID,
			"key":   service.Key,
			"value": service.Value,
		})
	}
}

//Fetches a service from the database.
func GetService(c *gin.Context) {
	id := c.Param("id")
	service, isServicePresent := controller.GetServiceByID(id)
	if isServicePresent {
		c.JSON(http.StatusOK, gin.H{
			"id":    service.ID,
			"value": service.Value,
			"key":   service.Key,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"id": "invalid",
		})
	}
}

//Updates a service in the database.
func UpdateService(c *gin.Context) {
	var patchRequest controller.PostRequest
	err := c.BindJSON(&patchRequest)
	if err != nil {
		fmt.Println("incorect input format")
		log.Fatal(err)
	}
	id := c.Param("id")
	if service, isServicePresent := controller.UpdateServiceById(id, patchRequest); isServicePresent {
		c.JSON(http.StatusOK, gin.H{
			"id":    service.ID,
			"value": service.Value,
			"key":   service.Key,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"id": "invalid",
		})
	}
}

//Deletes a service from the database.
func DeleteService(c *gin.Context) {
	id := c.Param("id")
	isServicePresent := controller.DeleteById(id)
	if !isServicePresent {
		c.JSON(http.StatusNotFound, gin.H{
			"id": "invalid",
		})
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}

}
