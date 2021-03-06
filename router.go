package main

import (
	"net/http"

	"github.com/gudladona87/kubeinfo/clients"
	"github.com/gudladona87/kubeinfo/handlers"
	"github.com/gudladona87/kubeinfo/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/square/go-jose.v2/json"
)

//health is just the root path of the http server
func health(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	resp := models.Response{Message: "Healthy!"}
	b, _ := json.Marshal(resp)

	w.Write(b)
}

func registerHandlers() (*httprouter.Router, error) {
	router := httprouter.New()

	//Injecting ClientSet dependency to PodInfoHandler.
	cliSet, err := clients.NewClientSet()
	if err != nil {
		return nil, err
	}
	pod := handlers.PodInfoHandler{CoreClient: cliSet.CoreV1()}

	//Register the handles
	router.GET("/", health)
	router.GET("/pods", pod.ListPods)
	return router, nil
}
