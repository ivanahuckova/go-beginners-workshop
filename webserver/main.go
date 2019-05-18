// Code taken from this tutorial by Peter Bourgon
http://howistart.org/posts/go/1/

package main

// First, we need to import the net/http package from the standard library.
import (
	"fmt"
	"net/http"
)

var (
	port = "8080"
)

func main() {
	// Then, in the main function, we install a handler function at the root path of our webserver.
	// http.HandleFunc operates on the default HTTP router, officially called a ServeMux.
	http.HandleFunc("/", hello)
	http.HandleFunc("/goodbye", goodbye)
	
	// We start the HTTP server on port 8080 and with the default ServeMux via http.ListenAndServe.
	// That’s a synchronous, or blocking, call, which will keep the program alive until interrupted.
	fmt.Printf("Listening at localhost:%s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

// The function hello is an http.HandlerFunc, which means it has a specific type signature, and can be passed as an argument to HandleFunc.
// Every time a new request comes into the HTTP server matching the root path, the server will spawn a new goroutine executing the hello function.
// And the hello function simply uses the http.ResponseWriter to write a response to the client. Since http.ResponseWriter.
// Write takes the more general  []byte, or byte-slice, as a parameter, we do a simple type conversion of our string.
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("goodbye!"))
}