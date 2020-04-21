package webService

import (
	"encoding/json"
	"fmt"
	"lesson-5-hw-crud/common"
	"net/http"
)

type webServer interface {
	AddRoute(path string, handler func(http.ResponseWriter, *http.Request), method string)
	AddRouteWithParams(path string, handler func(http.ResponseWriter, *http.Request, map[string]string), method string)
	Start()
}

type resultError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type webService struct {
	usersInteractor usersInteractor
	server          webServer
	log             common.Logger
}

func New(webServer webServer, usersInteractor usersInteractor, log common.Logger) (ws *webService) {
	ws = new(webService)
	ws.log = log
	ws.server = webServer
	ws.usersInteractor = usersInteractor

	ws.initCrudRoutes()

	ws.server.AddRoute("/health", ws.Health, "")
	return
}

func (ws *webService) initCrudRoutes() {
	ws.server.AddRoute("/user", ws.createUser, http.MethodPost)
	ws.server.AddRouteWithParams("/user/{userId}", ws.readUser, http.MethodGet)
	ws.server.AddRouteWithParams("/user/{confId}", ws.updateUser, http.MethodPut)
	ws.server.AddRouteWithParams("/user/{confId}", ws.deleteUser, http.MethodDelete)
}

func (ws *webService) Start() {
	ws.server.Start()
}

func (ws *webService) sendResponse(w http.ResponseWriter, response interface{}, code int) {
	ws.log.Infof("Response %d (%s)", code, http.StatusText(code))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	res, err := json.Marshal(response)
	if err != nil {
		ws.sendError(w, fmt.Sprintf("Marshal response failed: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(res); err != nil {
		ws.log.Error(err.Error())
	}
}

func (ws *webService) sendError(w http.ResponseWriter, message string, code int) {
	ws.log.Errorf("Response %d (%s): %s", code, http.StatusText(code), message)
	res, _ := json.Marshal(resultError{
		Code:    code,
		Message: message,
	})

	if _, err := w.Write(res); err != nil {
		ws.log.Error(err.Error())
	}
}
