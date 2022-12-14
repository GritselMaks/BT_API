package app

import (
	"net/http"

	"github.com/GritselMaks/BT_API/internal/utils"
	"github.com/gorilla/mux"
)

func (s *Server) GetArticles() http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		s.logger.Debug("starting handler: GetArticles")
		defer s.logger.Debugf("finishing handler: GetArticles")
		articles, err := s.store.Articles().ShowArticles()
		if err != nil {
			s.logger.Errorf("failed getting data from store: %s", err.Error())
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		for i, article := range articles {
			url := utils.PreparePictureUrl(s.config.Host, s.config.Port, article.Date)
			articles[i].Url = url
		}
		ResponseWithJSON(w, http.StatusOK, articles)
	}))
}

func (s *Server) GetArticleWithDate() http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		s.logger.Debugf("starting handler: GetArticleWithDate")
		defer s.logger.Debugf("finishing handler: GetArticleWithDate")
		date, ok := mux.Vars(r)["date"]
		if !ok {
			s.logger.Errorf("failed getting request data")
			RespondWithError(w, http.StatusBadRequest, "Bad request data")
			return
		}
		article, err := s.store.Articles().ShowArticlebByDate(date)
		if err != nil {
			s.logger.Errorf("failed getting data from store: %s", err.Error())
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		url := utils.PreparePictureUrl(s.config.Host, s.config.Port, article.Date)
		article.Url = url
		ResponseWithJSON(w, http.StatusOK, article)
	}))
}

func (s *Server) GetPicture() http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		s.logger.Debugf("starting handler: GetPicture")
		defer s.logger.Debugf("finishing handler: GetPicture")
		date, ok := mux.Vars(r)["date"]
		if !ok {
			s.logger.Errorf("failed getting request data")
			RespondWithError(w, http.StatusBadRequest, "Bad request data")
			return
		}
		picture, err := s.pudgeStore.Get(date)
		if err != nil {
			s.logger.Errorf("failed getting data from store: %s", err.Error())
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		RespondWithPicture(w, http.StatusOK, picture)
	}))
}
