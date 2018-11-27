package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"

	"github.com/dlouvier/playlistback/models"
)

// CalculateSecret giving the clientid and clientsecret it calculates the token
func CalculateSecret(clientid string, clientsecret string) string {
	return base64.StdEncoding.EncodeToString([]byte(clientid + ":" + clientsecret))
}

// GetToken ssssss
func GetToken(url string, secret string) {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte("grant_type=client_credentials")))
	req.Header.Set("Authorization", "Basic "+secret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp := httpClient(req)

	var ans models.BodyResp

	err := json.Unmarshal(resp, &ans)
	if err != nil {
		panic(err)
	}

	os.Setenv("TOKEN", ans.AccessToken)
}

// LoginAPI parser the config.json configuration file
func LoginAPI() {
	config := ConfigParser()
	secret := CalculateSecret(config.ClientID, config.ClientSecret)
	GetToken(config.URLGetToken, secret)
}
