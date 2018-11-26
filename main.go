package main

import (
	"fmt"

	c "github.com/dlouvier/playlistback/controllers"
)

func main() {
	// Importing the package classes

	config := c.ConfigParser()
	secret := c.CalculateSecret(config.ClientID, config.ClientSecret)
	fmt.Println("Your secret is: " + secret)
	accesstoken := c.GetToken(config.URLGetToken, secret)
	fmt.Println("Your access token is: " + accesstoken)

	trackinfo := c.GetTrackInfo("https://api.spotify.com/v1/tracks/2TpxZ7JUBn3uw46aR7qd6V", accesstoken)
	fmt.Println("Enjoy the Track Info \n========================================\n\n" + trackinfo)
}
