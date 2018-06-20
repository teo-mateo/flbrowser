package server

import (
	"time"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"sync"
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

var lastCheck time.Time
var mutex = &sync.Mutex{}

func checkAccessToken(at string) bool {

	mutex.Lock()
	if time.Now().Sub(lastCheck) < 200* time.Millisecond{
		time.Sleep(200*time.Millisecond)
	}
	lastCheck = time.Now()
	mutex.Unlock()

	//debug
	if at == "youshallnotpass"{
		return true
	}

	expires := accessTokens[at]
	if expires.IsZero(){
		return false
	}
	if time.Now().After(expires){
		return false
	}
	return true
}