package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	log.Fatal(http.ListenAndServe(":"+GetProtEnvaironment(), GetRoutes()))
}

func getInfoCurp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, GetCurpByRenapo(mux.Vars(r)["encode_curp"]))
}

// GetRoutes return the url to api aplication
func GetRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/V1/getInfoCurp/{encode_curp}", getInfoCurp).Name("getInfoCurp").Methods("GET")
	return router
}

// GetCurpByRenapo return string html request
func GetCurpByRenapo(curp string) string {
	resp, _ := http.Post(os.Getenv("URL_STC_1")+curp+os.Getenv("URL_STC_2"), "application/json", nil)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return string(bodyBytes)
}

// ValidateRequest comparate and unencrypst data from two points
func ValidateRequest(key string) bool {
	return true
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
}

// GetProtEnvaironment return the port to deploy the api
func GetProtEnvaironment() string {
	port := os.Getenv("PORT")
	if port != "" {
		return port
	}
	return "80"
}
