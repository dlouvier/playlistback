package requests

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// BodyResp is a model to process data from the JSON
type BodyResp struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expiration  int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

// GetToken ssssss
func GetToken(url string, secret string) string {
	tr := &http.Transport{
		MaxIdleConns:       10,
		DisableCompression: true,
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport:     tr,
		CheckRedirect: nil,
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte("grant_type=client_credentials")))
	req.Header.Set("Authorization", "Basic "+secret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var ans BodyResp

	err = json.Unmarshal(body, &ans)
	if err != nil {
		panic(err)
	}
	fmt.Println(ans.AccessToken)

	return ans.AccessToken

}

// GetTrackInfo ssssss
func GetTrackInfo(url string, token string) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		DisableCompression: true,
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport:     tr,
		CheckRedirect: nil,
	}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

}

func init() {
	fmt.Println("Hola caracola")
}
