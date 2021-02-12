package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// If the route is not dreamCatcher then we throw an error
	// Add a /dreamCatcher path to the URL 
    if r.URL.Path != "/dreamCatcher" {
        http.Error(w, "404 not found.", http.StatusNotFound);
        return;
    }

	// We only support GET commands for now
    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound);
        return;
    }

    fmt.Fprintf(w, "Hey, to add a blog entry add a /form.html to your route");
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// We will parse the information passed into the website
	// Add a /form path to the URL to enter this info
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err);
        return;
    }
    fmt.Fprintf(w, "POST request successful\n");
    name := r.FormValue("name");
    address := r.FormValue("address");
	blog := r.FormValue("blog");

    fmt.Fprintf(w, "Name = %s\n", name);
    fmt.Fprintf(w, "Address = %s\n", address);
	fmt.Fprintf(w, "Blog = %s\n", blog);
}

func main() {
	// Add HTML files from ../static
	fileServer := http.FileServer(http.Dir("./static"));
    http.Handle("/", fileServer);

	// Add a route handle to the server (dreamCatcher)
	// Handle the function for POST 
    http.HandleFunc("/dreamCatcher", helloHandler)
	http.HandleFunc("/form", formHandler)

	// Notify the user that the server is running on port 8080
    fmt.Printf("Starting server at port 8080\n");
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err);
    }
}