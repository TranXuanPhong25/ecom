package main

import (
	"bytes"
	"net/http"
	"os"
)

func ExtractToken(r *http.Request) string {
	token := ""
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader != "" {
		token = getTokenFromAuthorizationHeader(authorizationHeader)
		if token != "" {
			return token
		}
	}

	cookieHeader := r.Header.Get("Cookie")
	if cookieHeader != "" {
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

func loadEnv() {
	JWTServiceAddr = os.Getenv("JWT_SERVICE_ADDR")
	if JWTServiceAddr == "" {
		JWTServiceAddr = "jwt-svc.services:50050"
	}

	OPAServerRouteURL = os.Getenv("OPA_SERVER_ROUTE_URL")
	if OPAServerRouteURL == "" {
		OPAServerRouteURL = "http://opa-server.opa:8181/v1/data/route"
	}

	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}
}
