**Package README**

## Introduction

This package provides a simple HTTP server that handles POST requests to create new posts. The server expects JSON-formatted data in the request body with information about the post's title, author, and body. Upon receiving a valid request, it prints the post details to the console and responds with a JSON message indicating success.

## Usage

### Example

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	// Register the handler for the root path
	http.HandleFunc("/", HandlePostRequest)

	// Start the server and listen on port 8080
	log.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
```

### API Endpoint

- **Endpoint:** `/`
- **Method:** POST

#### Request

The server expects a JSON-formatted POST request with the following structure:

```json
{
  "title":  "Post Title",
  "author": "Author Name",
  "body":   "Post Body"
}
```

- `title`: The title of the post.
- `author`: The name of the author of the post.
- `body`: The body content of the post.

#### Response

Upon successful processing of the request, the server responds with a JSON message:

```json
{
  "status":  "success",
  "message": "Author Name posted a new post with title - Post Title"
}
```

In case of errors, the server returns an appropriate HTTP status code along with an error message.

## Development

Please feel free to fork the repository and submit pull requests to add to or modify this package. Please open an issue on the repository if you run across any problems or have any recommendations for enhancements.
