package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Blog struct {
	Name    string `json: "name"`
	Address string `json "address"`
	Blog    string `json "blog"`
}

var blogs []Blog

func getBlogHandler(w http.ResponseWriter, r *http.Request) {
	// Convert all the blogs to JSON
	blogBytes, err := json.Marshal(blogs)

	// Print err if it occurs, otw proceed
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(blogBytes)
}

func createBlogHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of Blog
	blog := Blog{}

	// Parse the errors from thr Form
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the bird from the form info
	blog.Name = r.Form.Get("name")
	blog.Address = r.Form.Get("address")
	blog.Blog = r.Form.Get("blog")

	// Append our existing list of birds with a new entry
	blogs = append(blogs, blog)

	//Finally, we redirect the user to the original HTMl page (located at `/assets/`)
	http.Redirect(w, r, "/assets/", http.StatusFound)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Dream Catcher")
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", handler).Methods("GET")

	// Declare the static file directory and point it to the directory we just made
	staticFileDirectory := http.Dir("./static/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/static/" prefix when looking for files.
	// For example, if we type "/static/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for "./static/static/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/static/", instead of the absolute route itself
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/blog", getBlogHandler).Methods("GET")
	r.HandleFunc("/blog", createBlogHandler).Methods("POST")
	return r
}

func main() {
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := newRouter()
	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err.Error())
	}
}
