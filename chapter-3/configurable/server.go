package configurable

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "127.0.01:8080",
		Handler: nil,
	}

	server.ListenAndServe()
}
