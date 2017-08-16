package server

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	goscaleio "github.com/codedellemc/goscaleio"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	negroni "github.com/urfave/negroni"

	config "github.com/dvonthenen/kubernetes-scaleio-prom/config"
)

//RestServer representation for a REST API server
type RestServer struct {
	Config  *config.Config
	scaleIO *goscaleio.Client
	Server  *negroni.Negroni
}

//NewRestServer generates a new REST API server
func NewRestServer(cfg *config.Config) *RestServer {
	restServer := &RestServer{
		Config:  cfg,
		scaleIO: nil,
	}

	mux := mux.NewRouter()
	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		err := restServer.getScaleIOStats()
		if err != nil {
			log.Errorln("getScaleIOStats Failed:", err)
		}
		promhttp.Handler().ServeHTTP(w, r)
	}).Methods("GET")

	server := negroni.Classic()
	server.UseHandler(mux)

	restServer.Server = server

	return restServer
}
