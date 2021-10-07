package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"log"
	"net/http"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	env_go_proxy := os.Getenv("GOPROXY")

	w.Header().Set("ENV_GO_PROXY", env_go_proxy)
	w.Header().Set("VERSION", os.Getenv("VERSION"))
	w.WriteHeader(http.StatusOK)

	remote_addr := r.RemoteAddr

	for k, v := range r.Header {
		fmt.Fprintln(w, fmt.Sprintf("%s = %s", k, v))
	}
	glog.V(2).Info(fmt.Sprintf("recieve request from: %s and response code: %d\n", remote_addr, http.StatusOK))
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	flag.Parse()
	flag.Set("v", "4")

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/healthz", HealthHandler)

	glog.V(2).Info("Starting http server...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
