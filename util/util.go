package util

import (
	"net/http"
	"fmt"
	"bufio"
	"strings"
)

func FormatHttpRequest(r *http.Request) string {

	var s []string
	s=append(s, fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto))
	s=append(s, fmt.Sprintf("Host: %v", r.Host))
	for k,v := range r.Header{
		s = append(s, fmt.Sprintf("%v: %v", k, v))
	}

	defer r.Body.Close()
	scanner := bufio.NewScanner(r.Body)
	if scanner.Scan(){
		s = append(s, scanner.Text())
	}
	return strings.Join(s, "\n")
}
