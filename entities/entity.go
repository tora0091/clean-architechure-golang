package entities

import "time"

type Music struct {
	ID        int        `json:"id" gorm:"primary_key, autoincrement, column:id"`
	ISWC      string     `json:"iswc" gorm:"column:iswc"`
	Title     string     `json:"title" gorm:"column:title"`
	Time      int        `json:"time" gorm:"column:time"`
	Genre     string     `json:"genre" gorm:"column:genre"`
	ArtistID  int        `json:"artist_id" gorm:"column:artist_id"`
	CompanyID int        `json:"company_id" gorm:"column:company_id"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

type Artist struct {
	ID        int        `json:"id" gorm:"primary_key, autoincrement, column:id"`
	Name      string     `json:"name" gorm:"column:name"`
	Email     string     `json:"email" gorm:"column:email"`
	Birth     string     `json:"birth" gorm:"column:birth"`
	Gender    string     `json:"gender" gorm:"column:gender"`
	Address   string     `json:"address" gorm:"column:address"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

type Company struct {
	ID        int        `json:"id" gorm:"primary_key, autoincrement, column:id"`
	Name      string     `json:"name" gorm:"column:name"`
	Email     string     `json:"email" gorm:"column:email"`
	Address   string     `json:"address" gorm:"column:address"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}
