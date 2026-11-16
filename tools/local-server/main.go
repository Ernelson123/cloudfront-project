package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Tell Go to serve files from the "site" directory
	fs := http.FileServer(http.Dir("./site"))

	fmt.Println("🚀 Server running at http://localhost:8080")
	fmt.Println("Press Ctrl+C to stop")

	// Start listening on port 8080
	err := http.ListenAndServe(":8080", fs)
	if err != nil {
		fmt.Println("❌ Error starting server:", err)
	}
}