package webService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"lesson-5-hw-crud/domain"
)

type usersInteractor interface {
	Create(user *domain.User) (*domain.User, error)
	Read(id int64) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	Delete(id int64) error
}

func (ws *webService) createUser(w http.ResponseWriter, r *http.Request) {
	user, err := ws.readUserBody(r)
	if err != nil {
		ws.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	resUser, err := ws.usersInteractor.Create(user)
	if err != nil {
		ws.sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ws.sendResponse(w, resUser, http.StatusOK)
}

func (ws *webService) readUser(w http.ResponseWriter, r *http.Request, params map[string]string) {
	var (
		userId int64
		err    error
	)

	userId, err = ws.readUserId(params)
	if err != nil {
		ws.sendError(w, err.Error(), http.StatusBadRequest)
	}

	resUser, err := ws.usersInteractor.Read(userId)
	if err != nil {
		ws.sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ws.sendResponse(w, resUser, http.StatusOK)
}

func (ws *webService) updateUser(w http.ResponseWriter, r *http.Request, params map[string]string) {
	user, err := ws.readUserBody(r)
	if err != nil {
		ws.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID, err = ws.readUserId(params)
	if err != nil {
		ws.sendError(w, err.Error(), http.StatusBadRequest)
	}

	resUser, err := ws.usersInteractor.Update(user)
	if err != nil {
		ws.sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ws.sendResponse(w, resUser, http.StatusOK)
}

func (ws *webService) deleteUser(w http.ResponseWriter, r *http.Request, params map[string]string) {
	userId, err := ws.readUserId(params)
	if err != nil {
		ws.sendError(w, err.Error(), http.StatusBadRequest)
	}

	err = ws.usersInteractor.Delete(userId)
	if err != nil {
		ws.sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ws.sendResponse(w, nil, http.StatusOK)
}

func (ws *webService) readUserId(params map[string]string) (userId int64, err error) {
	if v, ok := params["userId"]; !ok || v == "" {
		err = fmt.Errorf("userId not set in params")
		return
	} else {
		userId, _ = strconv.ParseInt(v, 10, 63)
	}
	return
}

func (ws *webService) readUserBody(r *http.Request) (user *domain.User, err error) {
	user = new(domain.User)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(reqBody, user)
	if err != nil {
		return
	}

	return
}
