package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mmanjoura/nomad/backend"
	"github.com/mmanjoura/nomad/backend/category"
	js "github.com/mmanjoura/nomad/backend/serializer/json"
	"github.com/pkg/errors"
)

type CategoryHandler interface {
	FindCategoryByID(http.ResponseWriter, *http.Request)
	FindCategories(http.ResponseWriter, *http.Request)
	CreateCategory(http.ResponseWriter, *http.Request)
	UpdateCategory(http.ResponseWriter, *http.Request)
	DeleteCategory(http.ResponseWriter, *http.Request)
}

func (s *Server) categorySerializer(contentType string) category.CategorySerializer {

	return &js.Category{}
}

func (s *Server) FindCategoryByID(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}
	shopId := 1
	attr, err := s.CategoryService.FindOne(r.Context(), userId, shopId)
	if err != nil {
		log.Println(err)
	}
	responseBody, err := json.Marshal(attr)

	if err != nil {
		if errors.Cause(err) == category.ErrCategoryNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	setupResponse(w, contentType, responseBody, http.StatusOK)
}

func (s *Server) FindCategories(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")

	vars := mux.Vars(r)
	fmt.Print(vars["searchJoin"])
	//userId, err := strconv.Atoi(vars["id"])
	// if err != nil {
	// 	log.Println(vars)
	// }

	shopId := 1
	categoryId := 1

	// this filter is coming from ui
	filter := backend.Filter{}
	filter.Limit = 3
	categorys, _, err := s.CategoryService.FindAll(r.Context(), categoryId, shopId, filter)

	if err != nil {
		if errors.Cause(err) == category.ErrCategoryNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := json.Marshal(categorys)

	if err != nil {
		return
	}

	setupResponse(w, contentType, responseBody, http.StatusOK)

}

func (s *Server) CreateCategory(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	attr, err := s.categorySerializer(contentType).DecodeCategory(requestBody)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	//We need to get the userId
	userId := 1
	shopId := 1
	err = s.CategoryService.Create(r.Context(), userId, shopId, attr)
	if err != nil {
		if errors.Cause(err) == category.ErrCategoryInvalid {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := json.Marshal(attr)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated)
}

func (s *Server) UpdateCategory(w http.ResponseWriter, r *http.Request) {

	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	categoryId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}

	attr, err := s.categorySerializer(contentType).DecodeCategory(requestBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	shopId := 1
	_, err = s.CategoryService.Update(r.Context(), categoryId, shopId, *attr)
	if err != nil {
		if errors.Cause(err) == category.ErrCategoryInvalid {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := s.categorySerializer(contentType).EncodeCategory(attr)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated)
}

func (s *Server) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	vars := mux.Vars(r)
	categoryId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}
	shopId := 1
	err = s.CategoryService.Delete(r.Context(), categoryId, shopId)

	if err != nil {
		if errors.Cause(err) == category.ErrCategoryNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	setupResponse(w, contentType, nil, http.StatusOK)
}
