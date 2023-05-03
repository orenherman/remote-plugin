package main

import (
	"io"
	"net/http"
	"os"
	"plugin"
)

func main() {
	err := DownloadFile("plugin.so", "http://localhost:9000/oren.so")
	if err != nil {
		panic(err)
	}
	p, err := plugin.Open("./plugin.so")
	if err != nil {
		panic(err)
	}

	f, err := p.Lookup("Start")
	if err != nil {
		panic(err)
	}

	f.(func())()
}

func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
