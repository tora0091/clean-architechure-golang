package jsons

import (
	"clean-architechure-golang/entities"
)

type StatusAndMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type StatusAndArtist struct {
	Status int              `json:"status"`
	Data   *entities.Artist `json:"data"`
}

type StatusAndArtistList struct {
	Status int                `json:"status"`
	Data   []*entities.Artist `json:"data"`
}

type StatusAndCompany struct {
	Status int               `json:"status"`
	Data   *entities.Company `json:"data"`
}

type StatusAndCompanyList struct {
	Status int                 `json:"status"`
	Data   []*entities.Company `json:"data"`
}

type StatusAndMusic struct {
	Status int             `json:"status"`
	Data   *entities.Music `json:"data"`
}

type StatusAndMusicList struct {
	Status int               `json:"status"`
	Data   []*entities.Music `json:"data"`
}

type StatusAndMusicStructResponse struct {
	Status int                           `json:"status"`
	Data   *entities.MusicStructResponse `json:"data"`
}

type StatusAndMusicListStructResponse struct {
	Status int                             `json:"status"`
	Data   []*entities.MusicStructResponse `json:"data"`
}
