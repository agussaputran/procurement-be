package handler

import (
	"context"
	"net/http"

	graph "procurement-be/graph"
	generated "procurement-be/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func (h Handler) Graphql() {
	// set endpoint
	endpoint := h.router.Group("/graph")

	// set resolver
	graphConfig := generated.Config{
		Resolvers: &graph.Resolver{},
	}

	// graphConfig.Directives.LoggedIn = func(ctx context.Context, obj interface{}, next graphql.Resolver, loggedIn model.Login) (res interface{}, err error) {
	// 	if err := middleware.Authorization(ctx, ""); err != nil {
	// 		return nil, errors.New(err.Error())
	// 	}
	// 	return next(ctx)
	// }

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(graphConfig),
	)
	gqlHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		header := r.Header.Get("Authorization")
		AddHeader := context.WithValue(ctx, "Authorization", header)
		// md := metadata.Pairs(
		// 	"Authorization", header,
		// )
		// mdCtx := metadata.NewOutgoingContext(AddHeader, md)

		srv.ServeHTTP(w, r.WithContext(AddHeader))
	})

	endpoint.All("/query", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(gqlHandler)(c.Context())
		return nil
	})

	endpoint.All("/", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(playground.Handler("GraphQL playground", "/graph/query"))(c.Context())
		return nil
	})
}
