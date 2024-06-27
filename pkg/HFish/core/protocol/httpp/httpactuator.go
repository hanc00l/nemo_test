package httpp

import (
	"io"
	"net/http"
)

func envHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "Servlet/2.0")
	w.Header().Set("Content-Type", "application/vnd.spring-boot.actuator")
	io.WriteString(w, "<!DOCTYPE html><html><header><title>applicationConfig</title></header><body>local.server.port:18080<h2>\"_links\":</h2><h2>\"self\":</h2><h2>\"health\"</h2></body></html>")
}

func envHandlerIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "Servlet/2.0")
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "<!DOCTYPE html><html><header><title>applicationConfig</title></header><body>local.server.port:18080<h2>\"_links\":</h2><h2>\"self\":</h2><h2>\"health\"</h2></body></html>")
}
func startHttpActuatorServer(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", envHandlerIndex)
	mux.HandleFunc("/env", envHandler)

	http.ListenAndServe(addr, mux)
}
