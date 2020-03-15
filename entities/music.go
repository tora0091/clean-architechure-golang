package entities

import "time"

type Music struct {
	ID        int    `json:"id"`
	ISWC      string `json:"iswc"`
	Title     string `json:"title"`
	Time      int    `json:"time"`
	Genre     string `json:"genre"`
	Artist    int    `json:"artist"`
	Company   int    `json:"company"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Artist struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Birth     string `json:"birth"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Company struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
