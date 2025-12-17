package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	pb "github.com/TranXuanPhong25/ecom/services/jwt-service/proto"
)

var (
	TimeOut = 5000 * time.Millisecond // Default timeout
)

func validateToken(token string) (string, error) {
	if token == "" {
		return "", nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), TimeOut)
	defer cancel()
	client := getJWTSvcClient()
	response, err := client.ValidateToken(ctx, &pb.TokenRequest{Token: token})
	if err != nil {
		log.Printf("could not validate token: %v", err)
		return "", err
	}
	return response.GetUserId(), nil

}

func authorizeAccess(r *http.Request, isAuthenticated bool) (bool, error) {
	path := r.Header.Get("x-original-uri")
	if path == "" {
		path = r.URL.Path // fallback
	}
	method := r.Header.Get("x-original-method")
	if method == "" {
		method = r.Method // fallback
	}

	requestBody := AuthzRequest{Input: OPAInput{
		Method:        method,
		Path:          path,
		Authenticated: isAuthenticated,
	}}
	bytesReqBody, err := json.Marshal(requestBody)
	if err != nil {
		return false, err
	}
	request, err := http.NewRequest("POST", OPAServerRouteURL, bytes.NewReader(bytesReqBody))
	if err != nil {
		return false, err
	}
	client := &http.Client{
		Timeout: TimeOut,
	}
	response, err := client.Do(request)
	if err != nil {
		return false, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	if response.StatusCode != 200 {
		return false, err
	}
	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return false, err
	}

	data := AuthzResponse{}

	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		return false, err
	}

	if data.Result.Allow {
		return true, nil
	}

	return false, err
}
