package main

import (
	"fmt"

	"github.com/dlouvier/playlistback/classes"
)

func main() {
	config := classes.ConfigParser()
	secret := classes.CalculateSecret(config.ClientID, config.ClientSecret)
	fmt.Println("Your secret is: " + secret)
	accesstoken := classes.GetToken(config.URLGetToken, secret)
	fmt.Println("Your access token is: " + accesstoken)

	trackinfo := classes.GetTrackInfo("https://api.spotify.com/v1/tracks/2TpxZ7JUBn3uw46aR7qd6V", accesstoken)
	fmt.Println("Enjoy the Track Info \n========================================\n\n" + trackinfo)
}
