package webServer

import (
	"github.com/gorilla/mux"
	"lesson-5-hw-crud/common"
	"net/http"
)

type webServer struct {
	router *mux.Router
	addr   string
	log    common.Logger
}

func New(addr string, log common.Logger) (ws *webServer) {
	log.Debugf("%v", "Creating web server....")
	ws = new(webServer)
	ws.addr = addr
	ws.log = log
	ws.router = mux.NewRouter()
	ws.router.StrictSlash(true)
	return
}

func (ws *webServer) AddRoute(path string, handler func(http.ResponseWriter, *http.Request), method string) {
	r := ws.router.HandleFunc(path, handler)
	if method != "" {
		r.Methods(method)
	}
	ws.log.Debugf("Web server route for %s registered.", path)
}

func (ws *webServer) AddRouteWithParams(path string, handler func(http.ResponseWriter, *http.Request, map[string]string), method string) {
	r := ws.router.HandleFunc(path, withParams(handler))
	if method != "" {
		r.Methods(method)
	}
	ws.log.Debugf("Web server route for %s registered.", path)
}

func (ws *webServer) Start() {
	ws.log.Infof("Server successfully started on %s", ws.addr)

	if err := http.ListenAndServe(ws.addr, ws.router); err != nil {
		ws.log.Infof("Failed listen: %s", err.Error())
	}
}

func withParams(handler func(w http.ResponseWriter, r *http.Request, params map[string]string)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, mux.Vars(r))
	}
}
