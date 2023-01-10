package film

import (
	"github.com/gin-gonic/gin"
	"go-netflix-microservice/datasource"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type MockEmptyStore struct {
}

func (store MockEmptyStore) FindById(string) *datasource.FilmRecord {
	return nil
}

func (store MockEmptyStore) FindByGenreAndScoreHigherThan(string, float32) *[]datasource.FilmRecord {
	return nil
}

var joker = datasource.FilmRecord{
	Id:                  1,
	Title:               "Joker",
	Type:                "Movie",
	Genre:               "Crime, Drama, Thriller",
	CountryAvailability: "Lithuania,Poland,France,Italy,Spain,Greece,Belgium,Portugal,Netherlands,Germany,Iceland,Czech Republic",
	Languages:           "English",
	Score:               3.5,
	Tags:                "Dark Comedies,Crime Comedies,Dramas,Comedies,Crime Dramas,Swedish Movies",
	Summary:             "A practical jokers fake kidnapping at a bachelor party turns into a real abduction, forcing him to infiltrate a terrorist group and rescue the groom.",
}

type MockStore struct {
}

func (store MockStore) FindById(string) *datasource.FilmRecord {
	return &joker
}

func (store MockStore) FindByGenreAndScoreHigherThan(string, float32) *[]datasource.FilmRecord {
	return &[]datasource.FilmRecord{joker}
}

func TestFindByIdReturnsBadRequestWhenRecordDoesntFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	datasource.DbStore = MockEmptyStore{}
	context, _ := gin.CreateTestContext(httptest.NewRecorder())

	FindById(context)

	if context.Writer.Status() != http.StatusNotFound {
		t.Errorf("Expected 404 Status Not Found, but got %v", context.Writer.Status())
	}
}

func TestFindByGenreAndScoreHigherThanReturnsBadRequestWhenRecordsDontFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	datasource.DbStore = MockEmptyStore{}
	context := testContext("?genre=Documentary&score_from=9.0")

	Find(context)

	if context.Writer.Status() != http.StatusNotFound {
		t.Errorf("Expected 404 Status Not Found, but got %v", context.Writer.Status())
	}
}

func TestFindByIdReturnsOkWhenRecordFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	datasource.DbStore = MockStore{}
	context, _ := gin.CreateTestContext(httptest.NewRecorder())

	FindById(context)

	if context.Writer.Status() != http.StatusOK {
		t.Errorf("Expected 200 Status Ok, but got %v", context.Writer.Status())
	}
}

func TestFindByGenreAndScoreHigherThanReturnsOkWhenRecordsFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	datasource.DbStore = MockStore{}
	context := testContext("?genre=Documentary&score_from=9.0")

	Find(context)

	if context.Writer.Status() != http.StatusOK {
		t.Errorf("Expected 200 Status Ok, but got %v", context.Writer.Status())
	}
}

func testContext(rawURL string) *gin.Context {
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	parsedUrl, _ := url.Parse(rawURL)

	context.Request = &http.Request{
		URL: parsedUrl,
	}

	return context
}
