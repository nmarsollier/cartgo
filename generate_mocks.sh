mockgen -source=./service/dao.go -destination=./service/dao_mocks.go -package=service
mockgen -source=./security/dao.go -destination=./security/dao_mocks.go -package=security
mockgen -source=./rabbit/r_emit/rabbit_channel.go -destination=./rabbit/r_emit/rabbit_channel_mocks.go -package=r_emit
mockgen -source=./tools/db/mongo_collection.go -destination=./tools/db/mongo_collection_mocks.go -package=db