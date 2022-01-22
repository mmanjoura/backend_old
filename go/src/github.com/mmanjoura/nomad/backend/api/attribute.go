package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mmanjoura/nomad/backend"
	"github.com/mmanjoura/nomad/backend/attribute"
	js "github.com/mmanjoura/nomad/backend/serializer/json"
	"github.com/pkg/errors"
)

type AttributeHandler interface {
	FindAttributeByID(http.ResponseWriter, *http.Request)
	FindAttributes(http.ResponseWriter, *http.Request)
	CreateAttribute(http.ResponseWriter, *http.Request)
	UpdateAttribute(http.ResponseWriter, *http.Request)
	DeleteAttribute(http.ResponseWriter, *http.Request)
}

func (s *Server) attributeSerializer(contentType string) attribute.AttributeSerializer {

	return &js.Attribute{}
}

func (s *Server) FindAttributeByID(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}
	attr, err := s.AttributeService.FindOne(r.Context(), userId)
	if err != nil {
		log.Println(err)
	}
	responseBody, err := json.Marshal(attr)

	if err != nil {
		if errors.Cause(err) == attribute.ErrAttributeNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	setupResponse(w, contentType, responseBody, http.StatusOK)
}

func (s *Server) FindAttributes(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	// vars := mux.Vars(r)
	// userId, err := strconv.Atoi(vars["id"])
	// if err != nil {
	// 	log.Println(err)
	// }

	userId := 1

	// this filter is coming from ui
	filter := backend.Filter{}
	filter.Limit = 3
	attributes, _, err := s.AttributeService.FindAll(r.Context(), userId, filter)

	if err != nil {
		if errors.Cause(err) == attribute.ErrAttributeNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := json.Marshal(attributes)

	if err != nil {
		return
	}

	setupResponse(w, contentType, responseBody, http.StatusOK)

}

func (s *Server) CreateAttribute(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	attr, err := s.attributeSerializer(contentType).DecodeAttribute(requestBody)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	//We need to get the userId
	userId := 1
	err = s.AttributeService.Create(r.Context(), userId, attr)
	if err != nil {
		if errors.Cause(err) == attribute.ErrAttributeInvalid {
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

func (s *Server) UpdateAttribute(w http.ResponseWriter, r *http.Request) {

	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	attributeId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}

	attr, err := s.attributeSerializer(contentType).DecodeAttribute(requestBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = s.AttributeService.Update(r.Context(), attributeId, *attr)
	if err != nil {
		if errors.Cause(err) == attribute.ErrAttributeInvalid {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := s.attributeSerializer(contentType).EncodeAttribute(attr)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated)
}

func (s *Server) DeleteAttribute(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	vars := mux.Vars(r)
	attributeId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}
	err = s.AttributeService.Delete(r.Context(), attributeId)

	if err != nil {
		if errors.Cause(err) == attribute.ErrAttributeNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	setupResponse(w, contentType, nil, http.StatusOK)
}
