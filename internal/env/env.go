package env

import (
	"os"
	"strconv"
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

func new() *Configuration {
	return &Configuration{
		ServerName:        "cartgo",
		Port:              3003,
		GqlPort:           4003,
		RabbitURL:         "amqp://localhost",
		MongoURL:          "mongodb://localhost:27017",
		SecurityServerURL: "http://localhost:3000",
		CatalogServerURL:  "http://localhost:3002",
		FluentURL:         "localhost:24224",
	}
}

// Get Obtiene las variables de entorno del sistema
func Get() *Configuration {
	if config == nil {
		config = load()
	}

	return config
}

// Load file properties
func load() *Configuration {
	result := new()

	if value := os.Getenv("SERVER_NAME"); len(value) > 0 {
		result.ServerName = value
	}

	if value := os.Getenv("RABBIT_URL"); len(value) > 0 {
		result.RabbitURL = value
	}

	if value := os.Getenv("MONGO_URL"); len(value) > 0 {
		result.MongoURL = value
	}

	if value := os.Getenv("PORT"); len(value) > 0 {
		if intVal, err := strconv.Atoi(value); err == nil {
			result.Port = intVal
		}
	}

	if value := os.Getenv("GQL_PORT"); len(value) > 0 {
		if intVal, err := strconv.Atoi(value); err == nil {
			result.GqlPort = intVal
		}
	}

	if value := os.Getenv("AUTH_SERVICE_URL"); len(value) > 0 {
		result.SecurityServerURL = value
	}

	if value := os.Getenv("FLUENT_URL"); len(value) > 0 {
		result.FluentURL = value
	}

	if value := os.Getenv("CATALOG_SERVICE_URL"); len(value) > 0 {
		result.CatalogServerURL = value
	}

	return result
}
