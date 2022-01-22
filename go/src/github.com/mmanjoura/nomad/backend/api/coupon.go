package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mmanjoura/nomad/backend"
	"github.com/mmanjoura/nomad/backend/coupon"
	js "github.com/mmanjoura/nomad/backend/serializer/json"
	"github.com/pkg/errors"
)

type CouponHandler interface {
	FindCouponByID(http.ResponseWriter, *http.Request)
	FindCoupons(http.ResponseWriter, *http.Request)
	CreateCoupon(http.ResponseWriter, *http.Request)
	UpdateCoupon(http.ResponseWriter, *http.Request)
	DeleteCoupon(http.ResponseWriter, *http.Request)
}

func (s *Server) couponSerializer(contentType string) coupon.CouponSerializer {

	return &js.Coupon{}
}

func (s *Server) FindCouponByID(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}
	attr, err := s.CouponService.FindOne(r.Context(), userId)
	if err != nil {
		log.Println(err)
	}
	responseBody, err := json.Marshal(attr)

	if err != nil {
		if errors.Cause(err) == coupon.ErrCouponNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	setupResponse(w, contentType, responseBody, http.StatusOK)
}

func (s *Server) FindCoupons(w http.ResponseWriter, r *http.Request) {
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
	coupons, _, err := s.CouponService.FindAll(r.Context(), userId, filter)

	if err != nil {
		if errors.Cause(err) == coupon.ErrCouponNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := json.Marshal(coupons)

	if err != nil {
		return
	}

	setupResponse(w, contentType, responseBody, http.StatusOK)

}

func (s *Server) CreateCoupon(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	attr, err := s.couponSerializer(contentType).DecodeCoupon(requestBody)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	//We need to get the userId
	userId := 1
	err = s.CouponService.Create(r.Context(), userId, attr)
	if err != nil {
		if errors.Cause(err) == coupon.ErrCouponInvalid {
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

func (s *Server) UpdateCoupon(w http.ResponseWriter, r *http.Request) {

	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	couponId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}

	attr, err := s.couponSerializer(contentType).DecodeCoupon(requestBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = s.CouponService.Update(r.Context(), couponId, *attr)
	if err != nil {
		if errors.Cause(err) == coupon.ErrCouponInvalid {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := s.couponSerializer(contentType).EncodeCoupon(attr)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated)
}

func (s *Server) DeleteCoupon(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	vars := mux.Vars(r)
	couponId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}
	err = s.CouponService.Delete(r.Context(), couponId)

	if err != nil {
		if errors.Cause(err) == coupon.ErrCouponNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	setupResponse(w, contentType, nil, http.StatusOK)
}
