package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Kong/go-pdk"
	pb "github.com/TranXuanPhong25/ecom/jwt-service/proto"
)

var (
	TimeOut = 5000 * time.Millisecond // Default timeout
)

func validateToken(token string, kong *pdk.PDK) {
	isValidToken := false

	if token != "" {
		ctx, cancel := context.WithTimeout(context.Background(), TimeOut)
		defer cancel()
		r, err := client.ValidateToken(ctx, &pb.TokenRequest{Token: token})
		if err != nil {
			err := kong.Log.Err("could not validate token: %v", err)
			if err != nil {
				return
			}
			return
		}
		err = kong.ServiceRequest.SetHeader("X-User-Id", r.GetUserId())
		if err != nil {
			kong.Log.Err("could not set X-User-Id: %v", err)
			return
		}
		isValidToken = r.GetUserId() == ""
	}

	hasAccess, err := authorizeAccess(kong, isValidToken)

	if err != nil {
		kong.Log.Err("authorizeEndpoint: %v", err)
		kong.Response.Exit(
			http.StatusInternalServerError,
			[]byte(fmt.Sprintf(`{"error":"%v"}`, err.Error())),
			map[string][]string{})
	}

	if !hasAccess {
		kong.Response.Exit(http.StatusForbidden, []byte(`{"error":"You don't have access to this resource"}`), map[string][]string{
			"Content-Type": {"application/json"},
		})
	}

	if !isValidToken && !hasAccess {
		kong.Response.Exit(http.StatusUnauthorized, []byte(`{"error":"Invalid Token"}`), map[string][]string{
			"Content-Type": {"application/json"},
		})
	}

}

func authorizeAccess(kong *pdk.PDK, isAuthenticated bool) (bool, error) {
	method, err := kong.Request.GetMethod()
	if err != nil {
		return false, err
	}

	path, err := kong.Request.GetPath()
	if err != nil {
		return false, err
	}

	requestBody := OPARequest{Input: OPAInput{
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

	data := OPAResponse{}

	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		return false, err
	}

	if data.Result.Allow {
		return true, nil
	}

	return false, err
}
