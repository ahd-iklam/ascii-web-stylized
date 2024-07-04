package main

import "net/http"

//The http.ResponseWriter variable is responsible for putting together the response from the HTTP server. When we write to this variable, we are essentially sending data back to the client that made the HTTP request.

//http.Request data structure. It represents an HTTP request received by a server and contains various fields to provide information about the incoming request. Here are the main fields and methods available in the http.Request struct (method Url head body )

func excuteInexPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	// verify the request method before executing the template
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }
	templates.ExecuteTemplate(w, "index.html", nil)
	//executes the template, writing the generated HTML to the http.ResponseWriter
}
