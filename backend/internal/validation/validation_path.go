package validation

import (

	vt "github.com/VirusTotal/vt-go"
)

func ValidateURL(path string, apikey string) bool{
	client := vt.NewClient(apikey)

	_ = client
	return false
}