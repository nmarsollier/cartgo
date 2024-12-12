mockgen -source=./tools/httpx/client.go -destination=./tools/httpx/client_mocks.go -package=httpx
mockgen -source=./rabbit/emit/rabbit_channel.go -destination=./rabbit/emit/rabbit_channel_mocks.go -package=emit
mockgen -source=./tools/log/logrus_logger.go -destination=./tools/log/logrus_logger_mocks.go -package=log
# mockgen -source=./tools/db/query.go -destination=./tools/db/query_mocks.go -package=db
mockgen -source=./tools/db/update.go -destination=./tools/db/update_mocks.go -package=db
mockgen -source=./tools/db/connection.go -destination=./tools/db/connection_mocks.go -package=db
