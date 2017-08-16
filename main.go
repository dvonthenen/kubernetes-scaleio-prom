package main

import (
	"flag"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/dvonthenen/kubernetes-scaleio-prom/config"
	"github.com/dvonthenen/kubernetes-scaleio-prom/server"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.Infoln("Initializing the Prometheus metric collection for ScaleIO...")
}

func main() {
	cfg := config.NewConfig()
	fs := flag.NewFlagSet("kubernetes-scaleio-prom", flag.ExitOnError)
	cfg.AddFlags(fs)
	fs.Parse(os.Args[1:])

	level, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Warnln("Invalid log level. Defaulting to info.")
		level = log.InfoLevel
	} else {
		log.Infoln("Set logging to", cfg.LogLevel)
	}
	log.SetLevel(level)

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		log.Debugln(pair[0], "=", pair[1])
	}

	restServer := server.NewRestServer(cfg)
	restServer.Server.Run(":" + strconv.Itoa(cfg.RestPort))
}
