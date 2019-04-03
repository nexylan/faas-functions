package function

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"testing"
)

// TestPasswordGeneration test the Handle method with default parameters
func TestPasswordGeneration(t *testing.T) {
	response := decodeResponse(Handle([]byte{}))

	validateResponseCode(t, http.StatusOK, response.Code)
	validatePassword(t, response.Password, 8)
}

// TestPasswordGeneration test the Handle method with custom length provided
func TestPasswordGenerationWithCustomLength(t *testing.T) {
	length := 15
	stringifyLength := strconv.Itoa(length)
	response := decodeResponse(Handle([]byte("{\"length\":" + stringifyLength + "}")))

	validateResponseCode(t, http.StatusOK, response.Code)
	validatePassword(t, response.Password, length)
}

func TestPasswordGenerationWithCustomSpecification(t *testing.T) {
	length, upperCaseNum, digitNum, specialCharNum := 5, 3, 2, 3

	response := decodeResponse(Handle([]byte(
		fmt.Sprintf("{\"length\": %s, \"upper_case_num\": %s,\"digit_num\": %s,\"special_char_num\": %s}",
			strconv.Itoa(length),
			strconv.Itoa(upperCaseNum),
			strconv.Itoa(digitNum),
			strconv.Itoa(specialCharNum),
		),
	)))

	validateResponseCode(t, http.StatusOK, response.Code)
	validatePassword(t, response.Password, upperCaseNum+digitNum+specialCharNum)
}

func TestInvalidJSON(t *testing.T) {
	response := decodeResponseCode(Handle([]byte("{\"length\":\"fail\"}")))

	validateResponseCode(t, http.StatusBadRequest, response.Code)
}

func decodeResponse(reponse string) Response {
	r := Response{}

	_ = json.Unmarshal([]byte(reponse), &r)

	return r
}

func decodeResponseCode(reponse string) ResponseCode {
	rc := ResponseCode{}

	_ = json.Unmarshal([]byte(reponse), &rc)

	return rc
}

func validateResponseCode(t *testing.T, expectedCode int, responseCode int) {
	if expectedCode != responseCode {
		t.Errorf("Response code must be %d not %d", http.StatusOK, responseCode)
	}
}

func validatePassword(t *testing.T, password string, passwordLength int) {
	regex := regexp.MustCompile(`^[a-zA-Z0-9/\-+!\\:;\[\]{}_$#@]+$`)

	if len(password) != passwordLength {
		t.Errorf("Password length (%d) is not equal to %d", len(password), passwordLength)
	}

	if match := regex.FindString(password); len(match) == 0 {
		t.Errorf("Password \"%s\" formating is not valid", password)
	}
}
