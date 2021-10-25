package https

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	// cert.pem:SSL証明書, key.pem:サーバ用の秘密鍵
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
