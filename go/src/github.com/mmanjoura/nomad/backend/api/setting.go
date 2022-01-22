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
	"github.com/mmanjoura/nomad/backend/setting"
	"github.com/pkg/errors"
)

type SettingHandler interface {
	FindSettingByID(http.ResponseWriter, *http.Request)
	FindSettings(http.ResponseWriter, *http.Request)
	CreateSetting(http.ResponseWriter, *http.Request)
	UpdateSetting(http.ResponseWriter, *http.Request)
	DeleteSetting(http.ResponseWriter, *http.Request)
}

func (s *Server) settingSerializer(contentType string) setting.SettingSerializer {

	return &js.Setting{}
}

func (s *Server) FindSettingByID(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}
	attr, err := s.SettingService.FindOne(r.Context(), userId)
	if err != nil {
		log.Println(err)
	}
	responseBody, err := json.Marshal(attr)

	if err != nil {
		if errors.Cause(err) == setting.ErrSettingNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	setupResponse(w, contentType, responseBody, http.StatusOK)
}

func (s *Server) FindSettings(w http.ResponseWriter, r *http.Request) {
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
	settings, _, err := s.SettingService.FindAll(r.Context(), userId, filter)

	if err != nil {
		if errors.Cause(err) == setting.ErrSettingNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := json.Marshal(settings)

	if err != nil {
		return
	}

	setupResponse(w, contentType, responseBody, http.StatusOK)

}

func (s *Server) CreateSetting(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	attr, err := s.settingSerializer(contentType).DecodeSetting(requestBody)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	//We need to get the userId
	userId := 1
	err = s.SettingService.Create(r.Context(), userId, attr)
	if err != nil {
		if errors.Cause(err) == setting.ErrSettingInvalid {
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

func (s *Server) UpdateSetting(w http.ResponseWriter, r *http.Request) {

	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	settingId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}

	attr, err := s.settingSerializer(contentType).DecodeSetting(requestBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = s.SettingService.Update(r.Context(), settingId, *attr)
	if err != nil {
		if errors.Cause(err) == setting.ErrSettingInvalid {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := s.settingSerializer(contentType).EncodeSetting(attr)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated)
}

func (s *Server) DeleteSetting(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	vars := mux.Vars(r)
	settingId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}
	err = s.SettingService.Delete(r.Context(), settingId)

	if err != nil {
		if errors.Cause(err) == setting.ErrSettingNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	setupResponse(w, contentType, nil, http.StatusOK)
}
