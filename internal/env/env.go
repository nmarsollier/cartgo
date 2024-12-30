package env

import (
	"cmp"
	"os"

	"github.com/nmarsollier/commongo/strs"
)

// Configuration properties
type Configuration struct {
	ServerName        string `json:"serverName"`
	Port              int    `json:"port"`
	GqlPort           int    `json:"gqlPort"`
	RabbitURL         string `json:"rabbitUrl"`
	MongoURL          string `json:"mongoUrl"`
	SecurityServerURL string `json:"securityServerUrl"`
	CatalogServerURL  string `json:"catalogUrl"`
	FluentURL         string `json:"fluentUrl"`
}

var config *Configuration

// Get Obtiene las variables de entorno del sistema
func Get() *Configuration {
	if config == nil {
		config = load()
	}

	return config
}

// Load file properties
func load() *Configuration {
	return &Configuration{
		ServerName:        cmp.Or(os.Getenv("SERVER_NAME"), "cartgo"),
		Port:              cmp.Or(strs.AtoiZero(os.Getenv("PORT")), 3003),
		GqlPort:           cmp.Or(strs.AtoiZero(os.Getenv("GQL_PORT")), 4003),
		RabbitURL:         cmp.Or(os.Getenv("RABBIT_URL"), "amqp://localhost"),
		MongoURL:          cmp.Or(os.Getenv("MONGO_URL"), "mongodb://localhost:27017"),
		SecurityServerURL: cmp.Or(os.Getenv("AUTH_SERVICE_URL"), "http://localhost:3000"),
		CatalogServerURL:  cmp.Or(os.Getenv("CATALOG_SERVICE_URL"), "http://localhost:3002"),
		FluentURL:         cmp.Or(os.Getenv("FLUENT_URL"), "localhost:24224"),
	}
}
