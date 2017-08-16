package config

import (
	"flag"
	"strconv"
)

//consts exported out of package
const (
	//VersionInt in INT form
	VersionInt = 1

	//VersionStr in string form
	VersionStr = "0.0.1"

	//DefaultRestPort REST API port
	DefaultRestPort = 80

	//DefaultScaleIOAPIVersion version string
	DefaultScaleIOAPIVersion = "2.0"

	//DefaultScaleIOUsername the username for accessing scaleio APIs
	DefaultScaleIOUsername = "admin"
)

//Config is the representation of the config
type Config struct {
	LogLevel string
	Debug    bool
	RestPort int

	ScaleIOName string
	ScaleIOID   string
	APIVersion  string
	Username    string
	Password    string
	Endpoint    string
}

//AddFlags adds flags to the command line parsing
func (cfg *Config) AddFlags(fs *flag.FlagSet) {
	fs.StringVar(&cfg.LogLevel, "loglevel", cfg.LogLevel, "Set the logging level")
	fs.BoolVar(&cfg.Debug, "debug", cfg.Debug, "Debug mode")
	fs.IntVar(&cfg.RestPort, "port", cfg.RestPort, "Port to serve up REST endpoint")

	fs.StringVar(&cfg.ScaleIOName, "scaleio.name", cfg.ScaleIOName, "ScaleIO Cluster Name")
	fs.StringVar(&cfg.ScaleIOID, "scaleio.id", cfg.ScaleIOID, "ScaleIO Cluster ID")
	fs.StringVar(&cfg.APIVersion, "scaleio.apiversion", cfg.APIVersion, "ScaleIO API Version")
	fs.StringVar(&cfg.Username, "scaleio.username", cfg.Username, "ScaleIO Username")
	fs.StringVar(&cfg.Password, "scaleio.password", cfg.Password, "ScaleIO Password")
	fs.StringVar(&cfg.Endpoint, "scaleio.endpoint", cfg.Endpoint, "REST API Endpoint")
}

//NewConfig creates a new Config object
func NewConfig() *Config {
	return &Config{
		LogLevel:    env("LOG_LEVEL", "info"),
		Debug:       envBool("DEBUG", "false"),
		RestPort:    envInt("REST_PORT", strconv.Itoa(DefaultRestPort)),
		ScaleIOName: env("CLUSTER_NAME", "scaleio"),
		ScaleIOID:   env("CLUSTER_ID", ""),
		APIVersion:  env("API_VERSION", DefaultScaleIOAPIVersion),
		Username:    env("SCALEIO_USERNAME", DefaultScaleIOUsername),
		Password:    env("SCALEIO_PASSWORD", ""),
		Endpoint:    env("SCALEIO_ENDPOINT", ""),
	}
}
