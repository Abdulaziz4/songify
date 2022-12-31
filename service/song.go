package service

import (
	"github.com/Abdulaziz4/songify/model"
	"gorm.io/gorm"
)

type SongService struct {
	db *gorm.DB
}

func (s *SongService) GetAllSongs() ([]model.Song, error) {
	songs := []model.Song{}

	result := s.db.Find(&songs)
	if result.Error != nil {
		return nil, result.Error
	}

	return songs, nil
}

func (s *SongService) GetById(id string) (*model.Song, error) {
	var song model.Song

	result := s.db.Find(&song)
	if result.Error != nil {
		return nil, result.Error
	}

	return &song, nil
}

func (s *SongService) CreateSong(name, url string) (*model.Song, error) {
	song := model.Song{Name: name, Url: url}

	result := s.db.Create(&song)
	if result.Error != nil {
		return nil, result.Error
	}

	return &song, nil
}

func (s *SongService) DeleteSong(id string) error {
	res := s.db.Delete(&model.Song{}, id)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func NewSongService(db *gorm.DB) *SongService {
	return &SongService{db: db}
}
