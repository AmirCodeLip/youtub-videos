package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type HomeApi struct {
	errorHandler *ErrorHandler
}

func NewHomeApi(errorHandler *ErrorHandler) *HomeApi {
	return &HomeApi{errorHandler}
}

func (homeApi *HomeApi) Hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, parseErr := strconv.ParseInt(vars["id"], 10, 64)
	if parseErr != nil {
		homeApi.errorHandler.BadGateway(w, "something bad has happened please try later.")
		return
	}
	fmt.Fprintln(w, "hello %id", id)
}
