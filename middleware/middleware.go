package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/azizyogo/bank-app/common/auth"
	constanta "github.com/azizyogo/bank-app/common/const"
	"github.com/azizyogo/bank-app/pkg/config"
)

func Authorize(next http.HandlerFunc, minRoleAccess int8) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			http.Error(w, "Missing auth token", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(tokenHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid auth token", http.StatusUnauthorized)
			return
		}

		secretKey := config.LoadConfig().JWTSecret
		claims, err := auth.GetJWTClaims(tokenParts[1], secretKey)
		if err != nil {
			http.Error(w, "Invalid auth token", http.StatusUnauthorized)
			return
		}

		expirationTime := time.Unix(claims.ExpiresAt.Unix(), 0)
		if expirationTime.Before(time.Now()) {
			http.Error(w, "Token Expired", http.StatusUnauthorized)
			return
		}

		if claims.Role < minRoleAccess {
			http.Error(w, "Unauthorize User", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), constanta.CLAIMS_CONTEXT_KEY, claims)
		r = r.WithContext(ctx)

		next(w, r)
	}
}
