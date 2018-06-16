package server

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"encoding/base64"
	"fmt"
)

func TestBasicAuthOK(t *testing.T){
	fmt.Println("-----TestBasicAuthOK-----")

	usernamePassword := base64.StdEncoding.EncodeToString([]byte("username1:password1"))

	req, err := http.NewRequest("GET", "/login", nil)
	req.Header.Set("Authorization", "Basic "+usernamePassword)
	if err != nil{
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mockBasicAuth)
	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK{
		t.Errorf("Got %v want %v", recorder.Code, http.StatusOK)
	}

	setCookieHeader := recorder.Header().Get("Set-Cookie")
	if setCookieHeader == ""{
		t.Errorf("auth did not set cookie")
	}
	fmt.Println(setCookieHeader)
}

func TestGenerateAccessToken(t *testing.T){
	at, expires, err := generateAccessToken()
	if err != nil{
		t.Fatal(err)
	}

	fmt.Println("-----TestGenerateAccessToken-----")
	fmt.Println(expires)
	fmt.Println(at)
}

func mockBasicAuth(w http.ResponseWriter, r *http.Request){
	doBasicAuth(w, r, "username1", "password1")
}

func TestBasicAuthGenerate(t *testing.T){
	fmt.Println("-----TestBasicAuthGenerate-----")
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("husr:hpwd")))
}