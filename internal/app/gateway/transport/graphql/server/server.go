package server

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/coretech/dtc-api-gateway/internal/app/gateway/service/transaction"
	"github.com/coretech/dtc-api-gateway/internal/app/gateway/transport/graphql/generated"
	"github.com/coretech/dtc-api-gateway/internal/app/gateway/transport/graphql/resolver"
	"github.com/coretech/dtc-api-gateway/pkg/client/hdm-go-transactions-api/genqlient"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8090"

type Server struct {
	server *handler.Server
}

func New() *Server {
	txClient := genqlient.New("http://localhost:7300/v1/query")
	txService := transaction.New(txClient)

	r := resolver.New(txService)
	return &Server{
		server: handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: r})),
	}
}

func (s *Server) Start(ctx context.Context) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(corsMiddleware())

	s.server.Use(&debug.Tracer{})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", s.server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	return http.ListenAndServe(":"+port, router)
}

func (s *Server) Stop(ctx context.Context) error {
	return nil
}

func corsMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fromReq := r.Header.Get("Origin")
			w.Header().Set("Access-Control-Allow-Origin", fromReq)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Accept-Language, Cookie, Authorization, Origin, Referer")
			w.Header().Set("Access-Control-Max-Age", "86400")
			next.ServeHTTP(w, r)
		})
	}
}
