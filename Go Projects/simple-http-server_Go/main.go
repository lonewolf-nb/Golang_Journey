package main

import (
	"fmt"
	"log"
	"net/http"
)

// formHandler handles POST requests sent to the "/form" route.
// It reads user-submitted data (name and address) from an HTML form and prints it.
func formHandler(w http.ResponseWriter, r *http.Request) {

	// ParseForm() parses the form data from the POST request body.
	// If an error occurs while parsing, it sends an error message to the browser.
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v", err)
		return
	}

	// If parsing succeeds, acknowledge the POST request.
	fmt.Fprintf(w, "POST request successful\n")

	// Retrieve the values of form fields "name" and "address".
	name := r.FormValue("name")
	address := r.FormValue("address")

	// Display the extracted values in the browser.
	fmt.Fprintf(w, "Name : %s\n", name)
	fmt.Fprintf(w, "Address : %s\n", address)
}

// helloRouteHandler handles GET requests for the "/hello" route.
// It returns a simple greeting message when accessed correctly.
func helloRouteHandler(w http.ResponseWriter, r *http.Request) {

	// Check if the user requested exactly the "/hello" path.
	// If not, return a 404 error.
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Validate that the HTTP method is GET; otherwise, reject the request.
	if method := r.Method; method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	// Send a simple text response back to the browser.
	fmt.Fprintf(w, "hello!")
}

// main is the entry point of the program.
// It configures routes, starts the HTTP server, and listens on port 8080.
func main() {

	// fileServer serves static files (HTML, CSS, JS) from the "static" directory.
	fileServer := http.FileServer(http.Dir("./static"))

	// Route "/" is handled by the file server (serving static files).
	http.Handle("/", fileServer)

	// Route "/form" is handled by the formHandler function.
	http.HandleFunc("/form", formHandler)

	// Route "/hello" is handled by the helloRouteHandler function.
	http.HandleFunc("/hello", helloRouteHandler)

	// Print a message to indicate the server is starting.
	fmt.Printf("Starting server at port 8080\n")

	// Start listening for requests on port 8080.
	// If any error occurs while starting, log.Fatal() stops execution and prints the error.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
