package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

// BuildSecureHash creates an HMAC-SHA256 hash from all vpc_* and user_* params,
// sorted alphabetically and concatenated as key=value&key=value...
// The hash key is a hex-encoded string as supplied by OnePay.
func BuildSecureHash(params url.Values, hashKey string) string {
	var keys []string
	for k := range params {
		if strings.HasPrefix(k, "vpc_") || strings.HasPrefix(k, "user_") {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	var parts []string
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%s", k, params.Get(k)))
	}
	rawData := strings.Join(parts, "&")

	key, _ := hex.DecodeString(hashKey)
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(rawData))
	return strings.ToUpper(hex.EncodeToString(mac.Sum(nil)))
}

// VerifySecureHash verifies the vpc_SecureHash returned from OnePay.
// It excludes vpc_SecureHash itself when building the verification hash.
func VerifySecureHash(params url.Values, hashKey string, receivedHash string) bool {
	filtered := make(url.Values)
	for k, v := range params {
		if k != "vpc_SecureHash" {
			filtered[k] = v
		}
	}
	expected := BuildSecureHash(filtered, hashKey)
	return strings.EqualFold(expected, receivedHash)
}
