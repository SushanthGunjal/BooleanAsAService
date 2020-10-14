package controller

import "github.com/SushanthGunjal/BooleanAsAService/database"

type PostRequest struct {
	Value bool   `json:"value"`
	Key   string `json:"key"`
}

var repository database.RepositoryController
