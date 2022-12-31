package router

import (
	"net/http"

	"github.com/Abdulaziz4/songify/service"
	"github.com/gin-gonic/gin"
)

type SongRouter struct {
	s *service.SongService
}

type CreateSongRequest struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (r *SongRouter) Get(c *gin.Context) {
	songs, err := r.s.GetAllSongs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, songs)
}

func (r *SongRouter) GetById(c *gin.Context) {
	song, err := r.s.GetById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": song})
}

func (r *SongRouter) Post(c *gin.Context) {
	var req CreateSongRequest
	c.BindJSON(&req)

	r.s.CreateSong(req.Name, req.Url)

	c.JSON(http.StatusOK, gin.H{"data": "Song created"})
}

func (r *SongRouter) Delete(c *gin.Context) {
	err := r.s.DeleteSong(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Song deleted"})
}

func NewSongRouter(s *service.SongService) *SongRouter {
	return &SongRouter{s: s}
}
