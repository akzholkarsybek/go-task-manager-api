package middleware

import (
	"net/http"
	"strings"
	"context"

	"github.com/Akakazkz/go-task-manager-api/internal/auth"
)

type contextKey string
const UserIDKey contextKey = "userID"

func GetUserID(ctx context.Context) (int64, bool){
	userID, ok := ctx.Value(UserIDKey).(int64)
	return userID, ok
}

func Auth(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == ""{
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "invalid authorization header", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		userID, _, err := auth.ParseJWT(tokenString)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(
			r.Context(), 
			UserIDKey, 
			userID,
		)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}