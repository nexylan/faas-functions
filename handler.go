package function

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type PasswordSpec struct {
	Length int
}

type Response struct {
	Code     int
	Password string
}

func Handle(req []byte) string {
	var passwordSpec PasswordSpec

	passwordSpec.Length = 8

	if err := json.Unmarshal(req, &passwordSpec); err != nil && len(req) != 0 {
		log.Fatal("Incorrect JSON", err)
	}

	encodedResponse, _ := json.Marshal(Response{Code: http.StatusOK, Password: generate(passwordSpec.Length)})

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
