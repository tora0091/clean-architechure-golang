package entities

import "time"

type Music struct {
	ID        int       `json:"id" gorm:"primary_key, column:id"`
	ISWC      string    `json:"iswc" gorm:"column:iswc"`
	Title     string    `json:"title" gorm:"column:title"`
	Time      int       `json:"time" gorm:"column:time"`
	Genre     string    `json:"genre" gorm:"column:genre"`
	Artist    int       `json:"artist" gorm:"column:artist"`
	Company   int       `json:"company" gorm:"column:company"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}

type Artist struct {
	ID        int       `json:"id" gorm:"primary_key, column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Email     string    `json:"email" gorm:"column:email"`
	Birth     string    `json:"birth" gorm:"column:birth"`
	Gender    string    `json:"gender" gorm:"column:gender"`
	Address   string    `json:"address" gorm:"column:address"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}

type Company struct {
	ID        int       `json:"id" gorm:"primary_key, column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Email     string    `json:"email" gorm:"column:email"`
	Address   string    `json:"address" gorm:"column:address"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}
