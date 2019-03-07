package function

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

// PasswordSpec is function body reference
type PasswordSpec struct {
	Length int
}

// ResponseCode is code of each response
type ResponseCode struct {
	Code     int
}

// Response is the structure of function response
type Response struct {
	Code     ResponseCode
	Password string
}

// Handle is the function main method triggered each time the function is called
func Handle(req []byte) string {
	var passwordSpec PasswordSpec
	var encodedResponse []byte
	var err error

	passwordSpec.Length = 8

	if len(req) != 0 {
		err = json.Unmarshal(req, &passwordSpec)
	}

	if err != nil {
		marshalJSON, _ := json.Marshal(ResponseCode{Code: http.StatusBadRequest})
		encodedResponse = marshalJSON
	} else {
		marshalJSON, _ := json.Marshal(Response{Code: ResponseCode{Code:http.StatusOK}, Password: generate(passwordSpec.Length)})
		encodedResponse = marshalJSON
	}

	return string(encodedResponse)
}

func generate(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())

	allowedChars := []rune("ABCDEFGHIJKLMNOPQRSTUVXYZabcdefghijklmnopqrstuvxyz0123456789/-+!\\:;[]{}_$#@")
	passwordArray := make([]rune, length)

	for index := range passwordArray {
		passwordArray[index] = allowedChars[rand.Intn(len(allowedChars))]
	}

	return string(passwordArray)
}
