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
