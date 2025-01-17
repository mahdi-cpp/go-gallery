package repository_drawer

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-gallery/cache"
	"github.com/mahdi-cpp/api-go-gallery/model"
	"github.com/mahdi-cpp/api-go-gallery/utils"
)

func RestMusic() map[string]any {
	return gin.H{
		"name":   "Music",
		"photos": musicDTO.Photos,
	}
}

var musicDTO MusicDTO

type MusicDTO struct {
	Name   string
	Photos []model.PhotoBase
}

func GetDrawerMusic(folder string) {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var count = len(photos)

	var index = 0
	var nameIndex = 0

	for i := 0; i < count; i++ {
		var photoBase model.PhotoBase

		if nameIndex >= len(utils.MovieNames) {
			nameIndex = 0
		}
		musicDTO.Name = utils.MovieNames[nameIndex]

		photoBase = photos[index]
		photoBase.Key = -1
		photoBase.ThumbSize = 540
		photoBase.Crop = 1
		photoBase.Round = 0
		photoBase.PaintWidth = 1080
		photoBase.PaintHeight = 1080

		musicDTO.Photos = append(musicDTO.Photos, photoBase)
		nameIndex++
		index++
	}

	index = 0
}
