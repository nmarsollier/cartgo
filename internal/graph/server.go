package server

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nmarsollier/cartgo/internal/engine/env"
	"github.com/nmarsollier/cartgo/internal/engine/log"
	"github.com/nmarsollier/cartgo/internal/graph/model"
	"github.com/nmarsollier/cartgo/internal/graph/schema"
)

func Start(logger log.LogRusEntry) {
	port := env.Get().GqlPort
	srv := handler.NewDefaultServer(model.NewExecutableSchema(model.Config{Resolvers: &schema.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger.Info("GraphQL playground on port : ", port)
	logger.Error(http.ListenAndServe(fmt.Sprintf(":%d", env.Get().GqlPort), nil))
}
