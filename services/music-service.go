package services

type MusicService interface {
	FindAll()
}

type musicService struct{}

func NewMusicService() MusicService {
	return &musicService{}
}

func (s *musicService) FindAll() {

}
