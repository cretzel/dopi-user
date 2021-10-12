package router

import (
	"context"
	"log"
	"net/http"
)

type AuthMiddleware struct {
}

func (amw *AuthMiddleware) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
		w.Header().Add("Content-Type", "application/json")

		uri := r.RequestURI
		if uri == "/api/user/login" || uri == "/api/user/logout" || uri == "/api/user/refresh" {
			next.ServeHTTP(w, r)
			return
		}

		claims, err := amw.checkAuth(r)
		if err != nil {
			Error(w, 401, "Not authenticated")
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (amw *AuthMiddleware) checkAuth(r *http.Request) (*DopiClaims, error) {
	jwtCookie, err := r.Cookie("dopi_jwt")
	if err != nil {
		return nil, err
	}

	token, err := VerifyToken(jwtCookie.Value)
	if err != nil {
		return nil, err
	}

	dopiClaims, err := GetClaims(token)
	if err != nil {
		return nil, err
	}

	return dopiClaims, nil
}
