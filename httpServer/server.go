package httpServerCC

import (
	"fmt"
	"log"
	"net/http"
)

func Launch(port string, message string) {

	if port == "" {
		port = "8080"
	}

	if message == "" {
		message = "Hello, World!"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, message)
	})

	fmt.Println(fmt.Sprintf("Server started on the %s port...", port))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
