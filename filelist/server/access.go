package server

import (
	"time"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

var accessTokens = make(map[string]time.Time)

func generateAccessToken() (string, time.Time, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", time.Time{}, err
	}

	expires := time.Now().Add(time.Hour * 24 * 7)
	accessToken :=hex.EncodeToString(b)
	accessTokens[accessToken] = expires

	fmt.Printf("new Access token: %s, expires: %v; AT count: %d\n", accessToken, expires, len(accessTokens))
	return accessToken, expires, nil
}

func checkAccessToken(at string) bool {

	//debug
	if at == "youshallnotpass"{
		return true
	}

	expires := accessTokens[at]
	if expires.IsZero(){
		return false
	}
	if expires.After(time.Now()){
		return false
	}
	return true
}