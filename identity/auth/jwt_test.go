package auth

import (
	"testing"
)

func TestExpired(t *testing.T) {
	jwtExpiryInSeconds = -1
	jwtSecretKey = "test"
	token := CreateJWT(123, "string", false)
	_, _, notVaild := DecodeJWT(token)
	if notVaild == false {
		t.Errorf("got %t, wanted %t", notVaild, true)
	}
}

func TestNotExpired(t *testing.T) {
	jwtExpiryInSeconds = 1
	jwtSecretKey = "test"
	token := CreateJWT(123, "st", false)
	_, _, notValid := DecodeJWT(token)
	if notValid == true {
		t.Errorf("got %t, wanted %t", notValid, false)
	}
}
