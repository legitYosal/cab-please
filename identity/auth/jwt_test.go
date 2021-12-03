package auth

import (
	"testing"
)

func TestExpired(t *testing.T) {
	jwtExpiryInSeconds = -1
	jwtSecretKey = "test"
	token := CreateJWT(123, "string")
	_, _, notVaild := DecodeJWT(token)
	if notVaild == false {
		t.Errorf("got %t, wanted %t", notVaild, true)
	}
}

func TestNotExpired(t *testing.T) {
	jwtExpiryInSeconds = 1
	jwtSecretKey = "test"
	token := CreateJWT(123, "st")
	_, _, notValid := DecodeJWT(token)
	if notValid == true {
		t.Errorf("got %t, wanted %t", notValid, false)
	}
}
