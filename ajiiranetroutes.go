package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setAJIIRANetHandlerIn(roleRouter *mux.Router) {

	var ajiiraNetHandler AJIIRANetHandler
	roleRouter.HandleFunc("", ajiiraNetHandler.ProcessData).Methods(http.MethodPost)

}
