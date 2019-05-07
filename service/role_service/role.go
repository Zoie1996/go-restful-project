package role_service

import (
	"go-restful-project/models"
	"go-restful-project/pkg/e"
	"log"
)

// CheckRoleExist 检查用户是否存在
func CheckRoleExist(id int) int {
	exists, err := models.ExistRoleByID(id)
	if err != nil {
		log.Printf("[role] get role fail %s", err)
		// GetInfo := app.GetInfo(http.StatusNotFound, e.ERROR)
		return e.ERROR

	}
	if !exists {
		return e.ERROR_ROLE_NOT_FOUND
	}
	return e.SUCCESS
}
