package models

import (
	"github.com/jinzhu/gorm"
)

type Ranking struct {
	gorm.Model
	Score  	 int64 `json:"score"`
  UserID   uint  `json:"userId"`
}
