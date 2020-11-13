package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// shapes map to globally hold all the items when server is running
var shapes = make(map[string]Shape)

// hello function to respond to http requests
func hello(w http.ResponseWriter, req *http.Request) {
	// print to console
	fmt.Println("/hello with", req.Method)

	// set response header
	w.Header().Set("Content-Type", "application/json")

	// create map for response json
	response := map[string]string{
		"message": "Hello World",
	}

	// encode map and write it to response
	_ = json.NewEncoder(w).Encode(response)
}

// handleShapes will return all the shapes stored globally for GET request
func handleShapes(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(&shapes)
	default:
		http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
	}
}

// handleRectangle will handle POST requests to add new shape in shapes map
func handleRectangle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodPost:
		// read query parameter "name"
		names, ok := req.URL.Query()["name"]

		// check if parameter is present and not empty
		if !ok || len(names[0]) < 1 {
			http.Error(w, "Invalid name for shape", http.StatusBadRequest)
			return
		}

		// check if anything already stored with this name
		_, present := shapes[names[0]]

		if present {
			http.Error(w, "Shape name already exists", http.StatusBadRequest)
			return
		}

		// if body in request is missing throw StatusBadRequest error
		if req.Body == nil {
			http.Error(w, "Body required", http.StatusBadRequest)
			return
		}

		// decode request body json to Rectangle struct
		var rectangle Rectangle
		err := json.NewDecoder(req.Body).Decode(&rectangle)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		// store the Rectangle to Shape map using its reference.
		shapes[names[0]] = &rectangle
		fmt.Println(rectangle)

		// create success message map for json response
		response := map[string]string{
			"message": "shape created successfully",
		}
		_ = json.NewEncoder(w).Encode(response)
	default:
		http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
	}
}

func handleCircle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodPost:
		names, ok := req.URL.Query()["name"] // get the query param "name"

		if !ok || len(names[0]) < 1 {
			http.Error(w, "Name is not defined or is empty", http.StatusBadRequest)
			return
		}

		// check if anything already stored with this name
		_, present := shapes[names[0]]

		if present {
			http.Error(w, "Shape name already exists", http.StatusBadRequest)
			return
		}

		// if body in request is missing throw StatusBadRequest error
		if req.Body == nil {
			http.Error(w, "Body required", http.StatusBadRequest)
			return
		}

		// decode request body json to Circle struct
		var circle Circle
		err := json.NewDecoder(req.Body).Decode(&circle)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		// add circle to the shapes map
		shapes[names[0]] = &circle

		// create response object
		response := map[string]string{
			"message": "Circle shape (" + names[0] + ") was created successfully",
		}
		_ = json.NewEncoder(w).Encode(response)
	default:
		http.Error(w, "This HTTP method is not allowed", http.StatusMethodNotAllowed)
	}

}

func main() {

	// define endpoints for the webserver
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/shapes", handleShapes)
	http.HandleFunc("/shapes/rect", handleRectangle)
	http.HandleFunc("/shapes/circle", handleCircle)

	// start web server and ignore any errors
	_ = http.ListenAndServe(":3000", nil)
}
