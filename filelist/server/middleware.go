package server

import (
	"net/http"
	"fmt"
)

func secure (next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r * http.Request){
		fmt.Println("in secure middleware")
		h:= r.Header.Get(ATHeader)

		if checkAccessToken(h){
			next(w,r)
		} else {
			fmt.Printf("%d http.StatusUnauthorized\n", http.StatusUnauthorized)
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
