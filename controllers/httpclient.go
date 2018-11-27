package controllers

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
)

func httpClient(req *http.Request) []byte {
	tr := &http.Transport{
		MaxIdleConns:       10,
		DisableCompression: true,
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport:     tr,
		CheckRedirect: nil,
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body

}
