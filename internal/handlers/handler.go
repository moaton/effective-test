package handlers

import (
	"effective-test/internal/models"
	"effective-test/internal/service"
	"effective-test/pkg/logger"
	"effective-test/pkg/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Handler interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) *mux.Router {
	router := mux.NewRouter()
	handler := &handler{
		service: service,
	}

	router.HandleFunc("/users", handler.GetUsers).Methods("GET")
	router.HandleFunc("/user", handler.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", handler.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", handler.DeleteUser).Methods("DELETE")
	return router
}

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	type response struct {
		Data  []models.User `json:"data"`
		Total int64         `json:"total"`
	}

	params := models.Params{
		Limit:  10,
		Offset: 0,
	}
	limitStr := r.FormValue("limit")
	offsetStr := r.FormValue("offset")
	params.Query = r.FormValue("query")
	if params.Query != "" {
		params.Query = strings.ToLower(params.Query)
	}
	if limitStr != "" {
		params.Limit, err = strconv.ParseInt(limitStr, 10, 0)
		if err != nil {
			logger.Errorf("Limit strconv.ParseInt err %v", err)
		}
	}

	if offsetStr != "" {
		params.Offset, err = strconv.ParseInt(offsetStr, 10, 0)
		if err != nil {
			logger.Errorf("Limit strconv.ParseInt err %v", err)
		}
	}

	users, total, err := h.service.GetUsers(ctx, params)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseOk(w, response{
		Data:  users,
		Total: total,
	})
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	type response struct {
		ID int64 `json:"id"`
	}

	var user models.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	if user.Name == "" || user.Surname == "" {
		utils.ResponseError(w, http.StatusBadRequest, "name or surname is empty")
		return
	}

	id, err := h.service.CreateUser(ctx, &user)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseOk(w, response{
		ID: id,
	})
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	type response struct {
		Message string `json:"message"`
	}

	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	var user models.User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&user)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	user.ID = id
	err = h.service.UpdateUser(ctx, &user)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseOk(w, response{
		Message: "success",
	})
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	type response struct {
		Message string `json:"message"`
	}

	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = h.service.DeleteUser(ctx, id)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseOk(w, response{
		Message: "success",
	})
}
