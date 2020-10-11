package controller

import (
	"bytes"
	"encoding/gob"
	"log"
	"net/http"
	"regexp"
)

// GetBytes self explanatory
func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//OnlyCharsMameh parses input to only characters
func OnlyCharsMameh(input string) string {
	// Make a Regex to say we only want letters
	reg, err := regexp.Compile("[^a-z]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(input, "")
}

//OnlyCountryCode parses input to only match country code
func OnlyCountryCode(input string) string {
	// Make a Regex to say we only want letters
	reg, err := regexp.Compile("[^A-Z]")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(input, "")[0:2]
}

//OnlyInt parses input to get only integers
func OnlyInt(input string) string {
	// Make a Regex to say we only want numbers
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(input, "")
}

//WriteError helps returning error
func WriteError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(message))
}
