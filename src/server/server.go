package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/geffersonFerraz/grinha-de-backend-2024-q1-http3/config"
	"github.com/geffersonFerraz/grinha-de-backend-2024-q1-http3/src/controller"
	"github.com/quic-go/quic-go/http3"
)

type (
	Server interface {
		PrepareRoutes()
		Listen(ctx context.Context, url string, port string)
	}
	server struct {
		transaction controller.TransactionsController
		handler     http.Handler
	}
)

func NewServer(controller controller.Controller) Server {

	return &server{transaction: controller.NewTransactionController()}
}

func (s *server) PrepareRoutes() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /clientes/{id}/extrato", s.transaction.ConsultaExtrato)
	mux.HandleFunc("POST /clientes/{id}/transacoes", s.transaction.CreateTransaction)
	s.handler = mux
}

func (s *server) Listen(ctx context.Context, url string, port string) {
	srvUrl := fmt.Sprintf("%s:%s", config.CFG.SERVER_HOST, config.CFG.SERVER_PORT)
	err := http3.ListenAndServeQUIC(srvUrl, "./server-cert.pem", "./server.key", s.handler)

	if err != nil {
		panic(err)
	}

}
