package env

import (
	"os"
	"strconv"
)

// Configuration properties
type Configuration struct {
	Port              int    `json:"port"`
	GqlPort           int    `json:"gqlPort"`
	RabbitURL         string `json:"rabbitUrl"`
	SecurityServerURL string `json:"securityServerUrl"`
	CatalogServerURL  string `json:"catalogUrl"`
	FluentUrl         string `json:"fluentUrl"`
	PostgresURL       string `json:"postgresUrl"`
}

var config *Configuration

func new() *Configuration {
	return &Configuration{
		Port:              3003,
		GqlPort:           4003,
		RabbitURL:         "amqp://localhost",
		SecurityServerURL: "http://localhost:3000",
		CatalogServerURL:  "http://localhost:3002",
		FluentUrl:         "localhost:24224",
		PostgresURL:       "postgres://postgres@localhost:5432/postgres",
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

	if value := os.Getenv("RABBIT_URL"); len(value) > 0 {
		result.RabbitURL = value
	}

	if value := os.Getenv("POSTGRES_URL"); len(value) > 0 {
		result.PostgresURL = value
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
		result.FluentUrl = value
	}

	if value := os.Getenv("CATALOG_SERVICE_URL"); len(value) > 0 {
		result.CatalogServerURL = value
	}

	return result
}
