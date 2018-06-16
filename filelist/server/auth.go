package server

import (
	"net/http"
	"strings"
	"errors"
	"encoding/base64"
)

const FLCookieName = "flbrowser-access-token"
const ATHeader = "FLAccessToken"

func doBasicAuth(w http.ResponseWriter, r *http.Request, expectedUsername string, expectedPwd string) error {

	auth := r.Header.Get("Authorization")
	if strings.Index(auth, "Basic") != 0{
		return errors.New("missing <Authorization> header")
	}

	s := strings.Split(auth, " ")
	if len(s) != 2{
		return errors.New("bad <Authorization> header")
	}

	if s[0] != "Basic"{
		return errors.New("only supporting <Basic> auth")
	}

	//username:pwd
	decoded, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil{
		return errors.New("bad <Authorization> header")
	}

	up := string(decoded)
	s = strings.Split(up, ":")
	if len(s) != 2{
		return errors.New("bad <Authorization> header")
	}

	if s[0]!=expectedUsername || s[1] != expectedPwd{
		return errors.New("bad <Authorization> header")
	}

	//generate access token
	accessToken, expires, err := generateAccessToken()
	if err != nil{
		return err
	}

	//set access token cookie
	http.SetCookie(w, &http.Cookie{
		Name:FLCookieName,
		Value:accessToken,
		Expires:expires,
	})

	//login ok
	return nil
}
