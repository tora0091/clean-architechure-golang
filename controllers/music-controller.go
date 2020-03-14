package controllers

type MusicController interface {
	FindAll()
}

type musicController struct{}

func NewMusicController() MusicController {
	return &musicController{}
}

func (c *musicController) FindAll() {

}
