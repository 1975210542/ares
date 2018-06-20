package models

import (
	"time"

	//	"github.com/jinzhu/gorm"
	//	"github.com/satori/go.uuid"
)

type BaseModel struct {
	CreateAt  time.Time  `json:"-"`
	UpdateAt  time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
