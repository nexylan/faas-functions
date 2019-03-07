package function

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"testing"
)

func TestPasswordGeneration(t *testing.T) {
	response := decodeResponse(Handle([]byte{}))

	validateResponseCode(t, response.Code)
	validatePassword(t, response.Password, 8)
}

func TestPasswordGenerationWithCustomLength(t *testing.T) {
	length := 15
	stringifyLength := strconv.Itoa(length)
	response := decodeResponse(Handle([]byte("{\"Length\":" + stringifyLength + "}")))

	validateResponseCode(t, response.Code)
	validatePassword(t, response.Password, length)
}

func decodeResponse(reponse string) Response {
	r := Response{}

	_ = json.Unmarshal([]byte(reponse), &r)

	return r
}

func validateResponseCode(t *testing.T, code int) {
	if http.StatusOK != code {
		t.Errorf("Response code must be %d", http.StatusOK)
	}
}

func validatePassword(t *testing.T, password string, passwordLength int) {
	regex := regexp.MustCompile(`^[a-zA-Z0-9\/\-+!\\:;\[\]{}_$#@]+$`)

	if len(password) != passwordLength {
		t.Errorf("Password %s length is not equal to %d", password, passwordLength)
	}

	if match := regex.FindString(password); len(match) == 0 {
		t.Errorf("Password %s formating is not valid", password)
	}
}
