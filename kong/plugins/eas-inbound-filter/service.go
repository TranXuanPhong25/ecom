package main

import (
	"context"
	"time"

	"github.com/Kong/go-pdk"
	pb "github.com/TranXuanPhong25/ecom/jwt-service/proto"
)

var (
	TimeOut = 5000 * time.Millisecond // Default timeout
)

func validateToken(token string, kong *pdk.PDK) {
	ctx, cancel := context.WithTimeout(context.Background(), TimeOut)
	defer cancel()
	r, err := client.ValidateToken(ctx, &pb.TokenRequest{Token: token})
	if err != nil {
		kong.Log.Err("could not validate token: %v", err)
		return
	}

	if r.GetUserId() == "" {
		kong.Response.Exit(401, []byte(`{"error":"Unauthorized"}`), map[string][]string{
			"Content-Type": {"application/json"},
		})
	}
	kong.ServiceRequest.SetHeader("X-User-Id", r.GetUserId())
}
