rm -rf ./tests/mocks

set -e

mockgen -source=./internal/cart/repository.go -destination=./tests/mocks/cart_repository.go -package=mocks
mockgen -source=./internal/cart/service.go -destination=./tests/mocks/cart_service.go -package=mocks
mockgen -source=./internal/services/service.go -destination=./tests/mocks/services_service.go -package=mocks
