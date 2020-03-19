package jsons

import (
	"clean-architechure-golang/entities"
)

type ResponseMessage struct {
	// Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseArtist struct {
	// Status int              `json:"status"`
	Data *entities.Artist `json:"data"`
}

type ResponseArtistList struct {
	// Status int                `json:"status"`
	Data []*entities.Artist `json:"data"`
}

type ResponseCompany struct {
	// Status int               `json:"status"`
	Data *entities.Company `json:"data"`
}

type ResponseCompanyList struct {
	// Status int                 `json:"status"`
	Data []*entities.Company `json:"data"`
}

type ResponseMusic struct {
	// Status int             `json:"status"`
	Data *entities.Music `json:"data"`
}

type ResponseMusicList struct {
	// Status int               `json:"status"`
	Data []*entities.Music `json:"data"`
}

type ResponseMusicStructResponse struct {
	// Status int                           `json:"status"`
	Data *entities.MusicStructResponse `json:"data"`
}

type ResponseMusicListStructResponse struct {
	// Status int                             `json:"status"`
	Data []*entities.MusicStructResponse `json:"data"`
}
