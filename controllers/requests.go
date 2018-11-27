package controllers

import (
	"net/http"
	"os"
)

// GetRequests ssssss
func GetRequests(url string) string {

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN"))
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
