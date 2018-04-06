package hello

import (
	"fmt"
	"net/http"
)

// init is run before the application starts serving.
func main() {
	// Handle all requests with path /hello with the helloHandler function.
	http.HandleFunc("/", helloHandler)

	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from the Go app")
}
