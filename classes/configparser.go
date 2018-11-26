package classes

import (
	"encoding/json"
	"io/ioutil"

	"github.com/dlouvier/playlistback/models"
)

// ConfigParser parser the config.json configuration file
func ConfigParser() models.AppConfig {
	var appconfig models.AppConfig
	file, _ := ioutil.ReadFile("confs/config.json")
	err := json.Unmarshal(file, &appconfig)
	if err != nil {
		panic(err)
	}

	return appconfig
}
