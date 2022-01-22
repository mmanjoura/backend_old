package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mmanjoura/nomad/backend"
	js "github.com/mmanjoura/nomad/backend/serializer/json"
	"github.com/mmanjoura/nomad/backend/shipping"
	"github.com/pkg/errors"
)

type ShippingHandler interface {
	FindShippingByID(http.ResponseWriter, *http.Request)
	FindShippings(http.ResponseWriter, *http.Request)
	CreateShipping(http.ResponseWriter, *http.Request)
	UpdateShipping(http.ResponseWriter, *http.Request)
	DeleteShipping(http.ResponseWriter, *http.Request)
}

func (s *Server) shippingSerializer(contentType string) shipping.ShippingSerializer {

	return &js.Shipping{}
}

func (s *Server) FindShippingByID(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}
	attr, err := s.ShippingService.FindOne(r.Context(), userId)
	if err != nil {
		log.Println(err)
	}
	responseBody, err := json.Marshal(attr)

	if err != nil {
		if errors.Cause(err) == shipping.ErrShippingNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	setupResponse(w, contentType, responseBody, http.StatusOK)
}

func (s *Server) FindShippings(w http.ResponseWriter, r *http.Request) {
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
	shippings, _, err := s.ShippingService.FindAll(r.Context(), userId, filter)

	if err != nil {
		if errors.Cause(err) == shipping.ErrShippingNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := json.Marshal(shippings)

	if err != nil {
		return
	}

	setupResponse(w, contentType, responseBody, http.StatusOK)

}

func (s *Server) CreateShipping(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	attr, err := s.shippingSerializer(contentType).DecodeShipping(requestBody)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	//We need to get the userId
	userId := 1
	err = s.ShippingService.Create(r.Context(), userId, attr)
	if err != nil {
		if errors.Cause(err) == shipping.ErrShippingInvalid {
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

func (s *Server) UpdateShipping(w http.ResponseWriter, r *http.Request) {

	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	shippingId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}

	attr, err := s.shippingSerializer(contentType).DecodeShipping(requestBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = s.ShippingService.Update(r.Context(), shippingId, *attr)
	if err != nil {
		if errors.Cause(err) == shipping.ErrShippingInvalid {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := s.shippingSerializer(contentType).EncodeShipping(attr)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated)
}

func (s *Server) DeleteShipping(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	vars := mux.Vars(r)
	shippingId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}
	err = s.ShippingService.Delete(r.Context(), shippingId)

	if err != nil {
		if errors.Cause(err) == shipping.ErrShippingNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	setupResponse(w, contentType, nil, http.StatusOK)
}
