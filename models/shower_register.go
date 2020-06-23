package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ShowerRegister struct {
	gorm.Model
	BeginDate  time.Time `json:"beginDate"`
	EndDate 	 time.Time `json:"endDate"`
  UserID   	 uint 		 `json:"userId"`
}
