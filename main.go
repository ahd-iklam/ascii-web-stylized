package main

import (
	"fmt"
	"html/template" // to keep the HTML in a separate file, allowing us to change the layout of our edit page without modifying the underlying Go code.
	"net/http"
)

var templates = template.Must(template.ParseFiles("templates/index.html", "templates/result.html"))

//template.Must is a convenience wrapper that panics when passed a non-nil error value, and otherwise returns the *Template unaltered. A panic is appropriate here; if the templates can't be loaded the only sensible thing to do is exit the program.

//template.ParseFile(): will read the html file content and return *template.template

func main() {
	mux := http.NewServeMux()
	// ServeMux:= "Server Multiplexer," plays a crucial role in routing incoming HTTP requests to the appropriate handlers functions When a request comes in, the ServeMux examines the URL path and determines which registered handler should handle the request.

	//ServeMux: match the URL of incoming HTTP requests to the appropriate handler functions. When a request comes in, the ServeMux examines the URL path and determines which registered handler should handle the request.

	// If no handler matches a request, ServeMux can fall back to a default handler, usually resulting in a 404 Not Found response if no other handler is registered.
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	mux.HandleFunc("/", excuteInexPage)
	mux.HandleFunc("/ascii-art", excuteAsciiArtResult)

	err := http.ListenAndServe(":8080", mux) // Use the custom ServeMux
	//ListenAndServe: always return an error. In order to log that error we wrap the function call with log.Fatal.
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
