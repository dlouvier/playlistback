package main

import (
	"fmt"
	"os"

	c "github.com/dlouvier/playlistback/controllers"
)

func main() {
	// Importing the package classes
	c.LoginAPI()
	fmt.Println("Your access token is: " + os.Getenv("TOKEN"))

	response := c.GetRequests("https://api.spotify.com/v1/users/dlouvier")
	fmt.Println("Response from API: \n========================================\n\n" + response)
}
