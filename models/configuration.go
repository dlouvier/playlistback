package models

// AppConfig is a model to process data from the JSON
type AppConfig struct {
	ClientID     string `json:"ClientID"`
	ClientSecret string `json:"ClientSecret"`
	URLGetToken  string `json:"URLGetToken"`
}
