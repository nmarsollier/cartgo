mockgen -source=./tools/httpx/client.go -destination=./tools/httpx/client_mocks.go -package=httpx
mockgen -source=./rabbit/emit/rabbit_channel.go -destination=./rabbit/emit/rabbit_channel_mocks.go -package=emit
mockgen -source=./tools/db/mongo_collection.go -destination=./tools/db/mongo_collection_mocks.go -package=db