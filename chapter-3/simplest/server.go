package simplest

import (
	"net/http"
)

func main() {
	http.ListenAndServe("", nil)
}
