package postgresql_test

import (
	"testing"

	"github.com/GritselMaks/BT_API/internal/store"
	"github.com/GritselMaks/BT_API/internal/store/models"
	"github.com/GritselMaks/BT_API/internal/store/postgresql"
	"github.com/stretchr/testify/assert"
)

func TestArticleRepository_Create(t *testing.T) {
	s, teardown := postgresql.TestStore(t)
	defer teardown("article")
	err := s.Articles().Create(&models.Article{
		Date: "2022-11-29",
	})
	assert.NoError(t, err)
}

func TestArticleRepositoryFindByDate(t *testing.T) {
	s, teardown := postgresql.TestStore(t)
	defer teardown("article")
	date := "2022-11-29"

	_, err := s.Articles().ShowArticlebByDate(date)
	assert.EqualError(t, err, store.ErrNotFound.Error())

	s.Articles().Create(&models.Article{Date: date, Title: "Title", Explanation: "Explanation"})
	article, err := s.Articles().ShowArticlebByDate(date)
	assert.NoError(t, err)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, date, article.Date)
}

func TestArticleRepositoryShowArticles(t *testing.T) {
	s, teardown := postgresql.TestStore(t)
	defer teardown("article")
	a1 := models.Article{Date: "2022-11-29", Title: "Title", Explanation: "Explanation"}
	a2 := models.Article{Date: "2022-11-30", Title: "Title", Explanation: "Explanation"}

	res, err := s.Articles().ShowArticles()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(res))
	s.Articles().Create(&a1)
	s.Articles().Create(&a2)

	res, err = s.Articles().ShowArticles()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(res))
}
