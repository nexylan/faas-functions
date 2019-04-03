package function

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

// PasswordSpec is function body reference
type PasswordSpec struct {
	Length         int `json:"length,omitempty"`
	UpperCaseNum   int `json:"upper_case_num,omitempty"`
	DigitNum       int `json:"digit_num,omitempty"`
	SpecialCharNum int `json:"special_char_num,omitempty"`
}

// ResponseCode is code of each response
type ResponseCode struct {
	Code int `json:"code"`
}

// Response is the structure of function response
type Response struct {
	ResponseCode
	Password string `json:"password"`
}

// Handle is the function main method triggered each time the function is called
func Handle(req []byte) string {
	var passwordSpec PasswordSpec
	var encodedResponse []byte
	var err error

	passwordSpec.Length = 8
	passwordSpec.UpperCaseNum = 1
	passwordSpec.DigitNum = 1
	passwordSpec.SpecialCharNum = 1

	if len(req) != 0 {
		err = json.Unmarshal(req, &passwordSpec)
	}

	if err != nil {
		marshalJSON, _ := json.Marshal(ResponseCode{Code: http.StatusBadRequest})
		encodedResponse = marshalJSON
	} else {
		marshalJSON, _ := json.Marshal(Response{ResponseCode{Code: http.StatusOK}, generate(passwordSpec)})
		encodedResponse = marshalJSON
	}

	return string(encodedResponse)
}

func generate(passwordSpec PasswordSpec) string {
	rand.Seed(time.Now().UTC().UnixNano())

	upperCaseLetters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	digits := []rune("0123456789")
	specialChars := []rune("/-+!\\:;[]{}_$#@")
	allChars := append(
		append(upperCaseLetters, []rune("abcdefghijklmnopqrstuvwxyz")...),
		append(digits, specialChars...)...,
	)
	upperCaseUsed, digitUsed, specialChar := 0, 0, 0
	passwordLength := passwordSpec.Length
	minimalRequiredCharsLength := passwordSpec.UpperCaseNum + passwordSpec.DigitNum + passwordSpec.SpecialCharNum

	if passwordLength < minimalRequiredCharsLength {
		passwordLength = minimalRequiredCharsLength
	}

	passwordArray := make([]rune, passwordLength)

	for index := range passwordArray {
		if passwordSpec.UpperCaseNum != upperCaseUsed {
			passwordArray[index] = upperCaseLetters[rand.Intn(len(upperCaseLetters))]
			upperCaseUsed++
		} else if passwordSpec.DigitNum != digitUsed {
			passwordArray[index] = digits[rand.Intn(len(digits))]
			digitUsed++
		} else if passwordSpec.SpecialCharNum != specialChar {
			passwordArray[index] = specialChars[rand.Intn(len(specialChars))]
			specialChar++
		} else {
			passwordArray[index] = allChars[rand.Intn(len(allChars))]
		}
	}

	return string(passwordArray)
}
