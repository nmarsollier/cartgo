mockgen -source=./tools/http_client/client.go -destination=./tools/http_client/client_mocks.go -package=http_client
mockgen -source=./rabbit/emit/rabbit_channel.go -destination=./rabbit/emit/rabbit_channel_mocks.go -package=emit
mockgen -source=./tools/db/mongo_collection.go -destination=./tools/db/mongo_collection_mocks.go -package=db