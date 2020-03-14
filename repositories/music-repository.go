package repositories

type MusicRepository interface {
	FindAll()
}

type musicRepository struct{}

func NewMusicRepository() MusicRepository {
	return &musicRepository{}
}

func (r *musicRepository) FindAll() {

}
