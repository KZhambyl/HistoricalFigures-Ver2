package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/figures", app.createFigureHandler)
	router.HandlerFunc(http.MethodGet, "/v1/figures/:id", app.showFigureHandler)
	router.HandlerFunc(http.MethodPut, "/v1/figures/:id", app.updateFigureHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/figures/:id", app.deleteFigureHandler)

	return router
}
