package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	// Set environment variables
	os.Setenv("RABBIT_URL", "amqp://test-rabbit-url")
	os.Setenv("MONGO_URL", "mongodb://test-mongo-url")
	os.Setenv("PORT", "8080")
	os.Setenv("AUTH_SERVICE_URL", "http://test-auth-service-url")
	os.Setenv("FLUENT_URL", "test-fluent-url")
	os.Setenv("CATALOG_SERVICE_URL", "http://test-catalog-service-url")
}

func TestLoad(t *testing.T) {
	// Set environment variables
	os.Setenv("RABBIT_URL", "amqp://test-rabbit-url")
	os.Setenv("MONGO_URL", "mongodb://test-mongo-url")
	os.Setenv("PORT", "8080")
	os.Setenv("AUTH_SERVICE_URL", "http://test-auth-service-url")
	os.Setenv("FLUENT_URL", "test-fluent-url")
	os.Setenv("CATALOG_SERVICE_URL", "http://test-catalog-service-url")

	// Load configuration
	config := load()

	// Assert values
	assert.Equal(t, "amqp://test-rabbit-url", config.RabbitURL)
	assert.Equal(t, 8080, config.Port)
	assert.Equal(t, "http://test-auth-service-url", config.SecurityServerURL)
	assert.Equal(t, "test-fluent-url", config.FluentUrl)
	assert.Equal(t, "http://test-catalog-service-url", config.CatalogServerURL)

	// Clear environment variables
	os.Unsetenv("RABBIT_URL")
	os.Unsetenv("MONGO_URL")
	os.Unsetenv("PORT")
	os.Unsetenv("AUTH_SERVICE_URL")
	os.Unsetenv("FLUENT_URL")
	os.Unsetenv("CATALOG_SERVICE_URL")
}
