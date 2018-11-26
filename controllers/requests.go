package controllers

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dlouvier/playlistback/models"
)

// CalculateSecret giving the clientid and clientsecret it calculates the token
func CalculateSecret(clientid string, clientsecret string) string {
	return base64.StdEncoding.EncodeToString([]byte(clientid + ":" + clientsecret))
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

	var ans models.BodyResp

	err = json.Unmarshal(body, &ans)
	if err != nil {
		panic(err)
	}

	// To Debug: blabla, _ := ioutil.ReadAll(req.Body) // ioutil.ReadAll(resp.Request.Header)
	// blabla := resp.Request.Header
	// blabla := resp.Request.URL
	// fmt.Println(blabla)
	// fmt.Println("Body: " + string(body))

	//fmt.Println("Body: " + string(body))
	//fmt.Println("Secret: " + secret)

	return ans.AccessToken

}

// GetTrackInfo ssssss
func GetTrackInfo(url string, token string) string {
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

	return string(body)
}
