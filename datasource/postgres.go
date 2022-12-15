package datasource

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

type FilmRecord struct {
	Id                  int
	Title               string
	Genre               string
	Tags                string
	Languages           string
	Type                string
	Score               float32
	CountryAvailability string
	Summary             string
}

const selectByIdQuery = "select id, title, genre, tags, languages, type, score, country_availability, summary from film where id = $1"
const selectByGenreAndScoreHigherThanOrderByScoreDesc = "select id, title, genre, tags, languages, type, score, country_availability, " +
	"summary from film where genre like $1 and score >= $2 order by score desc"

var db *sql.DB

func InitDatabase(datasource string) {
	database, err := sql.Open("postgres", datasource)

	if err != nil {
		panic(err.Error())
	}

	db = database
}

func FindById(id string) *FilmRecord {
	idInt, err := strconv.Atoi(id)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	row := db.QueryRow(selectByIdQuery, idInt)

	var record FilmRecord
	if err := row.Scan(
		&record.Id,
		&record.Title,
		&record.Genre,
		&record.Tags,
		&record.Languages,
		&record.Type,
		&record.Score,
		&record.CountryAvailability,
		&record.Summary); err != nil {
		log.Println(err.Error())
		return nil
	}

	return &record
}

func FindByGenreAndScoreHigherThan(genre string, score float32) *[]FilmRecord {
	rows, err := db.Query(selectByGenreAndScoreHigherThanOrderByScoreDesc, "%" + genre + "%", score)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	var records []FilmRecord

	for rows.Next() {
		var record FilmRecord
		if err := rows.Scan(
			&record.Id,
			&record.Title,
			&record.Genre,
			&record.Tags,
			&record.Languages,
			&record.Type,
			&record.Score,
			&record.CountryAvailability,
			&record.Summary); err != nil {
			log.Println(err.Error())
		}
		records = append(records, record)
	}

	return &records
}
