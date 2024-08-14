mockgen -source=./tools/http_client/client.go -destination=./tools/http_client/client_mocks.go -package=http_client
mockgen -source=./rabbit/r_emit/rabbit_channel.go -destination=./rabbit/r_emit/rabbit_channel_mocks.go -package=r_emit
mockgen -source=./tools/db/mongo_collection.go -destination=./tools/db/mongo_collection_mocks.go -package=db