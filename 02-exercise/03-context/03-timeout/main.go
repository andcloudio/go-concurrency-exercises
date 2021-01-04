package main

import (
	"io"
	"log"
	"net/http"
	"os"
  "context"
  "time"
)

func main() {

	// TODO: set a http client timeout

	req, err := http.NewRequest("GET", "https://andcloud.io", nil)
	if err != nil {
		log.Fatal(err)
	}

  ctx, cancel := context.WithTimeout(req.Context(), 500*time.Millisecond)
  defer cancel()

  req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	// Close the response body on the return.
	defer resp.Body.Close()

	// Write the response to stdout.
	io.Copy(os.Stdout, resp.Body)
}
