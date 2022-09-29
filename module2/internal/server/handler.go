package server

import (
	"net/http"
	"os"
)

func HandleCopyHeader(resp http.ResponseWriter, req *http.Request) {
	const key = "Content-Type"
	resp.Header().Add(key, req.Header.Get(key))
}

func HandleGetEnv(resp http.ResponseWriter, req *http.Request) {
	const env = "ENV"
	resp.Write([]byte(env + "=" + os.Getenv(env)))
}

func HandleClientIP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte(req.RemoteAddr))
}

func HandleSuccess(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
}
