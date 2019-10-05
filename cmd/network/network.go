package network

import (
	"io/ioutil"
	"log"
	"net/http"
)

// MakeGetRequest makes a HTTP GET request taking in the URL as the parameter.
// The response body is returned as a string.
func MakeGetRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	return string(body), nil
}