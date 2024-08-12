mockgen -source=./service/dao.go -destination=./service/mocks.go -package=service
mockgen -source=./security/dao.go -destination=./security/mocks.go -package=security
mockgen -source=./rabbit/r_emit/emitter.go -destination=./rabbit/r_emit/mocks.go -package=r_emit
mockgen -source=./tools/db/mongo_collection.go -destination=./tools/db/mocks.go -package=db