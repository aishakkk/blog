package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBody struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Body   string `json:"body"`
}

type ResponseBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var requestBody RequestBody
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&requestBody); err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		if requestBody.Title == "" || requestBody.Author == "" || requestBody.Body == "" {
			http.Error(w, "Invalid JSON message", http.StatusBadRequest)
			return
		}

		fmt.Printf("Received new post: %s\n", requestBody.Title)
		title := requestBody.Title
		author := requestBody.Author

		response := ResponseBody{
			Status:  "success",
			Message: author + " posted a new post with title - " + title,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
