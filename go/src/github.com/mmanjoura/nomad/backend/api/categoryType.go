package api

import (
	"encoding/json"
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

type CategoryTypeHandler interface {
	FindCategoryTypeByID(http.ResponseWriter, *http.Request)
	FindCategoryTypes(http.ResponseWriter, *http.Request)
	CreateCategoryType(http.ResponseWriter, *http.Request)
	UpdateCategoryType(http.ResponseWriter, *http.Request)
	DeleteCategoryType(http.ResponseWriter, *http.Request)
}

func (s *Server) categoryTypeSerializer(contentType string) category.CategorySerializer {

	return &js.Category{}
}

func (s *Server) FindCategoryTypeByID(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}
	shopId := 1
	userId = 1
	attr, err := s.CategoryService.TypeFindOne(r.Context(), userId, shopId)
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

func (s *Server) FindCategoryTypes(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	// vars := mux.Vars(r)
	// userId, err := strconv.Atoi(vars["id"])
	// if err != nil {
	// 	log.Println(err)
	// }

	userId := 1
	shopId := 1

	// this filter is coming from ui
	filter := backend.Filter{}
	filter.Limit = 3
	categoryTypes, _, err := s.CategoryService.TypeFindAll(r.Context(), userId, shopId, filter)

	if err != nil {
		if errors.Cause(err) == category.ErrCategoryNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := json.Marshal(categoryTypes)

	if err != nil {
		return
	}

	setupResponse(w, contentType, responseBody, http.StatusOK)

}

func (s *Server) CreateCategoryType(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	ctg, err := s.categoryTypeSerializer(contentType).DecodeCategoryType(requestBody)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	//We need to get the userId
	userId := 1
	shopId := 1
	err = s.CategoryService.TypeCreate(r.Context(), userId, shopId, ctg)
	if err != nil {
		if errors.Cause(err) == category.ErrCategoryTypeInvalid {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := json.Marshal(ctg)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated)
}

func (s *Server) UpdateCategoryType(w http.ResponseWriter, r *http.Request) {

	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	categoryTypeId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}

	attr, err := s.categoryTypeSerializer(contentType).DecodeCategoryType(requestBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	shopId := 1
	_, err = s.CategoryService.TypeUpdate(r.Context(), categoryTypeId, shopId, *attr)
	if err != nil {
		if errors.Cause(err) == category.ErrCategoryTypeInvalid {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := s.categoryTypeSerializer(contentType).EncodeCategoryType(attr)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated)
}

func (s *Server) DeleteCategoryType(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	vars := mux.Vars(r)
	categoryTypeId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}
	shopId := 1
	err = s.CategoryService.TypeDelete(r.Context(), categoryTypeId, shopId)

	if err != nil {
		if errors.Cause(err) == category.ErrCategoryTypeNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	setupResponse(w, contentType, nil, http.StatusOK)
}
