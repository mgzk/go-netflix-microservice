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

type Store interface {
	FindById(id string) *FilmRecord
	FindByGenreAndScoreHigherThan(genre string, score float32) *[]FilmRecord
}

var DbStore Store

type DatabaseStore struct {
	Db *sql.DB
}

func (store DatabaseStore) FindById(id string) *FilmRecord {
	idInt, err := strconv.Atoi(id)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	row := store.Db.QueryRow(selectByIdQuery, idInt)

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

func (store DatabaseStore) FindByGenreAndScoreHigherThan(genre string, score float32) *[]FilmRecord {
	rows, err := store.Db.Query(selectByGenreAndScoreHigherThanOrderByScoreDesc, "%"+genre+"%", score)

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

func InitDatabaseStore(datasource string) {
	database, err := sql.Open("postgres", datasource)

	if err != nil {
		panic(err.Error())
	}

	DbStore = DatabaseStore{
		Db: database,
	}
}
