package utils

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
	"ugly-friend/config"
	"ugly-friend/models"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const claimsKey contextKey = "claims"

func JWTMiddleware(next http.Handler, secretKey []byte, skippedRoutes []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip JWT validation for specific routes
		for _, skippedRoute := range skippedRoutes {
			if r.URL.Path == skippedRoute {
				next.ServeHTTP(w, r)
				return
			}
		}

		// Token validation logic here...
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Token validation continued...
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := validateJWT(tokenString, secretKey)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		// Extract the status field from the claims map
		status, ok := (*claims)["status"].(string)
		if !ok || status != "active" {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add claims to context and proceed
		ctx := context.WithValue(r.Context(), claimsKey, claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func validateJWT(tokenString string, secretKey []byte) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signatrue method")
		}

		return secretKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse the token: %s", err.Error())
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to parse Claims")
	}

	return claims, nil
}

func GenerateJWT() (string, error) {
	cfg, err := config.MustLoad()
	if err != nil {
		return "", fmt.Errorf("failed to load config: %w", err)
	}

	claims := &models.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)), // Token is valid for 30 days
			Issuer:    cfg.JWT.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKeyBytes := []byte(cfg.JWT.SecretKey)

	tokenString, err := token.SignedString(secretKeyBytes)
	if err != nil {
		return "", fmt.Errorf("error: failed to sign token: %w", err)
	}

	return tokenString, nil
}
