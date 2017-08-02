package server

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	negroni "github.com/urfave/negroni"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	config "github.com/dvonthenen/kubernetes-scaleio-prom/config"

	goscaleio "github.com/codedellemc/goscaleio"
)

//RestServer representation for a REST API server
type RestServer struct {
	Config  *config.Config
	Server  *negroni.Negroni
	scaleIO *goscaleio.Client
	promCounters *prometheus.SummaryVec

	sync.Mutex
}

//NewRestServer generates a new REST API server
func NewRestServer(cfg *config.Config) *RestServer {
	restServer := &RestServer{
		Config: cfg,
		scaleIO: nil,
		promCounters: nil,
	}

	mux := mux.NewRouter()
	mux.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		getVersion(w, r, restServer)
	}).Methods("GET")
	/*
		mux.HandleFunc("/api/state", func(w http.ResponseWriter, r *http.Request) {
			getState(w, r, restServer)
		}).Methods("GET")
		mux.HandleFunc("/api/state", func(w http.ResponseWriter, r *http.Request) {
			setState(w, r, restServer)
		}).Methods("POST")
		mux.HandleFunc("/api/node/state", func(w http.ResponseWriter, r *http.Request) {
			setNodeState(w, r, restServer)
		}).Methods("POST")
		mux.HandleFunc("/api/node/device", func(w http.ResponseWriter, r *http.Request) {
			setNodeDevices(w, r, restServer)
		}).Methods("POST")
		mux.HandleFunc("/api/node/ping", func(w http.ResponseWriter, r *http.Request) {
			setNodePing(w, r, restServer)
		}).Methods("POST")
		mux.HandleFunc("/api/fake", func(w http.ResponseWriter, r *http.Request) {
			setFakeData(w, r, restServer)
		}).Methods("POST")
		mux.HandleFunc("/ui", func(w http.ResponseWriter, r *http.Request) {
			displayState(w, r, restServer)
		}).Methods("GET")
		//TODO delete this below when a real UI is embedded
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			displayState(w, r, restServer)
		}).Methods("GET")
	*/
	server := negroni.Classic()
	server.UseHandler(mux)

	restServer.Server = server

	return restServer
}
