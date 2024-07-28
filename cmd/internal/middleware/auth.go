package middleware

import (
    "context"
    "net/http"
    "strings"
    "price-alert-app/pkg/auth"
)

func JWTMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Skip authentication for /login endpoint
        if r.URL.Path == "/login" {
            next.ServeHTTP(w, r)
            return
        }

        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing token", http.StatusUnauthorized)
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            http.Error(w, "Invalid token format", http.StatusUnauthorized)
            return
        }

        token := parts[1]
        claims, err := auth.ValidateJWT(token)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Set the claims in the request context if needed
        ctx := context.WithValue(r.Context(), "claims", claims)
        r = r.WithContext(ctx)

        next.ServeHTTP(w, r)
    })
}
