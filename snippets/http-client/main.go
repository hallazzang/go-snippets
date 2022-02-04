package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetWithDefaultHTTPClient() {
	// Most of the time, it is not recommended using the default HTTP client.
	// But for small examples like this one, it is acceptable.
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() // Don't forget to close resp.Body!
	// Also, you have to read the body to the end of its stream
	// to get the resources correctly disposed.

	// Read the response body as a byte slice.
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("GetWithDefaultHTTPClient: %s\n", string(content))
}

func PostWithCustomHTTPClient() {
	req, err := http.NewRequest("POST", "https://httpbin.org/post", strings.NewReader("foo=bar"))
	if err != nil {
		panic(err)
	}

	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// This time, rather than reading the body as a byte slice,
	// decode its JSON content and store it to a variable.
	var data struct {
		Data    string
		Headers struct {
			UserAgent string `json:"User-Agent"` // Can specify custom JSON key
		}
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		panic(err)
	}

	fmt.Printf("PostWithCustomHTTPClient: Data=%s, UserAgent=%s\n", data.Data, data.Headers.UserAgent)
}

func PostJSON() {
	// This time, create a structured request data.
	data := struct {
		Foo string `json:"some_other_name"`
		Bar int
	}{
		Foo: "baz",
		Bar: 1234,
	}

	// Encode it as JSON.
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "https://httpbin.org/anything", bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json") // Set Content-Type properly.

	// Use http.DefaultClient for sending our request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("PostJSON: %s\n", string(content))
}

func main() {
	GetWithDefaultHTTPClient()
	fmt.Println("====================")
	PostWithCustomHTTPClient()
	fmt.Println("====================")
	PostJSON()
}
