package hello

import (
	"fmt"
	"net/http"
)

// init is run before the application starts serving.
func init() {
	// Handle all requests with path /hello with the helloHandler function.
	http.HandleFunc("/", helloHandler)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from the Go app")
}
