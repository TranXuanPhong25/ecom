package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	pb "github.com/TranXuanPhong25/ecom/jwt-service/proto"
	"google.golang.org/grpc"
)

var (
	JWTServiceAddr    string
	OPAServerRouteURL string
	JWTSvcClient      pb.JWTServiceClient
	Port              string
	conn              *grpc.ClientConn
	once              sync.Once
)

// JWT Claims response từ external verify service
type JWTClaims struct {
	UserID string   `json:"user_id"`
	Roles  []string `json:"roles"`
	Email  string   `json:"email"`
	Exp    int64    `json:"exp"`
}

// Handler cho Envoy ExtAuth endpoint
func extAuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract JWT token từ header
	token := ExtractToken(r)
	// Step 1: Verify JWT với external service
	userId, err := validateToken(token)
	if err != nil {
		log.Printf("JWT verification failed: %v", err)
		sendResponse(w, http.StatusForbidden, false, "Invalid token")
		return
	}
	isValidToken := userId != ""
	hasAccess, err := authorizeAccess(r, isValidToken)

	if err != nil {
		log.Printf("authorizeEndpoint: %v", err)
		sendResponse(w, http.StatusInternalServerError, false, "Authorization service error")
		return
	}

	if !hasAccess {
		sendResponse(w, http.StatusForbidden, false, "Access denied")
		return
	}

	if !isValidToken && !hasAccess {
		sendResponse(w, http.StatusUnauthorized, false, "Invalid token")
		return
	}

	// Step 2: Check authorization với external service

	if err != nil {
		log.Printf("Authorization check failed: %v", err)
		sendResponse(w, http.StatusInternalServerError, false, "Authorization service error")
		return
	}

	// Success
	w.Header().Set("X-User-Id", userId)

	sendResponse(w, http.StatusOK, true, "")
}

// Send standardized response
func sendResponse(w http.ResponseWriter, statusCode int, allowed bool, reason string) {
	response := ExtAuthzResponse{
		Allow:  allowed,
		Reason: reason,
	}
	w.WriteHeader(statusCode)
	if !allowed {
		json.NewEncoder(w).Encode(response)
	}
}

// Health check handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Unix(),
		"services": map[string]string{
			"jwt_verify": JWTServiceAddr,
			"authz":      OPAServerRouteURL,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	loadEnv()

	// Setup routes
	http.HandleFunc("/", extAuthHandler)
	http.HandleFunc("/health", healthHandler)

	log.Printf("Starting auth proxy service on port %s", Port)
	log.Printf("JWT Verify URL: %s", JWTServiceAddr)
	log.Printf("Authorization URL: %s", OPAServerRouteURL)
	// Start server
	if err := http.ListenAndServe(":"+Port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
