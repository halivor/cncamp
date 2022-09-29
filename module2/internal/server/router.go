package server

import "net/http"

func NewRouter() {
	http.HandleFunc("/api/v1/header", HandleCopyHeader)
	http.HandleFunc("/api/v1/env", HandleGetEnv)
	http.HandleFunc("/api/v1/client_ip", HandleClientIP)
	http.HandleFunc("/api/v1/healthz", HandleSuccess)
	http.HandleFunc("/healthz", HandleSuccess)
}
