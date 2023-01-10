package film

import (
	"github.com/gin-gonic/gin"
	"go-netflix-microservice/datasource"
	"net/http"
)

type Film struct {
	Title               string  `json:"title"`
	Genre               string  `json:"genre"`
	Tags                string  `json:"tags"`
	Languages           string  `json:"languages"`
	Type                string  `json:"type"`
	Score               float32 `json:"score"`
	CountryAvailability string  `json:"country-availability"`
	Summary             string  `json:"summary"`
}

type Films struct {
	Count int    `json:"count"`
	Films []Film `json:"films"`
}

type Filter struct {
	Genre     string  `form:"genre"`
	ScoreFrom float32 `form:"score_from"`
}

const id = "id"

func FindById(c *gin.Context) {
	id := c.Param(id)

	record := datasource.DbStore.FindById(id)

	if record == nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, mapRecord(record))
}

func Find(c *gin.Context) {
	var filter Filter
	err := c.BindQuery(&filter)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	records := datasource.DbStore.FindByGenreAndScoreHigherThan(filter.Genre, filter.ScoreFrom)

	if records == nil {
		c.Status(http.StatusNotFound)
		return
	}

	films := mapRecordArray(records)

	c.JSON(http.StatusOK, Films{
		Count: len(films),
		Films: films,
	})
}

func mapRecord(record *datasource.FilmRecord) Film {
	return Film{
		Title:               record.Title,
		Genre:               record.Genre,
		Tags:                record.Tags,
		Languages:           record.Languages,
		Type:                record.Type,
		Score:               record.Score,
		CountryAvailability: record.CountryAvailability,
		Summary:             record.Summary,
	}
}

func mapRecordArray(records *[]datasource.FilmRecord) []Film {
	var films []Film

	for _, record := range *records {
		film := mapRecord(&record)
		films = append(films, film)
	}

	return films
}
