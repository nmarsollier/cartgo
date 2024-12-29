rm -rf ./tests/mocks

set -e

mockgen -source=./internal/cart/repository.go -destination=./tests/mocks/cart_repository.go -package=mocks
mockgen -source=./internal/cart/service.go -destination=./tests/mocks/cart_service.go -package=mocks
mockgen -source=./internal/rabbit/consume/consume_article_exist.go -destination=./tests/mocks/consume_article_exist.go -package=mocks
mockgen -source=./internal/rabbit/consume/consume_logout.go -destination=./tests/mocks/consume_logout.go -package=mocks
mockgen -source=./internal/rabbit/consume/consume_order_placed.go -destination=./tests/mocks/consume_order_placed.go -package=mocks
mockgen -source=./internal/services/service.go -destination=./tests/mocks/services_service.go -package=mocks
