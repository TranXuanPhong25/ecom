package main

import (
	"bytes"

	"github.com/Kong/go-pdk"
)

func getTokenFromHeaders(kong *pdk.PDK) string {
	token := ""
	authorizationHeader, getAuthorizationError := kong.Request.GetHeader("Authorization")
	if getAuthorizationError == nil {
		token = getTokenFromAuthorizationHeader(authorizationHeader)
		if token != "" {
			return token
		}
	}

	cookieHeader, getCookieError := kong.Request.GetHeader("Cookie")
	if getCookieError == nil {
		token = getTokenFromCookieHeader(cookieHeader)
		if token != "" {
			return token
		}
	}
	return token
}

func getTokenFromAuthorizationHeader(authorizationHeader string) string {
	if len(authorizationHeader) > 7 && authorizationHeader[:7] == "Bearer " {
		return authorizationHeader[7:]
	}
	return ""
}

func getTokenFromCookieHeader(cookieHeader string) string {
	tokenField := []string{"access_token", "AccessToken", "ACCESS_TOKEN"}
	// Parse the cookie header to extract the token
	cookieParts := bytes.Split([]byte(cookieHeader), []byte("; "))
	for _, part := range cookieParts {
		for _, field := range tokenField {
			if bytes.HasPrefix(part, []byte(field+"=")) {
				return string(part[len(field)+1:])
			}
		}
	}
	return ""
}
