package filepreview

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ServiceConfig struct {
	endPoint  url.URL
	authToken string
}

//Post a file to create a file preview
func GeneratePreview(fileLocation string, config ServiceConfig) ([]byte, error) {
	client := &http.Client{}
	postBody := fmt.Sprintf(`{"url":"%s","sizes":["600>"],"format":"png","metadata":["all"]}`, fileLocation)
	req, err := http.NewRequest("POST", config.endPoint.String(), bytes.NewReader([]byte(postBody)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", config.authToken)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	return body, err
}
