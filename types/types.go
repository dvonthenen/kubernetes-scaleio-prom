package types

//Version describes the version of the REST API
type Version struct {
	VersionInt int               `json:"versionint"`
	VersionStr string            `json:"versionstr"`
	BuildStr   string            `json:"buildstr,omitempty"`
	KeyValue   map[string]string `json:"keyvalue,omitempty"`
}
