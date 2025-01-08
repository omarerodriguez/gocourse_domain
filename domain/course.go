package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	ID        string     `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	Name      string     `json:"name" gorm:"type:char(50);not null"`
	StartDate time.Time  `json:"start_date"`
	EndDate   time.Time  `json:"end_date"`
	User      *User      `gorm:"-"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

func (course *Course) BeforeCreate(tx *gorm.DB) (err error) {
	if course.ID == "" {
		course.ID = uuid.New().String()
	}
	return
}
