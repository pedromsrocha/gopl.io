// Search https://xkcd.com/ for a comic by number and download its  image
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Image struct {
	Img string
}

func main() {
	issueNo := url.QueryEscape(os.Args[1])
	resp, err := http.Get("https://xkcd.com/" + issueNo + "/info.0.json")
	if err != nil {
		fmt.Printf("Get failed: %v", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("search query failed: %s", resp.Status)
		os.Exit(1)
	}
	var result Image
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("JSON unmarshling failed: %v", err)
		os.Exit(1)
	}
	resp2, err := http.Get(result.Img)
	if err != nil {
		fmt.Printf("Get failed: %v", err)
		os.Exit(1)
	}
	defer resp2.Body.Close()
	if resp2.StatusCode != http.StatusOK {
		fmt.Printf("search query failed: %s", resp2.Status)
		os.Exit(1)
	}
	fileName := os.Args[1] + ".png"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("create fialed: %v", err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = io.Copy(file, resp2.Body)
	if err != nil {
		fmt.Printf("copy failed: %v", err)
		os.Exit(1)
	}
}
