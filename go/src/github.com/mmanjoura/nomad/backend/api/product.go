package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mmanjoura/nomad/backend"
	"github.com/mmanjoura/nomad/backend/product"
	js "github.com/mmanjoura/nomad/backend/serializer/json"
	"github.com/pkg/errors"
)

type ProductHandler interface {
	FindProductByID(http.ResponseWriter, *http.Request)
	FindProducts(http.ResponseWriter, *http.Request)
	CreateProduct(http.ResponseWriter, *http.Request)
	UpdateProduct(http.ResponseWriter, *http.Request)
	DeleteProduct(http.ResponseWriter, *http.Request)
}

func (s *Server) productSerializer(contentType string) product.ProductSerializer {

	return &js.Product{}
}

func (s *Server) FindProductByID(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}
	attr, err := s.ProductService.FindOne(r.Context(), userId)
	if err != nil {
		log.Println(err)
	}
	responseBody, err := json.Marshal(attr)

	if err != nil {
		if errors.Cause(err) == product.ErrProductNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	setupResponse(w, contentType, responseBody, http.StatusOK)
}

func (s *Server) FindProducts(w http.ResponseWriter, r *http.Request) {
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
	products, _, err := s.ProductService.FindAll(r.Context(), userId, filter)

	if err != nil {
		if errors.Cause(err) == product.ErrProductNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := json.Marshal(products)

	if err != nil {
		return
	}

	setupResponse(w, contentType, responseBody, http.StatusOK)

}

func (s *Server) CreateProduct(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	attr, err := s.productSerializer(contentType).DecodeProduct(requestBody)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	//We need to get the userId
	userId := 1
	err = s.ProductService.Create(r.Context(), userId, attr)
	if err != nil {
		if errors.Cause(err) == product.ErrProductInvalid {
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

func (s *Server) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	productId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}

	attr, err := s.productSerializer(contentType).DecodeProduct(requestBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = s.ProductService.Update(r.Context(), productId, *attr)
	if err != nil {
		if errors.Cause(err) == product.ErrProductInvalid {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := s.productSerializer(contentType).EncodeProduct(attr)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated)
}

func (s *Server) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	vars := mux.Vars(r)
	productId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}
	err = s.ProductService.Delete(r.Context(), productId)

	if err != nil {
		if errors.Cause(err) == product.ErrProductNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	setupResponse(w, contentType, nil, http.StatusOK)
}
