package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/dlouvier/playlistback/models"
)

// CalculateSecret giving the clientid and clientsecret it calculates the token
func CalculateSecret(clientid string, clientsecret string) string {
	return base64.StdEncoding.EncodeToString([]byte(clientid + ":" + clientsecret))
}

// GetToken ssssss
func GetToken(url string, secret string) string {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte("grant_type=client_credentials")))
	req.Header.Set("Authorization", "Basic "+secret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp := httpClient(req)

	var ans models.BodyResp

	err := json.Unmarshal(resp, &ans)
	if err != nil {
		panic(err)
	}

	return ans.AccessToken
}

// GetTrackInfo ssssss
func GetTrackInfo(url string, token string) string {

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp := httpClient(req)

	if err != nil {
		panic(err)
	}

	return string(resp)
}

// To Debug: blabla, _ := ioutil.ReadAll(req.Body) // ioutil.ReadAll(resp.Request.Header)
// blabla := resp.Request.Header
// blabla := resp.Request.URL
// fmt.Println(blabla)
// fmt.Println("Body: " + string(body))

//fmt.Println("Body: " + string(body))
//fmt.Println("Secret: " + secret)
