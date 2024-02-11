package controller

import (
	"encoding/json"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/geffersonFerraz/grinha-de-backend-2024-q1-http3/config"
	"github.com/geffersonFerraz/grinha-de-backend-2024-q1-http3/src/interfaces"
	"github.com/geffersonFerraz/grinha-de-backend-2024-q1-http3/src/usecase"
)

type (
	TransactionsController interface {
		CreateTransaction(w http.ResponseWriter, r *http.Request)
		ConsultaExtrato(w http.ResponseWriter, r *http.Request)
	}

	transactions struct {
		useCases usecase.TransactionsUseCase
	}
)

func NewTransactionController(uc usecase.TransactionsUseCase) TransactionsController {
	return &transactions{useCases: uc}
}

func (t *transactions) ConsultaExtrato(w http.ResponseWriter, r *http.Request) {
	if config.CFG.ENABLE_LOG {
		println("alohomora  " + time.Now().Local().String())
	}
	id := r.PathValue("id")
	w.Header().Add("Content-Type", "application/json")

	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Cliente não encontrado"))
		return
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Cliente não encontrado"))
		return
	}

	result, status := t.useCases.ConsultaExtrato(idInt)
	if status != 200 {
		w.WriteHeader(status)
		json, _ := json.Marshal(result)
		w.Write(json)
		return
	}

	w.WriteHeader(status)
	json, _ := json.Marshal(result)
	w.Write(json)

}

func (t *transactions) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if config.CFG.ENABLE_LOG {
		println("expelliarmus  " + time.Now().Local().String())
	}
	id := r.PathValue("id")
	w.Header().Add("Content-Type", "application/json")

	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Cliente não encontrado"))
		return
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Cliente não encontrado"))
		return
	}

	request := &interfaces.RequestCreateTransaction{}

	err = json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Deu ruim no payload"))
		return
	}

	request.ID = idInt

	operacao := []string{"c", "d"}
	if !slices.Contains(operacao, request.Tipo) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Deu ruim no payload - 'c' ou 'd' animal"))
		return
	}

	if request.Valor <= 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Deu ruim no payload - valor errado "))
		return
	}

	if !(len(request.Descricao) >= 1 && len(request.Descricao) <= 10) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Deu ruim no payload - escreveu um livro, nao uma descrição"))
		return
	}

	result, status := t.useCases.CreateTransaction(*request)
	if status != 200 {
		w.WriteHeader(status)
		json, _ := json.Marshal(result)
		w.Write(json)
		return
	}
	w.WriteHeader(http.StatusOK)
	json, _ := json.Marshal(result)
	w.Write(json)
}
