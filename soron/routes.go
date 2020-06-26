package soron

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)
import "soron/soron/scan"

func setRoutes(router *mux.Router) {

	//Scan CRUD?
	router.HandleFunc("/scan", scan.ScanHandler).Methods("POST")

	//Healtz endpoint
	router.HandleFunc("/healtz", HealtzHandler)

	//Debug echo function
	router.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		body, _ := ioutil.ReadAll(r.Body)
		_, _ = fmt.Fprintf(w, "%s", string(body))
	}).Methods("POST")
}
